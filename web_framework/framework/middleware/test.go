package middleware

import (
	"fmt"
	"pratice/web_framework/framework/gin"
)

func Test1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("test1 middleware start")
		ctx.Next()
		fmt.Println("test1 middleware end")
	}
}

func Test2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("test2 middleware start")
		ctx.Next()
		fmt.Println("test2 middleware end")
	}
}

func Test3() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("test3 middleware start")
		ctx.Next()
		fmt.Println("test3 middleware end")
	}
}

func Test4() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("test4 middleware start")
		ctx.Next()
		fmt.Println("test4 middleware end")
	}
}
