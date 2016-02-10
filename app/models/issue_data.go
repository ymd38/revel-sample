package models

import (
	"regexp"
	"strconv"

	"github.com/revel/revel"
)

type IssueData struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Source     string `json:"source"`
	Detail     string `json:"detail"`
	Priority   int    `json:"priority"`     //対応状況(0:重要, 1:緊急, 2:その他)',
	Status     int    `json:"status"`       //`対応状況(0:完了, 1:未対応)',
	Limit      int64  `json:"-"`            //期限日 YYYYMMDD
	LimitStr   string `json:"limit" db:"-"` //期限日 YYYYMMDD
	Created    int64  `json:"-"`
	CreatedStr string `json:"created,omitempty" db:"-"`
	Updated    int64  `json:"-"`
	UpdatedStr string `json:"updated,omitempty" db:"-"`
}

func (issuedata *IssueData) Validate(v *revel.Validation) {
	v.Check(
		issuedata.Title,
		revel.Required{},
		revel.MaxSize{1024},
		revel.MinSize{1},
	).Message("titel is validate error")

	v.Check(
		issuedata.Source,
		revel.Required{},
		revel.MaxSize{1024},
		revel.MinSize{1},
	).Message("source is validate error")

	v.Check(
		issuedata.Detail,
		revel.Required{},
		revel.MaxSize{5120},
		revel.MinSize{1},
	).Message("detail is validate error")

	v.Check(
		issuedata.Priority,
		revel.Required{},
	).Message("priority is validate error")

	v.Match(strconv.Itoa(issuedata.Status), regexp.MustCompile(`\d{1}`)).Message("status is validate error")
	v.Match(issuedata.LimitStr, regexp.MustCompile(`\d{8}`)).Message("limit is validate error")

	if v.HasErrors() {
		errmap := v.ErrorMap()
		for e := range errmap {
			revel.ERROR.Println(e)
		}
	}
}

/* サービス毎の対応状況 */
type ServiceIssueView struct {
	ServiceID        int    `json:"-"`
	IssueId          int    `json:"issue_id"`
	IssueTitle       string `json:"issue_title"`
	IssuePriority    int    `json:"-"`
	IssuePriorityStr string `json:"issue_priority" db:"-"`
	StatusCode       int    `json:"-"`
	Status           string `json:"status" db:"-"`
	ReflectDate      int64  `json:"-"`
	ReflectDateStr   string `json:"ReflectDate,omitempty" db:"-"`
}
