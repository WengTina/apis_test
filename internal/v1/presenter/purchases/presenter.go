package purchases

import (
	"eirc.app/internal/v1/resolver/purchases"
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
	PurchasesResolver purchases.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		PurchasesResolver: purchases.New(db),
	}
}
