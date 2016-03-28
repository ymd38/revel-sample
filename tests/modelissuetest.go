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
	okIssueData := IssueData{Title: "TitleTest", Source: "SourceTest", Detail: "DetailTest", Priority: 0, Status: 1, LimitStr: "20160501"}
	ngIssueData := []IssueData{
		{Title: "", Source: "", Detail: "", Priority: 0, Status: 1, LimitStr: ""},
		{Title: "TitleTest", Source: "", Detail: "DetailTest", Priority: 0, Status: 1, LimitStr: "20160501"},
		{Title: "TitleTest", Source: "SourceTest", Detail: "", Priority: 0, Status: 1, LimitStr: "20160501"},
		{Title: "TitleTest", Source: "SourceTest", Detail: "DetailTest", Priority: 0, Status: 1, LimitStr: "MOJI"},
	}

	//正常パターン
	if err := t.issue.Create(&okIssueData); err != nil {
		t.Assert(false)
	}

	//エラーパターン
	for i := 0; i < len(ngIssueData); i++ {
		if err := t.issue.Create(&ngIssueData[i]); err == nil {
			t.Assert(false)
		}
	}
}

func (t *ModelIssueTest) TestUpdate() {
	issueList := t.issue.GetListAll()
	for i := 0; i < len(issueList); i++ {
		issueList[i].Title = "this is testcode"
		if err := t.issue.Update(&issueList[i]); err != nil {
			t.Assert(false)
		}
		break
	}
}

func (t *ModelIssueTest) TestGetList() {
	issueList := t.issue.GetListAll()
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
