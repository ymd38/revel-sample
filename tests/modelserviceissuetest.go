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
	type InputIDs struct {
		issueid, serviceid int
	}

	okInputIDs := InputIDs{1, 1}
	ngInputIDs := []InputIDs{{0, 0}, {0, 1}, {1, 0}}

	if err := t.serviceissue.Create(okInputIDs.issueid, okInputIDs.serviceid); err != nil {
		t.Assert(false)
	}

	for i := 0; i < len(ngInputIDs); i++ {
		if err := t.serviceissue.Create(ngInputIDs[i].issueid, ngInputIDs[i].serviceid); err == nil {
			t.Assert(false)
		}
	}
}

func (t *ModelServiceIssueTest) After() {
	t.Rollback()
}
