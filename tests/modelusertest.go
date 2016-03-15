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
	okUserData := UserData{Name: "h-yamada", MailAddr: "h-yamada@mediba.jp", Password: "h-yamada"}
	ngUserDataList := []UserData{
		{Name: "", MailAddr: "", Password: ""},
		{Name: "", MailAddr: "h-yamada@mediba.jp", Password: "h-yamada"},
		{Name: "h-yamada", MailAddr: "h-yamadamediba.jp", Password: "h-yamada"},
		{Name: "h-yamada", MailAddr: "h-yamada@mediba.jp", Password: ""},
	}

	if err := t.user.Create(&okUserData); err != nil { //正常系
		t.Assert(false)
	}

	//NGパターン
	for i := 0; i < len(ngUserDataList); i++ {
		if err := t.user.Create(&ngUserDataList[i]); err == nil { //タイトルは必須なのでエラーになる
			t.Assert(false)
		}
	}
}

func (t *ModelUserTest) After() {
	t.Rollback()
}
