package handler

import (
	"atm/internal/models"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) register(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		h.errLogger.Print(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	userData, err := json.Marshal(user)
	if err != nil {
		h.errLogger.Print(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	resp, err := http.Post("http://localhost:8181/auth/register", "application/json", bytes.NewBuffer(userData))
	if err != nil {
		h.errLogger.Print(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var respError map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&respError); err != nil {
			h.errLogger.Print(err)
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "error parsing response from authentication service"})
		}
		return c.JSON(resp.StatusCode, respError)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "successfully created"})
}

func (h *Handler) login(c echo.Context) error {
	return nil
}
