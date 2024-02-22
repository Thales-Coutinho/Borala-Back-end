package controllers

import "github.com/gin-gonic/gin"

func ListCities(c *gin.Context) {
	cities := []map[string]string{
		{
			"id":   "Guaruja",
			"name": "Guaruja",
		},
		{
			"id":   "Santos",
			"name": "Santos",
		},
		{
			"id":   "Praia Grande",
			"name": "Praia Grande",
		},
	}

	c.JSON(200, gin.H{
		"cities": cities,
	})
}
