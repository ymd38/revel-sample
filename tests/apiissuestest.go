package tests

import "github.com/revel/revel/testing"

type ApiIssuesTest struct {
	testing.TestSuite
}

func (t *ApiIssuesTest) Before() {
	println("Set up")
}

func (t *ApiIssuesTest) TestThatIndexPageWorks() {
	t.Get("/api/issues")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *ApiIssuesTest) After() {
	println("Tear down")
}
