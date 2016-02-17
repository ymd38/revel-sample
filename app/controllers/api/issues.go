package controllers

import (
	"security-cop/app/models"
	"strconv"
	"sync"

	"github.com/revel/revel"
)

type ApiIssues struct {
	ApiController
	models.Issue
}

//insert issue data
func (c *ApiIssues) Create() revel.Result {
	issue_data := &models.IssueData{}
	if err := models.BindJsonParams(c.App.Request.Body, issue_data); err != nil {
		return c.App.RenderJson(&ErrorResponse{ERR_VALIDATE, ErrorMessage(ERR_VALIDATE)})
	}

	//issue := &models.Issue{}
	//err := issue.Create(issue_data)
	err := c.Issue.Create(issue_data)
	if err != nil {
		return c.App.RenderJson(&ErrorResponse{ERR_VALIDATE, ErrorMessage(ERR_FATAL)})
	}

	return c.App.RenderJson(&Response{OK, issue_data})
}

//update issue data
func (c *ApiIssues) Update(id int) revel.Result {
	return nil
}

//respones is detail of issue by id.
func (c *ApiIssues) Show(id int) revel.Result {
	issue_list := c.Issue.GetIssueList("where id=" + strconv.Itoa(id))
	return c.App.RenderJson(&Response{OK, issue_list})
}

//respones is list of issue.
func (c *ApiIssues) List(q string) revel.Result {
	issue_list := c.Issue.GetIssueList("")
	return c.App.RenderJson(&Response{OK, issue_list})
}

//list issues of service
func (c *ApiIssues) Service(serviceid int, status string) revel.Result {
	issue_list := c.Issue.GetServiceIssueList(serviceid, status)
	return c.App.RenderJson(&Response{OK, issue_list})
}

func (c *ApiIssues) Relation(issueid int) revel.Result {
	var wg sync.WaitGroup
	service_list := c.Issue.GetCreateTarget(issueid)
	serviceissue := &models.ServiceIssue{}
	for i := 0; i < len(service_list); i++ {
		wg.Add(1)
		go func(iid int, sid int) {
			defer wg.Done()
			si_data := &models.ServiceIssueData{Id: 0, IssueId: iid, ServiceId: sid}
			serviceissue.Create(si_data)
		}(issueid, service_list[i].Id)
	}
	wg.Wait()
	return nil
}
