package tests

import (
	. "security-cop/app/models"

	"github.com/revel/revel/testing"
)

type ModelIssueTest struct {
	testing.TestSuite
	issue Issue
	GorpController
}

func (t *ModelIssueTest) Before() {
	InitDB()
	t.Begin()
}

func (t *ModelIssueTest) TestCreate() {
	issueData := &IssueData{}
	if err := t.issue.Create(issueData); err == nil { //Titleは必須なのでエラーになる
		t.Assert(false)
	}

	issueData.Title = "TitleTest"
	if err := t.issue.Create(issueData); err == nil { //Sourceは必須なのでエラーになる
		t.Assert(false)
	}

	issueData.Source = "SourceTest"
	if err := t.issue.Create(issueData); err == nil { //Detailは必須なのでエラーになる
		t.Assert(false)
	}

	issueData.Detail = "DetailTest"
	issueData.Priority = 0
	issueData.Status = 1

	issueData.LimitStr = "test"
	if err := t.issue.Create(issueData); err == nil { //日付フォーマットエラーになる
		t.Assert(false)
	}
	issueData.LimitStr = "20160501"
	if err := t.issue.Create(issueData); err != nil { //正常系
		t.Assert(false)
	}

}

func (t *ModelIssueTest) TestGetList() {
	issueList := t.issue.GetList()
	if len(issueList) == 0 {
		t.AssertNotFound()
	}
}

func (t *ModelIssueTest) TestGetByID() {
	issueList := t.issue.GetByID(1)
	if len(issueList) == 0 {
		t.AssertNotFound()
	}
}

func (t *ModelIssueTest) After() {
	t.Rollback()
}
