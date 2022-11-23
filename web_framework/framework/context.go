package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Context struct {
	writer      http.ResponseWriter
	req         *http.Request
	ctx         context.Context
	controllers []Controller // controllers
	index       int          // 当前处理到了那个handler
}

func NewContext(writer http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		writer: writer,
		req:    req,
		ctx:    req.Context(),
		index:  -1, // 为了+1之后默认是0，执行第一个controller
	}
}

// implement context
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Context) Err() error {
	return c.ctx.Err()
}

func (c *Context) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

// query
func (c *Context) QueryAll() map[string][]string {
	if c.req != nil {
		return c.req.URL.Query()
	}
	return map[string][]string{}
}

func (c *Context) QueryInt(key string, def int) int {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			intval, err := strconv.Atoi(vals[len-1])
			if err != nil {
				return def
			}
			return intval
		}
	}
	return def
}

func (c *Context) QueryString(key string, def string) string {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return vals[len-1]
		}
	}
	return def
}

func (c *Context) QueryArray(key string, def []string) []string {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

// writer
func (c *Context) Json(status int, data interface{}) {
	c.writer.WriteHeader(status)
	c.writer.Header().Set("Content-Type", "application/json")

	buf, err := json.Marshal(data)
	if err != nil {
		c.writer.WriteHeader(500)
		return
	}

	c.writer.Write(buf)
}

// 设置controllers
func (c *Context) SetControllers(controller []Controller) {
	c.controllers = controller
}

// 获取下一个controller
func (c *Context) Next() {
	c.index++
	if c.index < len(c.controllers) {
		controller := c.controllers[c.index]
		if controller != nil {
			controller(c)
		}
	}
}
