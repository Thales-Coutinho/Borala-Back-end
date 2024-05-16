package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Thales-Coutinho/Borala-Back-end/db"
	"github.com/Thales-Coutinho/Borala-Back-end/models"
)

func ListTrips(c *gin.Context) {

	var requestData struct {
		DepartureCity   string `json:"departure_city"`
		DestinationCity string `json:"destination_city"`
		Day             string `json:"day"`
	}

	c.BindJSON(&requestData)

	db := db.DataBaseConection()

	query := `SELECT t.id, d.name, d.score, t.hour, t.local, t.price, t.departure_city, t.destination_city, t.day 
				FROM trips t JOIN drivers d ON t.driver_cpf = d.cpf 
				WHERE t.departure_city = ? AND t.destination_city = ? AND t.day = ?`

	selectTrips, err := db.Query(query, requestData.DepartureCity, requestData.DestinationCity, requestData.Day)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error on ListTrips": err.Error(),
		})
		return
	}

	trips := []models.Trip{}
	for selectTrips.Next() {
		trip := models.Trip{}
		selectTrips.Scan(&trip.Id, &trip.Name, &trip.Score, &trip.Hour, &trip.Local, &trip.Price, &trip.DepartureCity, &trip.DestinationCity, &trip.Day)
		trips = append(trips, trip)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	c.JSON(http.StatusOK, trips)
}
