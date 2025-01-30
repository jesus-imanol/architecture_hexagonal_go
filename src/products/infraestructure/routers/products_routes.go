package routers

import (
	"demo/src/products/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, createProductController *controllers.CreateProductController,listProductController *controllers.ListProductController, deleteProductController *controllers.DeleteProductController, updateProductController *controllers.UpdateController) {
	v1 := r.Group("/v1/products")
	{
		v1.POST("/create", createProductController.CreateProduct)
		v1.GET("/", listProductController.GetAllProducts)
		v1.DELETE("/:id", deleteProductController.DeleteProduct)
		v1.PUT("/:id",updateProductController.UpdateProduct)
	}
}

/*
GET /products ---> handler
POST /products --> handler (CreateProduct_controller)
PUT /products ---> handler
DELETE /products-> handler
*/
