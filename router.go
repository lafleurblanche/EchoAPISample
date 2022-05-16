package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lafleurblanche/echoapisample/handler"
)

func newRouter() *echo.Echo {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    
    api := e.Group("/api")
    
    api.GET("/hello", handler.GetHello) // GET /api/hello

    return e
}
