package framework

import "net/http"

type Core struct {
	//routers map[string]map[string]Controller
	routers map[string]*Tree
}

func NewCore() *Core {
	return &Core{
		//routers: map[string]map[string]Controller{
		//	"GET":  make(map[string]Controller),
		//	"POST": make(map[string]Controller),
		//},
		routers: map[string]*Tree{
			"GET":    NewTree(),
			"POST":   NewTree(),
			"PUT":    NewTree(),
			"DELETE": NewTree(),
		},
	}
}

func (c *Core) Get(router string, controller Controller) {
	//c.routers["GET"][router] = controller
	c.routers["GET"].AddRouter(router, controller)
}

func (c *Core) Post(router string, controller Controller) {
	//c.routers["POST"][router] = controller
	c.routers["POST"].AddRouter(router, controller)
}

func (c *Core) Put(router string, controller Controller) {
	//c.routers["GET"][router] = controller
	c.routers["PUT"].AddRouter(router, controller)
}

func (c *Core) Delete(router string, controller Controller) {
	//c.routers["POST"][router] = controller
	c.routers["DELETE"].AddRouter(router, controller)
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	uri := r.RequestURI

	ctx := NewContext(w, r)
	if controller := c.FindController(method, uri); controller != nil {
		controller(ctx)
		return
	}
	ctx.Json(404, "NOT FOUND")
}

func (c *Core) FindController(method, uri string) Controller {
	methodTree := c.routers[method]
	if methodTree == nil {
		return nil
	}
	return methodTree.FindController(uri)
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(prefix, c)
}
