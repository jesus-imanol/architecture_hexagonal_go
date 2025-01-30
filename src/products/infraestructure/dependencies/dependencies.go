package dependencies

import (
	"demo/src/products/application"
	_"demo/src/products/domain/entities"
	"demo/src/products/infraestructure/adapters"
	"demo/src/products/infraestructure/controllers"
	"demo/src/products/infraestructure/routers"

	"fmt"

	"github.com/gin-gonic/gin"
)

func InitProducts(r *gin.Engine) {
	// ! Crear producto
	fmt.Println("Initializing")
	ps, err := adapters.NewMySQL()
	if err != nil {
		panic(err)
	}

	createProductUseCase := application.NewCreateProduct(ps)
	createProduct_controller := controllers.NewCreateProductController(createProductUseCase)
	//listar productosssss
	listProductUseCase := application.NewListProduct(ps)
    listProduct_controller := controllers.NewListProductController(listProductUseCase)
	//delete productosssss
	deleteProductUseCase := application.NewDeleteProduct(ps)
	deleteProduct_controller := controllers.NewDeleteProductController(deleteProductUseCase)
	//update productossss
	updateProductUseCase := application.NewUpdateProduct(ps)
	updateProduct_controller := controllers.NewUpdateProductController(updateProductUseCase)

	routers.ProductRoutes(r, createProduct_controller, listProduct_controller,deleteProduct_controller,updateProduct_controller)

}
