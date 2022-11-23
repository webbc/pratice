package main

import (
	"pratice/web_framework/framework"
)

func UserLoginController(c *framework.Context) {
	c.Json(200, "ok, UserLoginController")
}
