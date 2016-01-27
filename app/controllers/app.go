package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

/*
func (c App) login(username, password string) revel.Result {
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
			if err == nil {
				c.Session["user"] = username

hashedpassword, _ := bcrypt.GenerateFromPassword(
	[]byte(password), bcrypt.DefaultCost)

//err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))

c.Session["user"] = username
c.Session.SetDefaultExpiration()
*/

func (c App) Auth() revel.Result {

	if username, ok := c.Session["user"]; ok {
		revel.INFO.Println(username)
		//return c.getUser(username)
		return nil
	}
	//c.RenderArgs["user"] = c.Session["user"]

	return c.RenderText("Please login first")
}
