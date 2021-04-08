package grpcerrors

import (
	errs "example-service/internal/utils/errors"
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCError ...
type GRPCError struct {
	Status *status.Status
}

func (e *GRPCError) handleDefaultError(gc codes.Code, err error) {
	gStatus := status.New(gc, err.Error())
	e.Status, _ = gStatus.WithDetails(&errdetails.ErrorInfo{Reason: "INTERNAL"})
}

func (e *GRPCError) handleInternalError(gc codes.Code, err *errs.InternalError) {
	gStatus := status.New(gc, err.Error())
	errorInfo := &errdetails.ErrorInfo{
		Reason:   err.Code.String(),
		Metadata: err.Metadata,
	}
	e.Status, _ = gStatus.WithDetails(errorInfo)
}

func (e *GRPCError) handleFormError(err *errs.FromError) {
	gStatus := status.New(codes.InvalidArgument, err.Error())
	errorInfo := &errdetails.ErrorInfo{
		Reason: "BAD_REQUEST",
	}

	badRequest := &errdetails.BadRequest{}

	for _, field := range err.Fields {
		fieldViolation := &errdetails.BadRequest_FieldViolation{
			Field:       field.FieldName,
			Description: field.Description,
		}

		badRequest.FieldViolations = append(badRequest.FieldViolations, fieldViolation)
		fmt.Println("Fields", field.FieldName, field.Description)
	}

	e.Status, _ = gStatus.WithDetails(errorInfo, badRequest)
}

// Err ...
func (e *GRPCError) Err() error {
	return e.Status.Err()
}

// New ...
func New(anError error, gc ...codes.Code) *GRPCError {
	grpcCode := codes.Unknown
	if len(gc) > 0 {
		grpcCode = gc[0]
	}

	grpcError := &GRPCError{}

	switch anError.(type) {
	case *errs.FromError:
		grpcError.handleFormError(anError.(*errs.FromError))
	case *errs.InternalError:
		grpcError.handleInternalError(grpcCode, anError.(*errs.InternalError))
	default:
		grpcError.handleDefaultError(grpcCode, anError)
	}

	return grpcError
}

// Error ...
func Error(anError error, gc ...codes.Code) error {
	err := New(anError, gc...)
	return err.Err()
}
