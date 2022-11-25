package main

import "pratice/web_framework/framework/gin"

func UserLoginController(c *gin.Context) {
	c.JSON(200, "ok, UserLoginController")
}
