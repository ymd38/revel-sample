package models

import (
	"errors"
	"regexp"

	"github.com/revel/revel"
)

type ServiceData struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Start      int64  `json:"-"`
	StartStr   string `json:"start,,omitempty" db:"-"`
	End        int64  `json:"-"`
	EndStr     string `json:"end,omitempty" db:"-"`
	Created    int64  `json:"-"`
	CreatedStr string `json:"created,omitempty" db:"-"`
	Updated    int64  `json:"-"`
	UpdatedStr string `json:"updated,omitempty" db:"-"`
}

func (servicedata *ServiceData) Validate() error {
	var v revel.Validation
	v.Check(
		servicedata.Name,
		revel.Required{},
		revel.MaxSize{1024},
		revel.MinSize{1},
	)
	if v.HasErrors() {
		return errors.New("name is validate error")
	}

	if servicedata.StartStr != "" {
		v.Match(servicedata.StartStr, regexp.MustCompile(`\d{8}`))
		if v.HasErrors() {
			return errors.New("limit is validate error")
		}
	}

	if servicedata.EndStr != "" {
		v.Match(servicedata.EndStr, regexp.MustCompile(`\d{8}`))
		if v.HasErrors() {
			return errors.New("end is validate error")
		}
	}

	return nil
}
