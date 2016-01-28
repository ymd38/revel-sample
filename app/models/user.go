package models

import "github.com/revel/revel"

type User struct {
	Id       int    `json:"id"`
	Mail     string `json:"mail"`
	Password string `json:"password"`

	//LimitDate   time.Time
	//CreatedDate time.Time
	//UpdatedDate time.Time
}

func (user User) Validate(v *revel.Validation) {
	v.Required(user.Mail)
	v.Required(user.Password)
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
