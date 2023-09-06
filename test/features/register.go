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
var violations error
var t *testing.T

func returnViolationsIsNil() {
	assert.Equal(t, nil, violations)
}

func returnViolationsContains(message string) {
	assert.Equal(t, true, strings.Contains(violations.Error(), message))
}

func userRegister() {
	validate := validator.New()
	violations = validate.Struct(registerReq)
}

func payloadRegister(dataTable *godog.Table) {
	mapFields(&registerReq, dataTable)
}

func stepDefinitions(ctx *godog.ScenarioContext) {
	ctx.Step(`^the following payload:$`, payloadRegister)
	ctx.Step(`^user register$`, userRegister)
	ctx.Step(`^return violations is nil$`, returnViolationsIsNil)
	ctx.Step(`^return violations contains "([^"]*)"$`, returnViolationsContains)
}

func FeatureContextRegister(s *godog.ScenarioContext) {
	stepDefinitions(s)
}
