package infrastructure

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Setting up the routes
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "API is Running..."})
	})
	// routes.UserRoutes(httpRouter.Group("user"), db)
	port := os.Getenv("PORT")
	if port == "" {
		httpRouter.Run()
	} else {
		httpRouter.Run(port)
	}
}
