package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Routes() *echo.Echo {
	e := echo.New()

	mainGroup := e.Group("/api")

	// auth
	authGroup := mainGroup.Group("/auth")
	authGroup.POST("/register", h.register)
	authGroup.POST("/login", h.login)

	//some other services

	return e
}
