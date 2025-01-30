package controllers

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type UpdateController struct {
	updateProductUseCase *application.UpdateProduct
}

func NewUpdateProductController(updateUseCase *application.UpdateProduct) *UpdateController {
	return &UpdateController{updateProductUseCase: updateUseCase}
}

func (up *UpdateController) UpdateProduct(g *gin.Context) {
	var productToUpdate entities.UpdateProduct
	idParam := g.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid product ID"})
        return
    }
	idUpdate := int32(id)
	
	if err := g.ShouldBindJSON(&productToUpdate); err != nil {
        g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
        return
    }
	if err2 := up.updateProductUseCase.Execute(idUpdate, productToUpdate.Name, productToUpdate.Price, productToUpdate.Descripcion, productToUpdate.Stock); err2!= nil {
		g.JSON(http.StatusNotFound, gin.H{
                "detail": err2.Error(),
                "type": "products",    
        })
        return
    }
    response := gin.H{
		"data": gin.H{
			    "type": "products",
                "id":   idUpdate,
                "attributes": gin.H{
                    "name":  productToUpdate.Name,
                    "price": productToUpdate.Price,
                    "descripcion": productToUpdate.Descripcion,
                    "stock": productToUpdate.Stock,
                },
        },
    }
    g.JSON(http.StatusOK, response)

}