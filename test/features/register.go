package features

import (
	"akgo/db"
	"akgo/feature/auth"
	"github.com/cucumber/godog"
	"github.com/go-playground/assert/v2"
	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/mock"
	"strings"
	"testing"
)

var registerReq auth.RegisterReq
var violation error
var commandTag pgconn.CommandTag
var t *testing.T

func returnViolationIsNil() {
	assert.Equal(t, nil, violation)
}

func commandTagContains(data string) {
	assert.Equal(t, true, strings.Contains(commandTag.String(), data))
}

func returnViolationContains(message string) {
	assert.Equal(t, true, strings.Contains(violation.Error(), message))
}

func userRegister() {
	mockPool := new(PgMockPool)
	expectedCommandTag := pgconn.CommandTag("INSERT")
	mockPool.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(expectedCommandTag, nil)
	db.Pg = mockPool
	commandTag, violation = auth.DoRegister(registerReq)
}

func payloadRegister(dataTable *godog.Table) {
	mapFields(&registerReq, dataTable)
}

func stepDefinitions(ctx *godog.ScenarioContext) {
	ctx.Step(`^the following payload:$`, payloadRegister)
	ctx.Step(`^user register$`, userRegister)
	ctx.Step(`^return violation is nil$`, returnViolationIsNil)
	ctx.Step(`^return violation contains "([^"]*)"$`, returnViolationContains)
	ctx.Step(`^command tag contains "([^"]*)"$`, commandTagContains)
}

func FeatureContextRegister(s *godog.ScenarioContext) {
	stepDefinitions(s)
}
