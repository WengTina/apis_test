package roles

import (
	"eirc.app/internal/v1/service/roles"
	model "eirc.app/internal/v1/structure/roles"
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
	RolesService roles.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		RolesService: roles.New(db),
	}
}
