package views

import (
	"net/http"
	"focus/app/models"

	"github.com/labstack/echo/v4"
)

type H map[string]interface{}

// FindTickets is ... GET /tickets
func FindTickets(c echo.Context) error {
	var tickets []models.Ticket
	models.DB.Find(&tickets)

	return c.JSON(http.StatusOK, H{"data": tickets})
	// return c.JSON(http.StatusOK, nil)
}