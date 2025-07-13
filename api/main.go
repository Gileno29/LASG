package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"carreira-api/models"
)

func main() {
	router := gin.Default()

	router.POST("/logs", func(c *gin.Context) {
		var evento models.LogEvento

		if err := c.ShouldBindJSON(&evento); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Aqui você pode armazenar no banco ou processar o evento
		fmt.Printf("Evento recebido: %+v\n", evento)

		c.JSON(http.StatusOK, gin.H{
			"message": "Evento recebido com sucesso!",
			"user_id": evento.UserID,
		})
	})

	router.Run(":8080")
}

