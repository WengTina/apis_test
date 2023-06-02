package departments

import (
	"eirc.app/internal/v1/service/departments"
	model "eirc.app/internal/v1/structure/departments"
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
	DepartmentsService departments.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		DepartmentsService: departments.New(db),
	}
}
