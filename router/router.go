package router

import (
	"jwt-go/controllers"
	"jwt-go/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.UserRegister)
		userGroup.POST("/login", controllers.UserLogin)
	}

	productGroup := r.Group("/products")
	{
		productGroup.Use(middlewares.Authentication())
		productGroup.GET("/", controllers.FindAllProducts)
		productGroup.POST("/", controllers.CreateProduct)
		productGroup.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productGroup.GET("/:productId", middlewares.ProductAuthorization(), controllers.FindProductByID)
		productGroup.DELETE("/:productId", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return r
}
