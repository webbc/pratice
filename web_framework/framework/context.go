package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Context struct {
	writer http.ResponseWriter
	req    *http.Request
	ctx    context.Context
}

func NewContext(writer http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		writer: writer,
		req:    req,
		ctx:    req.Context(),
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
