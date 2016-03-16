package models

import (
	"errors"
	"regexp"

	"github.com/revel/revel"
)

// UserData is struct of user data management
// that uses database and api request/response
type UserData struct {
	ID       int    `json:"id"`
	MailAddr string `json:"mailaddr"`
	Password string `json:"password"`
	Created  int64  `json:"-"`
	Updated  int64  `json:"-"`
}

// Validate for userData of struct
func (userdata *UserData) Validate() error {
	var v revel.Validation

	v.Match(userdata.MailAddr, regexp.MustCompile(`([a-zA-Z0-9])+@mediba.jp`))
	if v.HasErrors() {
		return errors.New("mail address is validate error")
	}

	v.Check(
		userdata.Password,
		revel.Required{},
		revel.MinSize{4},
	)
	if v.HasErrors() {
		return errors.New("password is validate error")
	}

	return nil
}
