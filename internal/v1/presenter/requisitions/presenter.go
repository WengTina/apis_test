package requisitions

import (
	"eirc.app/internal/v1/resolver/requisitions"
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
	RequisitionsResolver requisitions.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		RequisitionsResolver: requisitions.New(db),
	}
}
