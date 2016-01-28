package models

import "github.com/revel/revel"

type Issue struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Source string `json:"source"`
	Detail string `json:"detail"`
	//LimitDate   time.Time
	//CreatedDate time.Time
	//UpdatedDate time.Time
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
}
