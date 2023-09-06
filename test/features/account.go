package features

import (
	"akgo/feature/account"
	"github.com/cucumber/godog"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	"testing"
)

var registerReq account.RegisterReq
var violations error
var t *testing.T

func returnViolationsIsNil() {
	assert.Equal(t, nil, violations)
}

func userRegister() {
	validate := validator.New()
	violations = validate.Struct(registerReq)
}

func payloadRegister(dataTable *godog.Table) {
	mapFields(&registerReq, dataTable)
}

func StepDefinitions(ctx *godog.ScenarioContext) {
	ctx.Step(`^the following valid register payload:$`, payloadRegister)
	ctx.Step(`^user register$`, userRegister)
	ctx.Step(`^return violations is nil$`, returnViolationsIsNil)
}

func FeatureContextAccount(s *godog.ScenarioContext) {
	StepDefinitions(s)
}
