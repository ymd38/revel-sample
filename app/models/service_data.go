package models

import (
	"regexp"

	"github.com/revel/revel"
)

type ServiceData struct {
	Id         int    `json:"id"`
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

func (servicedata *ServiceData) Validate(v *revel.Validation) {
	v.Check(
		servicedata.Name,
		revel.Required{},
		revel.MaxSize{1024},
		revel.MinSize{1},
	).Message("name is validate error")
	if v.HasErrors() {
		return
	}

	if servicedata.StartStr != "" {
		v.Match(servicedata.StartStr, regexp.MustCompile(`\d{8}`)).Message("limit is validate error")
		if v.HasErrors() {
			return
		}
	}

	if servicedata.EndStr != "" {
		v.Match(servicedata.EndStr, regexp.MustCompile(`\d{8}`)).Message("limit is validate error")
		if v.HasErrors() {
			return
		}
	}

}
