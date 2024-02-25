package controllers

import (
	"github.com/Thales-Coutinho/Borala-Back-end/models"
	"github.com/gin-gonic/gin"
)

func ListCities(c *gin.Context) {
	cities := []models.City{
		{Id: "Guaruja", Name: "Guaruja"},
		{Id: "Santos", Name: "Santos"},
		{Id: "Praia Grande", Name: "Praia Grande"},
	}

	c.JSON(200, cities)
}
