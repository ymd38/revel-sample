package controllers

import "github.com/revel/revel"

func init() {
	revel.OnAppStart(InitDB)
	//revel.InterceptMethod(App.Auth, revel.BEFORE)
}
