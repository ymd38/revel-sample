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
	Name     string `json:"name"`
	MailAddr string `json:"mailaddr"`
	Password string `json:"password"`
	Created  int64  `json:"-"`
	Updated  int64  `json:"-"`
}

// Validate for userData of struct
func (user *UserData) Validate() error {
	var v revel.Validation
	v.Match(user.Name, regexp.MustCompile(`\w{1,30}`))
	if v.HasErrors() {
		return errors.New("name is validate error")
	}

	/*
		v.Check(
			issuedata.Detail,
			revel.Required{},
			revel.MaxSize{5120},
			revel.MinSize{1},
		)
		if v.HasErrors() {
			return errors.New("detail is validate error")
		}

		v.Match(strconv.Itoa(issuedata.Priority), regexp.MustCompile(`\d{1}`))
		if v.HasErrors() {
			return errors.New("priority is validate error")
		}

		v.Match(strconv.Itoa(issuedata.Status), regexp.MustCompile(`\d{1}`))
		if v.HasErrors() {
			return errors.New("status is validate error")
		}
		v.Match(issuedata.LimitStr, regexp.MustCompile(`\d{8}`))
		if v.HasErrors() {
			return errors.New("limit is validate error")
		}
			)
	*/
	return nil
}
