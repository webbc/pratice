package main

import (
	"pratice/web_framework/framework/gin"
	"pratice/web_framework/provider/demo"
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
	//time.Sleep(time.Second * 10)
	//id := c.Param("id")
	//c.JSON(200, "ok, SubjectGetController == "+id)
	// 获取demo服务实例
	demoService := c.MustMake(demo.Key).(demo.Service)

	// 调用服务实例的方法
	foo := demoService.GetFoo()

	// 输出结果
	c.ISetOkStatus().IJson(foo)
}

func SubjectNameController(c *gin.Context) {
	c.JSON(200, "ok, SubjectNameController")
}
