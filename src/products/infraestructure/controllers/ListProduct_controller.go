package controllers

import (
    "demo/src/products/application"
    _"demo/src/products/domain/entities"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ListProductController struct {
    listProductUseCase *application.ListProduct
}

func NewListProductController(useCase *application.ListProduct) *ListProductController {
    return &ListProductController{listProductUseCase: useCase}
}

func (pc *ListProductController) GetAllProducts(c *gin.Context) {
    products, err := pc.listProductUseCase.Execute()
	fmt.Println("soy products",products)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var response []gin.H
    for _, product := range products {
        productResponse := gin.H{
            "type": "products",
            "id":   product.ID,
            "attributes": gin.H{
                "name":  product.Name,
                "price": product.Price,
                "descripcion": product.Descripcion,
                "stock": product.Stock,
            },
            "relationships": gin.H{
                "user_id": product.User_id,
            },
        }
        response = append(response, productResponse)
    }

    if len(products) > 0 {
        c.JSON(http.StatusOK, gin.H{"data": response})
    } else {
        fmt.Println("Products:", len(products))
      
        c.JSON(http.StatusOK, gin.H{
            "data": len(products),
            "message": "No se encontraron productos",
            "type": "products",
        })
    }
}
