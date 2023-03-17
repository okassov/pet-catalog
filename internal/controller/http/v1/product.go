package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/okassov/pet-catalog/internal/usecase"
	"github.com/okassov/pet-catalog/pkg/logger"
)

type productRoutes struct {
	a usecase.Catalog
	l logger.LoggerInterface
}

func newProductRoutes(handler *gin.RouterGroup, c usecase.Catalog, l logger.LoggerInterface) {

	r := &productRoutes{c, l}

	h := handler.Group("/")
	{
		h.GET("product/:id", r.GetProduct)
		h.GET("products", r.ListProducts)
		h.POST("product/:id", r.CreateProduct)
		h.PUT("product/:id", r.UpdateProduct)
		h.DELETE("product/:id", r.DeleteProduct)
	}
}

//	@Summary		Get Product
//	@Description	Return Single Product
//	@ID				GetProduct
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SignUpResponse
//	@Failure		500	{object}	response
//	@Router			/product/:id [get]
func (r *productRoutes) GetProduct(c *gin.Context) {

}

//	@Summary		List Products
//	@Description	Return Multiple Products
//	@ID				ListProduct
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SignUpResponse
//	@Failure		500	{object}	response
//	@Router			/products [get]
func (r *productRoutes) ListProducts(c *gin.Context) {

}

//	@Summary		Create Product
//	@Description	Create Product
//	@ID				CreateProduct
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SignUpResponse
//	@Failure		500	{object}	response
//	@Router			/product/:id [post]
func (r *productRoutes) CreateProduct(c *gin.Context) {

}

//	@Summary		Update Product
//	@Description	Update Product
//	@ID				UpdateProduct
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SignUpResponse
//	@Failure		500	{object}	response
//	@Router			/product/:id [put]
func (r *productRoutes) UpdateProduct(c *gin.Context) {

}

//	@Summary		Delete Product
//	@Description	Delete Product
//	@ID				DeleteProduct
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SignUpResponse
//	@Failure		500	{object}	response
//	@Router			/products [delete]
func (r *productRoutes) DeleteProduct(c *gin.Context) {

}
