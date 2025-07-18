package routes

import (
	"github.com/gin-gonic/gin"
	"cadastro-clientes-api/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	api := r.Group("/api")
	{
		api.GET("/clientes", controllers.GetClientes)
		api.GET("/clientes/:id", controllers.GetClienteByID)
		api.POST("/clientes", controllers.CreateCliente)
		api.PUT("/clientes/:id", controllers.UpdateCliente)
		api.DELETE("/clientes/:id", controllers.DeleteCliente)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "API funcionando",
		})
	})

	return r
}
