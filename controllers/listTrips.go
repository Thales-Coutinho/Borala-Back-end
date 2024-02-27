package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/Thales-Coutinho/Borala-Back-end/db"
	"github.com/Thales-Coutinho/Borala-Back-end/models"
)

func ListTrips(c *gin.Context) {
	departureCity := c.Request.URL.Query().Get("departureCity")
	destinationCity := c.Request.URL.Query().Get("destinationCity")
	day := c.Request.URL.Query().Get("day")

	db := db.DataBaseConection()
	query := "SELECT t.id, d.name, d.score, t.hour, t.local, t.price, t.DepartureCity, t.DestinationCity, t.day FROM trips t JOIN drivers d ON t.driverCPF = d.CPF WHERE t.DepartureCity = ? AND t.DestinationCity = ? AND t.day = ?"
	selectTrips, err := db.Query(query, departureCity, destinationCity, day)
	if err != nil {
		panic(err.Error())
	}

	trips := []models.Trip{}
	for selectTrips.Next() {
		trip := models.Trip{}
		selectTrips.Scan(&trip.Id, &trip.Name, &trip.Score, &trip.Hour, &trip.Local, &trip.Price, &trip.DepartureCity, &trip.DestinationCity, &trip.Day)
		trips = append(trips, trip)
	}
	c.JSON(200, trips)
}
