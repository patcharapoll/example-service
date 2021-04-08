package controller

import "go.uber.org/fx"

// Module ...
var Module = fx.Provide(
	NewHealthZController,
	NewPingPongController,
)
