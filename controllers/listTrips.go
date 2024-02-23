package controllers

import (
	"github.com/gin-gonic/gin"
)

func ListTrips(c *gin.Context) {
	departureCity := c.Request.URL.Query().Get("departureCity")
	destinationCity := c.Request.URL.Query().Get("destinationCity")
	day := c.Request.URL.Query().Get("day")

	trips := []map[string]string{
		{
			"id":              "1",
			"name":            "Gustavo",
			"score":           "2,1",
			"hour":            "08:30",
			"local":           "ponte metalica",
			"price":           "22,50",
			"departureCity":   "Santos",
			"destinationCity": "Praia Grande",
			"day":             "2024-02-29",
		},
		{
			"id":              "2",
			"name":            "Carmen",
			"score":           "5,0",
			"hour":            "12:30",
			"local":           "em frente a prefeitura",
			"price":           "30,00",
			"departureCity":   "Santos",
			"destinationCity": "Praia Grande",
			"day":             "2024-02-29",
		},
		{
			"id":              "3",
			"name":            "João",
			"score":           "4,3",
			"hour":            "17:40",
			"local":           "Pasteko",
			"price":           "30,00",
			"departureCity":   "Santos",
			"destinationCity": "Guaruja",
			"day":             "2024-02-29",
		},
	}

	filteredTrips := []map[string]string{}
	for _, trip := range trips {
		if trip["departureCity"] == departureCity && trip["destinationCity"] == destinationCity && trip["day"] == day {
			filteredTrips = append(filteredTrips, trip)
		}
	}

	c.JSON(200, gin.H{
		"trips": filteredTrips,
	})
}
