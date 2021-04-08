package controller

import (
	"context"
	port_v1 "example-service/pkg/api/v1"
)

// PingPongController ...
type PingPongController struct{}

// NewPingPongController ...
func NewPingPongController() *PingPongController {
	return new(PingPongController)
}

// StartPing ...
func (ctrl *PingPongController) StartPing(ctx context.Context, req *port_v1.PingPong) (*port_v1.PingPong, error) {
	return req, nil
}
