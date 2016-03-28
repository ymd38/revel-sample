package tests

import (
	. "security-cop/app/models"

	"github.com/revel/revel/testing"
)

type ModelUserTest struct {
	testing.TestSuite
	user User
	GorpController
}

func (t *ModelUserTest) Before() {
	InitDB()
	t.Begin()
}

func (t *ModelUserTest) TestCreate() {
	okUserData := UserData{MailAddr: "test@mediba.jp", Password: "test"}
	ngUserDataList := []UserData{
		{MailAddr: "", Password: ""},
		{MailAddr: "testmediba.jp", Password: "test"},
		{MailAddr: "test@mediba.jp", Password: ""},
	}

	if err := t.user.Create(&okUserData); err != nil { //正常系
		t.Assert(false)
	}

	//NGパターン
	for i := 0; i < len(ngUserDataList); i++ {
		if err := t.user.Create(&ngUserDataList[i]); err == nil {
			t.Assert(false)
		}
	}
}

func (t *ModelUserTest) TestGetToken() {
	//Token取得
	/*	okUserData := UserData{MailAddr: "test@mediba.jp", Password: "test"}
		if token := t.user.GetToken(&okUserData); token == "" {
			t.Assert(false)
		}
	*/
}

func (t *ModelUserTest) After() {
	t.Rollback()
}
