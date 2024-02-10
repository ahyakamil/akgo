package akmdc

import (
	"context"
	"fmt"
)

var Ctx map[string]context.Context

type contextKey string

const MdcKey = contextKey("mdc")

type MDC map[string]interface{}

func GetMDC() MDC {
	return GetMDCWithCtx("default")
}

func GetMDCWithCtx(ctxKey string) MDC {
	ctx, ok := Ctx[ctxKey]
	if !ok {
		print(fmt.Errorf("context with key %s not found", ctxKey))
		return nil
	}

	mdc, ok := ctx.Value(MdcKey).(MDC)
	if !ok {
		print(fmt.Errorf("mdc not found in context with key %s", ctxKey))
		return nil
	}
	return mdc
}

func Init() {
	Ctx = make(map[string]context.Context)
	Ctx["default"] = context.WithValue(context.Background(), MdcKey, make(MDC))
}
