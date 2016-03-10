package models

import (
	"errors"
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

func (issuedata *IssueData) Validate() error {
	var v revel.Validation
	v.Check(
		issuedata.Title,
		revel.Required{},
		revel.MaxSize{1024},
		revel.MinSize{1},
	)
	if v.HasErrors() {
		return errors.New("titel is validate error")
	}

	v.Check(
		issuedata.Source,
		revel.Required{},
		revel.MaxSize{1024},
		revel.MinSize{1},
	)
	if v.HasErrors() {
		return errors.New("source is validate error")
	}

	v.Check(
		issuedata.Detail,
		revel.Required{},
		revel.MaxSize{5120},
		revel.MinSize{1},
	)
	if v.HasErrors() {
		return errors.New("detail is validate error")
	}

	v.Match(strconv.Itoa(issuedata.Priority), regexp.MustCompile(`\d{1}`))
	if v.HasErrors() {
		return errors.New("priority is validate error")
	}

	v.Match(strconv.Itoa(issuedata.Status), regexp.MustCompile(`\d{1}`))
	if v.HasErrors() {
		return errors.New("status is validate error")
	}
	v.Match(issuedata.LimitStr, regexp.MustCompile(`\d{8}`))
	if v.HasErrors() {
		return errors.New("limit is validate error")
	}

	return nil
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
