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
	/*
		condition := " where s.serviceid=" + strconv.Itoa(serviceid)
		if status != "" {
			c.App.Validation.Match(status, regexp.MustCompile("[0-9]"))
			if c.App.Validation.HasErrors() {
				return c.App.RenderJson(&ErrorResponse{ERR_VALIDATE, ErrorMessage(ERR_VALIDATE)})
			}
			condition += " and s.status=" + status
		}

		sql := SERVICE_ISSUE_SQL + condition
		rows, err := controllers.Dbm.Select(models.ServiceIssueView{}, sql)
		if err != nil {
			panic(err)
		}

		issues := make([]models.ServiceIssueView, len(rows))
		cnt := 0
		for _, row := range rows {
			issue := row.(*models.ServiceIssueView)
			issues[cnt].IssueId = issue.IssueId
			issues[cnt].IssueTitle = issue.IssueTitle
			issues[cnt].IssuePriorityStr = util.GetPriority(issue.IssuePriority)
			issues[cnt].StatusCode = issue.StatusCode
			issues[cnt].Status = util.GetStatus(issue.StatusCode)
			if issue.ReflectDate > 0 {
				issues[cnt].ReflectDateStr = util.UnitTimeToDayString(issue.ReflectDate)
			}

			cnt++
		}

		return c.App.RenderJson(&Response{OK, issues})
	*/
	return nil
}
