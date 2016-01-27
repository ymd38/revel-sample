package controllers

import (
	"fmt"

	"github.com/revel/revel"
)

func init() {
	fmt.Println("init.go:init")
	revel.OnAppStart(InitDB)
	//revel.InterceptMethod(App.Auth, revel.BEFORE)
}
