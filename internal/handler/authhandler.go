package handler

import (
	"atm/internal/models"
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
	creds := models.User{}

	if err := c.Bind(&creds); err != nil {
		h.errorLogger.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user, err := h.service.Auth.CheckUserCreds(creds)
	if err != nil {
		h.errorLogger.Print(err)
		if err == models.ErrIncorrectPassword {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrNoUser)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// creating session and generating jwt token

	token, err := h.service.Auth.Login(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]string{
		"username": user.Username,
		"token":    token,
	}

	h.infoLogger.Print("Successfully logged in")

	return c.JSON(http.StatusCreated, response)
}
