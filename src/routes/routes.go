package routes

import (
	"handmedown-backend/src/controllers"
	"handmedown-backend/src/middleware"
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
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	route.Use(cors.New(config))

	// Login Register routes
	route.POST("/login", controllers.LoginHandler)
	route.POST("/register", controllers.RegisterHandler)

	// Product routes
	route.GET("/products", controllers.GetAllProducts)
	route.GET("/product-details/:id", controllers.GetProductDetail)
	route.GET("/user-products/:id", controllers.GetUserProducts)

	// Cart routes
	route.POST("/cart", middleware.AuthorizationMiddleware(), controllers.AddToCart)
	route.GET("/cart", middleware.AuthorizationMiddleware(), controllers.GetCart)
	route.DELETE("/cart", middleware.AuthorizationMiddleware(), controllers.DeleteCartItem)

	// Profile routes
	route.GET("/profile", middleware.AuthorizationMiddleware(), controllers.GetProfile)
	route.PATCH("/profile", middleware.AuthorizationMiddleware(), controllers.UpdateProfile)
	route.GET("/profile/:id", controllers.GetUserProfile)

	route.GET("/users", controllers.GetAllUser)

	return route
}
