package api

import (
	"github.com/customer-experience/api-prueba/controllers"
	_ "github.com/customer-experience/api-prueba/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//SetupRouter - Definicion de las rutas de las aplicaciones
func SetupRouter(r *gin.Engine) {

	v1 := r.Group("/api/v1")
	{
		pedidos := v1.Group("/pedidos")
		{
			pedidos.POST("", controllers.CrearPedido)
			pedidos.GET(":id", controllers.BuscarPedido)
			pedidos.PATCH(":id", controllers.ModificarPedido)
		}
	}
	r.GET("/api/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // ../api/doc/index.html
}
