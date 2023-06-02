package purchases

import (
	"eirc.app/internal/v1/service/purchases"
	model "eirc.app/internal/v1/structure/purchases"
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
	PurchasesService purchases.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		PurchasesService: purchases.New(db),
	}
}
