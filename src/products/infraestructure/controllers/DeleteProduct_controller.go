package controllers

import (
    "demo/src/products/application"
    _"demo/src/products/domain/entities"
    _"fmt"
    "net/http"
	"strconv"
    "github.com/gin-gonic/gin"
)


type DeleteProductController struct {
	deleteProductUseCase *application.DeleteProduct
}
func NewDeleteProductController(deleteUseCase *application.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{deleteProductUseCase: deleteUseCase}
}

func (dp *DeleteProductController) DeleteProduct(g *gin.Context) {
	idParam := g.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid product ID"})
        return
    }
	idDelete := int32(id)
	
	if err2 := dp.deleteProductUseCase.Execute(idDelete); err2!= nil {
        g.JSON(http.StatusNotFound, gin.H{
				"detail": err2.Error(),
				"type": "products",    
		})
        return
    }
	response := gin.H{
        "data": gin.H{
            "type": "comments", 
			"id": idParam,
			"message":"Producto eliminado con éxito",
            
        },
    }
	g.JSON(http.StatusOK, response)
}

	



