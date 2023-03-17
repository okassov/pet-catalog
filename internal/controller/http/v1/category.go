package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/okassov/pet-catalog/internal/usecase"
	"github.com/okassov/pet-catalog/pkg/logger"
)

type categoryRoutes struct {
	a usecase.Catalog
	l logger.LoggerInterface
}

func newCategoryRoutes(handler *gin.RouterGroup, c usecase.Catalog, l logger.LoggerInterface) {

	r := &categoryRoutes{c, l}

	h := handler.Group("/")
	{
		h.GET("category/:id", r.GetCategory)
		h.GET("categorys", r.ListCategories)
		h.POST("category/:id", r.CreateCategory)
		h.PUT("category/:id", r.UpdateCategory)
		h.DELETE("category/:id", r.DeleteCategory)
	}
}

//	@Summary		Get Category
//	@Description	Return Single Category
//	@ID				GetCategory
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SignUpResponse
//	@Failure		500	{object}	response
//	@Router			/Category/:id [get]
func (r *categoryRoutes) GetCategory(c *gin.Context) {

}

//	@Summary		List Categorys
//	@Description	Return Multiple Categorys
//	@ID				ListCategory
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SignUpResponse
//	@Failure		500	{object}	response
//	@Router			/Categorys [get]
func (r *categoryRoutes) ListCategories(c *gin.Context) {

}

//	@Summary		Create Category
//	@Description	Create Category
//	@ID				CreateCategory
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SignUpResponse
//	@Failure		500	{object}	response
//	@Router			/Category/:id [post]
func (r *categoryRoutes) CreateCategory(c *gin.Context) {

}

//	@Summary		Update Category
//	@Description	Update Category
//	@ID				UpdateCategory
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SignUpResponse
//	@Failure		500	{object}	response
//	@Router			/Category/:id [put]
func (r *categoryRoutes) UpdateCategory(c *gin.Context) {

}

//	@Summary		Delete Category
//	@Description	Delete Category
//	@ID				DeleteCategory
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SignUpResponse
//	@Failure		500	{object}	response
//	@Router			/Categorys [delete]
func (r *categoryRoutes) DeleteCategory(c *gin.Context) {

}
