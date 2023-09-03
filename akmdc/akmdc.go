package akmdc

import (
	"context"
)

var Ctx context.Context
type contextKey string
const MdcKey = contextKey("mdc")
type MDC map[string]interface{}
func GetMDC() MDC {
	if mdc, ok := Ctx.Value(MdcKey).(MDC); ok {
		return mdc
	}
	return nil
}