package tests

import (
	. "security-cop/app/models"

	"github.com/revel/revel/testing"
)

type ModelServiceIssueTest struct {
	testing.TestSuite
	serviceissue ServiceIssue
	GorpController
}

func (t *ModelServiceIssueTest) Before() {
	InitDB()
	t.Begin()
}

func (t *ModelServiceIssueTest) TestCreate() {
	issueid, serviceid := 0, 0
	if err := t.serviceissue.Create(issueid, serviceid); err == nil {
		t.Assert(false)
	}

	issueid = 1
	if err := t.serviceissue.Create(issueid, serviceid); err == nil {
		t.Assert(false)
	}

	serviceid = 1
	if err := t.serviceissue.Create(issueid, serviceid); err != nil {
		t.Assert(false)
	}
}

func (t *ModelServiceIssueTest) After() {
	t.Rollback()
}
