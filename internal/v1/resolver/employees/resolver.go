package employees

import (
	"eirc.app/internal/v1/service/employees"
	model "eirc.app/internal/v1/structure/employees"
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
	EmployeesService employees.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		EmployeesService: employees.New(db),
	}
}
