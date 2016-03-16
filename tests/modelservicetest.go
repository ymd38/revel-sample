package tests

import (
	. "security-cop/app/models"

	"github.com/revel/revel/testing"
)

type ModelServiceTest struct {
	testing.TestSuite
	service Service
	GorpController
}

func (t *ModelServiceTest) Before() {
	InitDB()
	t.Begin()
}

func (t *ModelServiceTest) TestCreate() {
	okServiceData := ServiceData{Name: "test", StartStr: "20160101", EndStr: "20251231"}
	ngServiceData := []ServiceData{
		{Name: "", StartStr: "20160101", EndStr: "20251231"},
		{Name: "test", StartStr: "test", EndStr: "20251231"},
		{Name: "test", StartStr: "20160101", EndStr: "test"},
	}

	//OKパターン
	if err := t.service.Create(&okServiceData); err != nil {
		t.Assert(false)
	}

	//NGパターン
	for i := 0; i < len(ngServiceData); i++ {
		if err := t.service.Create(&ngServiceData[i]); err == nil {
			t.Assert(false)
		}
	}
}

func (t *ModelServiceTest) TestGetList() {
	issueList := t.service.GetList()
	if len(issueList) == 0 {
		t.AssertNotFound()
	}
}

func (t *ModelServiceTest) TestGetByID() {
	issueList := t.service.GetByID(1)
	if len(issueList) == 0 {
		t.AssertNotFound()
	}
}

func (t *ModelServiceTest) After() {
	t.Rollback()
}
