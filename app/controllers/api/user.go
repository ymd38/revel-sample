package controllers

import "github.com/revel/revel"

type ApiUser struct {
	ApiController
}

func (c *ApiUser) Create() revel.Result {
	return nil
	/*
		user := &models.User{}
		c.BindParams(user)

		user.Validate(c.App.Validation)
		if c.App.Validation.HasErrors() {
			return c.App.RenderJson(&ErrorResponse{ERR_VALIDATE, ErrorMessage(ERR_VALIDATE)})
		}

		err := c.Txn.Insert(user)
		if err != nil {
			panic(err)
		}

		return c.App.RenderJson(&Response{OK, user})
	*/
}
