package routes

import (
	"database/sql"
	"net/http"
	"api.finance.com/controllers/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup, db *sql.DB) {
	userRoutes := rg.Group("/user")
	{
		userRoutes.POST("/create", func(c *gin.Context) {
			var userData user.UserRegister

			if err := c.ShouldBindJSON(&userData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := user.CreateUser(db, &userData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"message": "User created succesfully"})
		})

		userRoutes.POST("/", func(c *gin.Context) {
			var userData user.UserLogin

			if err := c.ShouldBindJSON(&userData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			token, err := user.LoginUser(db, &userData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.SetCookie(
				"auth_token",
				token.Auth_token,
				3600,
				"/",
				"",
				false,
				false,
			)

			c.SetCookie(
				"refresh_token",
				token.Refresh_token,
				2.419e+6,
				"/",
				"",
				false,
				true,
			)

			c.Status(http.StatusNoContent)
		})
	}
}
