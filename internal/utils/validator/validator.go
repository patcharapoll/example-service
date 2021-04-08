package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	errs "example-service/internal/utils/errors"
	vr "example-service/internal/utils/validator/validator_rule"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// CustomValidator ...
type CustomValidator struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

// Configure ...
func (cv *CustomValidator) Configure() error {
	v := validator.New()

	e := en.New()
	uni := ut.New(e, e)
	trans, found := uni.GetTranslator("en")

	if !found {
		return errors.New("not found en")
	}
	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		return err
	}

	cv.Validator = v
	cv.Trans = trans

	cv.RegisterRule(vr.NewImage64BitRule())
	cv.RegisterRule(vr.NewNoWhiteSpaceRule())
	cv.RegisterRule(vr.NewTokenRule())

	return nil
}

// RegisterRule ...
func (cv *CustomValidator) RegisterRule(rule vr.Rule) {
	vr.RegisterValidationRule(rule, cv.Validator, cv.Trans)
}

// Validate ...
func (cv *CustomValidator) Validate(structRule interface{}) error {
	if err := cv.Validator.Struct(structRule); err != nil {
		formError := errs.NewFormError("Wrong Input")

		for _, e := range err.(validator.ValidationErrors) {
			jsonFieldName := e.Field()
			if field, ok := reflect.TypeOf(structRule).Elem().FieldByName(e.Field()); ok {
				if jsonTag, ok := field.Tag.Lookup("json"); ok {
					jsonFieldName = strings.Split(jsonTag, ",")[0]
				}
			}

			param := e.Param()
			if param != "" {
				param = "=" + param
			}

			errorMsg := fmt.Sprintln("ERROR_%s%s: %s", strings.ToUpper(e.Tag()), param, e.Translate(cv.Trans))

			formError.AddErrorField(jsonFieldName, errorMsg)
		}

		return formError
	}
	return nil
}

// NewCustomValidator ...
func NewCustomValidator() (*CustomValidator, error) {
	var cv = &CustomValidator{}
	if err := cv.Configure(); err != nil {
		return nil, err
	}
	return cv, nil
}
