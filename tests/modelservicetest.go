package tests

import (
	. "security-cop/app/models"

	"github.com/revel/revel/testing"
)

type ModelServiceTest struct {
	testing.TestSuite
	service Service
}

func (t *ModelServiceTest) Before() {
	InitDB()
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
	println("Tear down")
}
