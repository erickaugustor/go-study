package controllers

import (
	"echo-api/src/utils"

	"net/http"

	"github.com/badoux/checkmail"
	"github.com/labstack/echo/v4"
)

func IsEmailValid(email string) bool {
	if has_error := checkmail.ValidateFormat(email); has_error != nil {
		return false
	}

	return true
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if IsEmailValid(email) != true {
		return c.String(http.StatusInternalServerError, "Error")
	}

	correct_password := "test"

	if password != correct_password {
		return c.String(http.StatusUnauthorized, "Not Logged")
	}

	token, _ := utils.CreateToken(email)

	return c.String(http.StatusOK, token)
}
