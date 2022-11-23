package main

import "pratice/web_framework/framework"

func SubjectAddController(c *framework.Context) {
	c.Json(200, "ok, SubjectAddController")
}

func SubjectListController(c *framework.Context) {
	c.Json(200, "ok, SubjectListController")
}

func SubjectDelController(c *framework.Context) {
	c.Json(200, "ok, SubjectDelController")
}

func SubjectUpdateController(c *framework.Context) {
	c.Json(200, "ok, SubjectUpdateController")
}

func SubjectGetController(c *framework.Context) {
	c.Json(200, "ok, SubjectGetController")
}

func SubjectNameController(c *framework.Context) {
	c.Json(200, "ok, SubjectNameController")
}
