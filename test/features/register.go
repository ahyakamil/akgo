package features

import (
	"akgo/feature/account"
	"github.com/cucumber/godog"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	"strings"
	"testing"
)

var registerReq account.RegisterReq
var violation error
var t *testing.T

func returnViolationIsNil() {
	assert.Equal(t, nil, violation)
}

func returnViolationContains(message string) {
	assert.Equal(t, true, strings.Contains(violation.Error(), message))
}

func userRegister() {
	validate := validator.New()
	violation = validate.Struct(registerReq)
}

func payloadRegister(dataTable *godog.Table) {
	mapFields(&registerReq, dataTable)
}

func stepDefinitions(ctx *godog.ScenarioContext) {
	ctx.Step(`^the following payload:$`, payloadRegister)
	ctx.Step(`^user register$`, userRegister)
	ctx.Step(`^return violation is nil$`, returnViolationIsNil)
	ctx.Step(`^return violation contains "([^"]*)"$`, returnViolationContains)
}

func FeatureContextRegister(s *godog.ScenarioContext) {
	stepDefinitions(s)
}
