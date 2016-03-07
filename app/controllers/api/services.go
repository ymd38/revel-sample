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
	serviceData := &ServiceData{}
	if err := BindJsonParams(c.App.Request.Body, serviceData); err != nil {
		return c.App.RenderJsonP(CALLBACK, &ErrorResponse{ERR_VALIDATE, err.Error()})
	}

	err := c.Service.Create(serviceData)
	if err != nil {
		return c.App.RenderJsonP(CALLBACK, &ErrorResponse{ERR_FATAL, err.Error()})
	}

	return c.App.RenderJsonP(CALLBACK, &Response{OK, serviceData})
}

//respones is detail of service by id.
func (c *ApiServices) Show(id int) revel.Result {
	serviceList := c.Service.GetByID(id)
	return c.App.RenderJsonP(CALLBACK, &Response{OK, serviceList})
}

//respones is list of service.
func (c *ApiServices) List() revel.Result {
	serviceList := c.Service.GetList()
	return c.App.RenderJsonP(CALLBACK, &Response{OK, serviceList})
}

func (c *ApiServices) Relation(id int) revel.Result {
	serviceIssue := new(ServiceIssue)
	serviceIssue.CreateByServiceID(id)
	return c.App.RenderJsonP(CALLBACK, &Response{OK, nil})
}
