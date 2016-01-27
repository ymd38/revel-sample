package models

import "github.com/revel/revel"

type Issue struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Source string `json:"source"`
	//LimitDate   time.Time
	//CreatedDate time.Time
	//UpdatedDate time.Time
}

func (issue Issue) Validate(v *revel.Validation) {
	v.Required(issue.Title)
	v.Required(issue.Source)
	/*
		v.Match(booking.CardNumber, regexp.MustCompile(`\d{16}`)).
			Message("Credit card number must be numeric and 16 digits")

		v.Check(booking.NameOnCard,
			revel.Required{},
			revel.MinSize{3},
			revel.MaxSize{70},
		)
	*/
}
