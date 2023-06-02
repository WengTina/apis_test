package products

import (
	"eirc.app/internal/v1/resolver/products"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	ProductsResolver products.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		ProductsResolver: products.New(db),
	}
}
