package middlewares

import (
	"github.com/agung96tm/miblog/lib"
	"github.com/labstack/echo/v4/middleware"
)

type GZipMiddleware struct {
	handler lib.HttpHandler
}

func NewGZipMiddleware(handler lib.HttpHandler) GZipMiddleware {
	return GZipMiddleware{
		handler: handler,
	}
}

func (m GZipMiddleware) Setup() {
	m.handler.Engine.Use(middleware.Gzip())
}
