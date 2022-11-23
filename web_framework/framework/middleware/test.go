package middleware

import (
	"fmt"
	"pratice/web_framework/framework"
)

func Test1() framework.Controller {
	return func(ctx *framework.Context) {
		fmt.Println("test1 middleware start")
		ctx.Next()
		fmt.Println("test1 middleware end")
	}
}

func Test2() framework.Controller {
	return func(ctx *framework.Context) {
		fmt.Println("test2 middleware start")
		ctx.Next()
		fmt.Println("test2 middleware end")
	}
}

func Test3() framework.Controller {
	return func(ctx *framework.Context) {
		fmt.Println("test3 middleware start")
		ctx.Next()
		fmt.Println("test3 middleware end")
	}
}

func Test4() framework.Controller {
	return func(ctx *framework.Context) {
		fmt.Println("test4 middleware start")
		ctx.Next()
		fmt.Println("test4 middleware end")
	}
}
