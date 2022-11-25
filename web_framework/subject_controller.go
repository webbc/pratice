package main

import (
	"pratice/web_framework/framework/gin"
	"time"
)

func SubjectListController(c *gin.Context) {
	c.JSON(200, "ok, SubjectListController")
}

func SubjectDelController(c *gin.Context) {
	c.JSON(200, "ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.JSON(200, "ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	time.Sleep(time.Second * 10)
	id := c.Param("id")
	c.JSON(200, "ok, SubjectGetController == "+id)
}

func SubjectNameController(c *gin.Context) {
	c.JSON(200, "ok, SubjectNameController")
}
