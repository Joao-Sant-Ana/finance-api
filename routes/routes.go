package routes

import (
	"github.com/gin-gonic/gin"
	"api.finance.com/config"
	"database/sql"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config, db *sql.DB) {
	version := r.Group("/" + cfg.Version)

	UserRoutes(version, db)
}
