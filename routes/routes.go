package routes

import (
	"github.com/Thales-Coutinho/Borala-Back-end/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/cities", controllers.ListCities)
	r.GET("/trips", controllers.ListTrips)
	r.Run()
}
