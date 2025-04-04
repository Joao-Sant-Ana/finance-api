package routes

import (
	"database/sql"
	"net/http"

	"api.finance.com/controllers/user"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string
}

func UserRoutes(rg *gin.RouterGroup, db *sql.DB) {
	userRoutes := rg.Group("/user")
	{
		userRoutes.POST("/", func(c *gin.Context) {
			var userData user.User
			
			if err := c.ShouldBindJSON(&userData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			if err := user.CreateUser(db, &userData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"message": "User created succesfully"})
		})
	}
}