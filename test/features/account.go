package features

import (
	"akgo/feature/account"
	"github.com/cucumber/godog"
	"github.com/go-playground/validator/v10"
	"log"
)

var registerReq account.RegisterReq

func returnSuccessInserted() error {
	return godog.ErrPending
}

func userRegister() error {
	validate := validator.New()
	violations := validate.Struct(registerReq)
	log.Print(violations)
	return violations
}

func payloadRegister(dataTable *godog.Table) error {
	err := mapFields(&registerReq, dataTable)
	log.Print(err)
	return err
}

func StepDefinitions(ctx *godog.ScenarioContext) {
	ctx.Step(`^the following valid register payload:$`, payloadRegister)
	ctx.Step(`^user register$`, userRegister)
	ctx.Step(`^return success inserted$`, returnSuccessInserted)
}

func FeatureContextAccount(s *godog.ScenarioContext) {
	StepDefinitions(s)
}
