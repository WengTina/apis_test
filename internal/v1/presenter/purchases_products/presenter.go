package purchases_products

import (
	"eirc.app/internal/v1/resolver/purchases_products"
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
	PurchasesProductsResolver purchases_products.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		PurchasesProductsResolver: purchases_products.New(db),
	}
}
