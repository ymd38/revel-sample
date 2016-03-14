package models

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

/* TODO:Validate */
