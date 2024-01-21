package handler

import (
	"atm/internal/models"
	"atm/internal/models/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) register(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		h.errLogger.Print(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	result, err := h.service.Auth.CreateNewUser(user)
	if err != nil {
		h.errLogger.Print(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	if result["status"] == "fail" {
		switch result["message"] {
		case auth.ErrInvalidEmail.Error(), auth.ErrInvalidPassword.Error():
			return c.JSON(http.StatusBadRequest, result["message"])
		case auth.ErrUsernameAlreadyTaken.Error():
			return c.JSON(http.StatusConflict, result["message"])
		case auth.ErrEmailAlreadyTaken.Error():
			return c.JSON(http.StatusConflict, result["message"])
		default:
			return c.JSON(http.StatusInternalServerError, "unknown error")
		}
	}

	h.infoLogger.Print("successfully created new user")
	return c.JSON(http.StatusOK, echo.Map{"message": "successfully created"})
}

func (h *Handler) login(c echo.Context) error {
	return nil
}
