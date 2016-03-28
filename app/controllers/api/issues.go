package controllers

import (
	"fmt"
	. "security-cop/app/models"

	"github.com/revel/revel"
)

type ApiIssues struct {
	ApiController
	Issue
}

//insert issue data
func (c *ApiIssues) Create() revel.Result {
	issue_data := IssueData{}
	fmt.Println(c.Request.Body)
	if err := BindJsonParams(c.Request.Body, &issue_data); err != nil {
		return c.Response(&ErrorResponse{ERR_VALIDATE, err.Error()})
	}

	err := c.Issue.Create(&issue_data)
	if err != nil {
		return c.Response(&ErrorResponse{ERR_FATAL, err.Error()})
	}

	return c.Response(&Response{OK, issue_data})
}

//update issue data
func (c *ApiIssues) Update(id int) revel.Result {
	return nil
}

//respones is detail of issue by id.
func (c *ApiIssues) Show(id int) revel.Result {
	issue_list := c.Issue.GetByID(id)
	return c.Response(&Response{OK, issue_list})
}

//respones is list of issue.
func (c *ApiIssues) List(q string) revel.Result {
	issue_list := c.Issue.GetList()
	return c.Response(&Response{OK, issue_list})
}

//list issues of service
func (c *ApiIssues) Service(serviceid int, status string) revel.Result {
	issue_list := c.Issue.GetServiceIssueList(serviceid, status)
	return c.Response(&Response{OK, issue_list})
}

func (c *ApiIssues) Relation(id int) revel.Result {
	service_issue := new(ServiceIssue)
	service_issue.CreateByIssueID(id)
	return nil
}
