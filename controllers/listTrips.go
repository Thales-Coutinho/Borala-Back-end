package controllers

import (
	"github.com/Thales-Coutinho/Borala-Back-end/models"
	"github.com/gin-gonic/gin"
)

func ListTrips(c *gin.Context) {
	departureCity := c.Request.URL.Query().Get("departureCity")
	destinationCity := c.Request.URL.Query().Get("destinationCity")
	day := c.Request.URL.Query().Get("day")

	trips := []models.Trip{
		{Id: "1", Name: "Gustavo", Score: "2,1", Hour: "08:30", Local: "ponte metalica", Price: "22,50", DepartureCity: "Santos", DestinationCity: "Praia Grande", Day: "2024-02-29"},
		{Id: "2", Name: "Carmen", Score: "5,0", Hour: "12:30", Local: "em frente a prefeitura", Price: "30,00", DepartureCity: "Santos", DestinationCity: "Praia Grande", Day: "2024-02-29"},
		{Id: "3", Name: "João", Score: "4,3", Hour: "17:40", Local: "Pasteko", Price: "30,00", DepartureCity: "Santos", DestinationCity: "Guaruja", Day: "2024-02-29"},
	}

	filteredTrips := []models.Trip{}
	for _, trip := range trips {
		if trip.DepartureCity == departureCity && trip.DestinationCity == destinationCity && trip.Day == day {
			filteredTrips = append(filteredTrips, trip)
		}
	}

	c.JSON(200, filteredTrips)
}
