package main

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var enforcer *casbin.Enforcer

// init 初始化
func init() {
	adapter, err := gormadapter.NewAdapter("mysql", "baochao:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}

	e, err := casbin.NewEnforcer("casbin/rbac_model.conf", adapter)
	if err != nil {
		log.Fatal(err)
	}
	enforcer = e

	e.LoadPolicy()
}

// GetPermissionsByUser 获取用户权限，用于前端展示
func GetPermissionsByUser(username string) [][]string {
	data, _ := enforcer.GetImplicitPermissionsForUser(username)
	return data
}

// CheckPermission 权限校验中间件
func CheckPermission(ctx *gin.Context) {
	// 获取角色，一般不会通过参数传递角色，而是根据用户当前的token解析得到用户的角色
	// 这里为了演示而简单处理
	role := ctx.Query("role")
	if role == "" {
		role = "anonymous"
	}

	// 根据角色,路由，请求方法校验权限
	hasPermission, err := enforcer.Enforce(role, ctx.Request.URL.Path, ctx.Request.Method)

	// 程序异常
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		ctx.Abort()
	}

	// 没有权限
	if !hasPermission {
		ctx.String(http.StatusUnauthorized, "StatusUnauthorized")
		ctx.Abort()
	}

	ctx.Next()
}

func main() {

	router := gin.Default()

	test := router.Group("test")
	{
		test.GET("/add", func(context *gin.Context) {
			enforcer.AddPolicy("member", "/article", "*")
			enforcer.SavePolicy()
			context.JSON(http.StatusOK, gin.H{
				"data": "ok",
			})
		})
	}

	// 根据用户
	router.GET("permission", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": GetPermissionsByUser(context.Query("user")),
		})
	})

	// 定义一组文章的接口
	article := router.Group("article").Use(CheckPermission)
	{
		// 列表查询
		article.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"data": "paginate",
			})
		})

		// 查看详情
		article.GET("/:id", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"data": "info:" + context.Param("id"),
			})
		})

		// 更新
		article.PUT("/:id", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"data": "update" + context.Param("id"),
			})
		})

		// 创建
		article.POST("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"data": "create",
			})
		})

		// 删除
		article.DELETE("/:id", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"data": "delete" + context.Param("id"),
			})
		})
	}

	router.Run(":8888")
}
