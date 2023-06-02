package purchases_products

import (
	"eirc.app/internal/v1/service/purchases_products"
	model "eirc.app/internal/v1/structure/purchases_products"
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
	PurchasesProductsService purchases_products.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		PurchasesProductsService: purchases_products.New(db),
	}
}
