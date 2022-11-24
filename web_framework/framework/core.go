package framework

import "net/http"

type Core struct {
	routers     map[string]*Tree
	middlewares []Controller // 全局中间件
}

func NewCore() *Core {
	return &Core{
		routers: map[string]*Tree{
			"GET":    NewTree(),
			"POST":   NewTree(),
			"PUT":    NewTree(),
			"DELETE": NewTree(),
		},
	}
}

func (c *Core) Get(router string, controller ...Controller) {
	controllers := append(c.middlewares, controller...)
	c.routers["GET"].AddRouter(router, controllers...)
}

func (c *Core) Post(router string, controller ...Controller) {
	controllers := append(c.middlewares, controller...)
	c.routers["POST"].AddRouter(router, controllers...)
}

func (c *Core) Put(router string, controller ...Controller) {
	controllers := append(c.middlewares, controller...)
	c.routers["PUT"].AddRouter(router, controllers...)
}

func (c *Core) Delete(router string, controller ...Controller) {
	controllers := append(c.middlewares, controller...)
	c.routers["DELETE"].AddRouter(router, controllers...)
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	uri := r.RequestURI

	ctx := NewContext(w, r)

	// 获取controller
	controllers := c.GetController(ctx, method, uri)

	// 设置controller
	ctx.SetControllers(controllers)

	// 调用Next
	ctx.Next()
}

// 获取路由
func (c *Core) GetController(ctx *Context, method, uri string) []Controller {

	// method的tree没有找到
	methodTree := c.routers[method]
	if methodTree == nil {
		return nil
	}

	// 没有找到对应的路由节点
	matchNode := methodTree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}

	// 解析路由参数
	params := methodTree.ParseParam(uri, matchNode)

	// 设置路由参数
	ctx.SetParam(params)

	return matchNode.controllers
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(prefix, c)
}

// 新增全局中间件
func (c *Core) Use(middlewares ...Controller) {
	c.middlewares = append(c.middlewares, middlewares...)
}
