package routes

import (
	"handmedown-backend/src/controllers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetRoutes mengatur semua rute
func SetRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("CLIENT_ORIGIN")}
	route.Use(cors.New(config))

	route.POST("/login", controllers.LoginHandler)
	route.POST("/register", controllers.RegisterHandler)
	route.GET("/products", controllers.GetAllProducts)

	route.GET("/users", controllers.GetAllUser)

	// route.GET("/profile", middleware.AuthMiddleware(), controller)

	return route
}
