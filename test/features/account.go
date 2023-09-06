package features

import (
	"akgo/feature/account"
	"github.com/cucumber/godog"
	"github.com/go-playground/validator/v10"
	"log"
)

var registerReq account.RegisterReq
var violations error

func returnSuccessInserted() {
	if violations != nil {
		log.Fatal("Expected violation nil, but", violations)
	}
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
	ctx.Step(`^return success inserted$`, returnSuccessInserted)
}

func FeatureContextAccount(s *godog.ScenarioContext) {
	StepDefinitions(s)
}
