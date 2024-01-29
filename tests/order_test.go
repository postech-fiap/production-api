package tests

import (
	"errors"
	"github.com/go-bdd/gobdd"
	"github.com/postech-fiap/production-api/internal/core/domain"
	"strconv"
	"testing"
)

func change(t gobdd.StepTest, ctx gobdd.Context, var1, var2 string) {
	order := domain.Order{Status: domain.Status(var1)}
	ctx.Set("validStatus", order.IsValidStatus(domain.Status(var2)))
}

func check(t gobdd.StepTest, ctx gobdd.Context, valid string) {
	received, err := ctx.GetBool("validStatus")
	if err != nil {
		t.Error(err)
		return
	}

	validBoolean, _ := strconv.ParseBool(valid)

	if validBoolean != received {
		t.Error(errors.New("the math does not work for you"))
	}
}

func TestScenarios(t *testing.T) {
	suite := gobdd.NewSuite(t)
	suite.AddStep(`I change from (\w+) to (\w+)`, change)
	suite.AddStep(`the result should equal (\w+)`, check)
	suite.Run()
}
