package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *Handler) Routes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// home
	e.GET("/api", h.IndexHandler)

	// auth
	e.POST("/api/register", h.register)
	e.POST("/api/login", h.login)

	return e
}
