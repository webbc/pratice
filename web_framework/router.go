package main

import (
	"pratice/web_framework/framework"
	"pratice/web_framework/framework/middleware"
)

func registerRouter(core *framework.Core) {

	core.Use(middleware.Test1(), middleware.Test2())

	// 静态路由+HTTP方法匹配
	core.Get("/user/login", UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject").Use(middleware.Test3())
	{
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", middleware.Test4(), SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}
