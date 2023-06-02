package requisitions

import (
	requisitions "eirc.app/internal/v1/service/requisition"
	model "eirc.app/internal/v1/structure/requisitions"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	RequisitionsService requisitions.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		RequisitionsService: requisitions.New(db),
	}
}
