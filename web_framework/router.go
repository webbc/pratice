package main

import (
	"pratice/web_framework/framework/gin"
	"pratice/web_framework/framework/middleware"
)

func registerRouter(core *gin.Engine) {

	core.Use(middleware.Test1(), middleware.Test2())

	// 静态路由+HTTP方法匹配
	core.GET("/user/login", UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", SubjectGetController)
		subjectApi.GET("/list/all", middleware.Test4(), SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}
	}
}
