package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) IndexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}
