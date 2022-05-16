package handler

import (
	"net/http"
    
	"github.com/labstack/echo"
)

func GetHello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}
