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
	serviceData := &ServiceData{}
	if err := t.service.Create(serviceData); err == nil { //タイトルは必須なのでエラーになる
		t.Assert(false)
	}

	serviceData.Name = "test"
	serviceData.StartStr = "test"
	if err := t.service.Create(serviceData); err == nil { //日付フォーマットエラーになる
		t.Assert(false)
	}
	serviceData.StartStr = "20160101"
	serviceData.EndStr = "test"
	if err := t.service.Create(serviceData); err == nil { //日付フォーマットエラーになる
		t.Assert(false)
	}

	serviceData.EndStr = "20251231"
	if err := t.service.Create(serviceData); err != nil { //正常系
		t.Assert(false)
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
