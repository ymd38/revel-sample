package models

import (
	"errors"
	"regexp"

	"github.com/revel/revel"
)

type ServiceIssueData struct {
	ID             int    `json:"id"`
	ServiceID      int    `json:"-"`
	IssueID        int    `json:"-"`
	Status         int    `json:"status"`
	Memo           string `json:"memo"`
	Reflectdate    int64  `json:"-"`
	ReflectdateStr string `json:"reflectdate,omitempty" db:"-"`
	Created        int64  `json:"-"`
	CreatedStr     string `json:"created,omitempty" db:"-"`
	Updated        int64  `json:"-"`
	UpdatedStr     string `json:"updated,omitempty" db:"-"`
}

func (serviceissuedata *ServiceIssueData) Validate() error {
	var v revel.Validation

	if serviceissuedata.ServiceID < 1 {
		return errors.New("required service id")
	}
	if serviceissuedata.IssueID < 1 {
		return errors.New("required issue id")
	}

	if serviceissuedata.ReflectdateStr != "" {
		v.Match(serviceissuedata.ReflectdateStr, regexp.MustCompile(`\d{8}`))
		if v.HasErrors() {
			return errors.New("Reflectdate is validate error")
		}
	}
	return nil
}
