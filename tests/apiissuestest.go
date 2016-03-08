package tests

import (
	. "security-cop/app/models"

	"github.com/revel/revel/testing"
)

type ApiIssuesTest struct {
	Issue
	testing.TestSuite
}

func (t *ApiIssuesTest) Before() {
	println("Set up")
}

func (t *ApiIssuesTest) TestIssues() {
	t.Get("/api/issues")
	t.AssertOk()
	t.AssertContentType("application/javascript; charset=utf-8")
}

func (t *ApiIssuesTest) After() {
	println("Tear down")
}
