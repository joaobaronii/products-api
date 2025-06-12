package controller

import (
	"api-produtos-go/model"
	"api-produtos-go/usecase"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id_product := ctx.Param("product_id")
	if id_product == "" {
		response := model.Response{
			Message: "Product ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	productId, err := strconv.Atoi(id_product)
	if err != nil {
		response := model.Response{
			Message: "Product ID must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}


	product, err := p.productUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id_product := ctx.Param("product_id")
	if id_product == "" {
		response := model.Response{
			Message: "Product ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	productId, err := strconv.Atoi(id_product)
	if err != nil {
		response := model.Response{
			Message: "Product ID must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUseCase.DeleteProduct(productId)
	if err == sql.ErrNoRows {
		response := model.Response {
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	response := model.Response{
		Message: "Product deleted successfully",
	}
	ctx.JSON(http.StatusOK, response)
}