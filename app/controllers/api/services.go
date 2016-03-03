package controllers

import (
	. "security-cop/app/models"

	"github.com/revel/revel"
)

type ApiServices struct {
	ApiController
	Service
}

func (c *ApiServices) Create() revel.Result {
	service_data := &ServiceData{}
	if err := BindJsonParams(c.App.Request.Body, service_data); err != nil {
		return c.App.RenderJson(&ErrorResponse{ERR_VALIDATE, ErrorMessage(ERR_VALIDATE)})
	}

	err := c.Service.Create(service_data)
	if err != nil {
		return c.App.RenderJson(&ErrorResponse{ERR_VALIDATE, ErrorMessage(ERR_FATAL)})
	}

	return c.App.RenderJson(&Response{OK, service_data})
}

//respones is detail of service by id.
func (c *ApiServices) Show(id int) revel.Result {
	service_list := c.Service.GetByID(id)
	return c.App.RenderJson(&Response{OK, service_list})
}

//respones is list of service.
func (c *ApiServices) List() revel.Result {
	service_list := c.Service.GetList()
	return c.App.RenderJson(&Response{OK, service_list})
}

func (c *ApiServices) Relation(id int) revel.Result {
	service_issue := new(ServiceIssue)
	service_issue.CreateByServiceID(id)
	return nil
}
