package models

import "github.com/revel/revel"

type Issue struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Source     string `json:"source"`
	Detail     string `json:"detail"`
	Priority   int    `json:"priority"` //対応状況(0:重要, 1:緊急, 2:その他)',
	Status     int    `json:"status"`   //`対応状況(0:完了, 1:未対応)',
	Limit      string `json:"limit"`    //期限日 YYYYMMDD
	Created    int64  `json:"-"`
	CreatedStr string `json:"created,omitempty" db:"-"`
	Updated    int64  `json:"-"`
	UpdatedStr string `json:"updated,omitempty" db:"-"`
}

func (issue Issue) Validate(v *revel.Validation) {
	v.Check(
		issue.Title,
		revel.Required{},
		revel.MaxSize{1024},
		revel.MinSize{1},
	)

	v.Check(
		issue.Source,
		revel.Required{},
		revel.MaxSize{1024},
		revel.MinSize{1},
	)

	v.Check(
		issue.Detail,
		revel.Required{},
		revel.MaxSize{5120},
		revel.MinSize{1},
	)

	v.Check(
		issue.Priority,
		revel.Required{},
	)

	v.Check(
		issue.Status,
		revel.Required{},
	)

	v.Check(
		issue.Limit,
		revel.Required{},
	)
}
