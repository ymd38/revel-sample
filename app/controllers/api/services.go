package controllers

import (
	"security-cop/app/models"
	"strconv"

	"github.com/revel/revel"
)

type ApiServices struct {
	ApiController
	models.Service
}

//respones is detail of issue by id.
func (c *ApiServices) Show(id int) revel.Result {
	service_list := c.Service.GetServiceList("where id=" + strconv.Itoa(id))
	return c.App.RenderJson(&Response{OK, service_list})
}

//respones is list of service.
func (c *ApiServices) List() revel.Result {
	service_list := c.Service.GetServiceList("")
	return c.App.RenderJson(&Response{OK, service_list})
}
