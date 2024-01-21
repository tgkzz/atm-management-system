package handler

import (
	"auth/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (h *Handler) register(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		h.errorLogger.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := h.service.Auth.CreateNewUser(user); err != nil {
		h.errorLogger.Print(err)
		if err == models.ErrInvalidEmail || err == models.ErrInvalidPassword {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if strings.Contains(err.Error(), "pq: повторяющееся значение ключа нарушает ограничение уникальности \"users_username_key\"") {
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrUsernameAlreadyTaken)
		}
		if strings.Contains(err.Error(), "pq: повторяющееся значение ключа нарушает ограничение уникальности \"users_email_key\"") {
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrEmailAlreadyTaken)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.infoLogger.Print("Successfully created new user")
	return c.JSON(http.StatusCreated, user)
}

func (h *Handler) login(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
