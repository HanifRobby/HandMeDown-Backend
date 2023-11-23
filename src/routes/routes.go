package routes

import (
	"handmedown-backend/src/controllers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// func Routes() {
// 	route := gin.Default()

// 	config := cors.DefaultConfig()
// 	config.AllowOrigins = []string{os.Getenv("CLIENT_ORIGIN")}
// 	route.Use(cors.New(config))

// 	route.GET("handmedown/products", controllers.GetAllProducts)

// }

// SetRoutes mengatur semua rute
func SetRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("CLIENT_ORIGIN")}
	route.Use(cors.New(config))

	// Menambahkan middleware yang menginjeksi DB ke setiap pengendali
	route.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	route.GET("/handmedown/products", controllers.GetAllProducts)

	return route
}
