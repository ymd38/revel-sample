package controllers

import (
	"security-cop/app/models"
	"strconv"

	"github.com/revel/revel"
)

type ApiIssues struct {
	ApiController
	models.Issue
}

//insert issue data
func (c *ApiIssues) Create() revel.Result {
	issue_data := &models.IssueData{}
	c.BindParams(issue_data)

	issue_data.Validate(c.App.Validation)
	if c.App.Validation.HasErrors() {
		return c.App.RenderJson(&ErrorResponse{ERR_VALIDATE, ErrorMessage(ERR_VALIDATE)})
	}

	/*  modelのinsertへ移行
	//gorp doesn't support time type. we use unix time on DB.
	issue.Created = time.Now().Unix()
	issue.Updated = time.Now().Unix()
	err := c.Txn.Insert(issue)
	if err != nil {
		panic(err)
	}

	return c.App.RenderJson(&Response{OK, issue})
	*/
	return nil
}

//update issue data
func (c *ApiIssues) Update(id int) revel.Result {
	issue_data := &models.IssueData{Id: id}
	c.BindParams(issue_data)

	issue_data.Validate(c.App.Validation)
	if c.App.Validation.HasErrors() {
		return c.App.RenderJson(&ErrorResponse{ERR_VALIDATE, ErrorMessage(ERR_VALIDATE)})
	}

	/* modelのupdateへ移行
	_, err := c.Txn.Update(issue)
	if err != nil {
		panic(err)
	}

	return c.App.RenderJson(&Response{OK, issue})
	*/
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
