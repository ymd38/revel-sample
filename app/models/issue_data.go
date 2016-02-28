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
	Priority   int    `json:"priority"`
	Status     int    `json:"status"`
	Limit      int64  `json:"-"`
	LimitStr   string `json:"limit" db:"-"`
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
	if v.HasErrors() {
		return
	}

	v.Check(
		issuedata.Source,
		revel.Required{},
		revel.MaxSize{1024},
		revel.MinSize{1},
	).Message("source is validate error")
	if v.HasErrors() {
		return
	}

	v.Check(
		issuedata.Detail,
		revel.Required{},
		revel.MaxSize{5120},
		revel.MinSize{1},
	).Message("detail is validate error")
	if v.HasErrors() {
		return
	}

	v.Check(
		issuedata.Priority,
		revel.Required{},
	).Message("priority is validate error")
	if v.HasErrors() {
		return
	}

	v.Match(strconv.Itoa(issuedata.Status), regexp.MustCompile(`\d{1}`)).Message("status is validate error")
	if v.HasErrors() {
		return
	}
	v.Match(issuedata.LimitStr, regexp.MustCompile(`\d{8}`)).Message("limit is validate error")
	if v.HasErrors() {
		return
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
