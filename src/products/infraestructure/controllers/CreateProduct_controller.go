package controllers

import (
    "demo/src/products/application"
    "demo/src/products/domain/entities"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

type CreateProductController struct {
    createProductUseCase *application.CreateProduct
}

func NewCreateProductController(createUseCase *application.CreateProduct) *CreateProductController {
    return &CreateProductController{createProductUseCase: createUseCase}
}

func (pc *CreateProductController) CreateProduct(c *gin.Context) {
    var product entities.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("Create Product")
    createdProduct, err := pc.createProductUseCase.Execute(product.Name, product.Price, product.Descripcion, product.Stock, product.User_id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    response := gin.H{
        "data": gin.H{
            "type": "products",
            "id":   createdProduct.ID,
            "attributes": gin.H{
                "name":  createdProduct.Name,
                "price": createdProduct.Price,
                "descripcion": createdProduct.Descripcion,
                "stock": createdProduct.Stock,
            },
            "relationships": gin.H{
                "user_id": createdProduct.User_id,
            },
        },
    }
    c.JSON(http.StatusCreated, response)
}

