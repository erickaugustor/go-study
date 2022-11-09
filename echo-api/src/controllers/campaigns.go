package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Campaign struct {
	Id uint8 `json:"id"`
	Name string `json:"name"`
}

func GetCampaigns(c echo.Context) error {
	campaign := Campaign{
		Name:  "Campaign Test",
		Id: 1,
	}

	return c.JSON(http.StatusOK, campaign)
}
