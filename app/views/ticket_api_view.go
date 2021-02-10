package views

import (
	"net/http"
	"focus/app/models"

	"github.com/labstack/echo/v4"
)

type H map[string]interface{}

// GetTickets is ...
func GetTickets(c echo.Context) error {
	// Connection handling of the database
	db := models.InitDb()
	defer db.Close()

	var tickets []models.Ticket
	// SELECT * FROM tickets
	db.Find(&tickets, []int{4, 6})

	// Display JSON result
	out := H{"tickets": tickets}
	return c.JSON(http.StatusOK, out)
}
