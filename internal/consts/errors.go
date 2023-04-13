package consts

import "github.com/gogf/gf/v2/errors/gcode"

var (
	ErrInvalidSignature = gcode.New(10001, "signature is invalid", nil)
	ErrInternal         = gcode.New(10002, "internal error", nil)
)
