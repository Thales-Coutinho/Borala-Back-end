package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Thales-Coutinho/Borala-Back-end/db"
	"github.com/Thales-Coutinho/Borala-Back-end/models"
)

func ListCities(c *gin.Context) {
	db := db.DataBaseConection()
	query := "SELECT id, name FROM cities"
	selectCities, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error on ListTrips": err.Error(),
		})
		return
	}

	cities := []models.City{}
	for selectCities.Next() {
		city := models.City{}
		selectCities.Scan(&city.Id, &city.Name)
		cities = append(cities, city)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	c.JSON(http.StatusOK, cities)
}
