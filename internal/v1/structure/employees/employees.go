package employees

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {
	//員工ID
	EmployeesID string `gorm:"primaryKey;column:employees_id;uuid_generate_v4()type:UUID;" json:"employees_id,omitempty"`
	//員工名稱
	EmployeesName string `gorm:"column:employees_name;type:TEXT;not null;" json:"employees_name,omitempty"`
	//角色ID
	Roles string `gorm:"column:roles;type:UUID;" json:"roles,omitempty"`
	//部門
	Department string `gorm:"column:department;type:TEXT;not null;" json:"department,omitempty"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at,omitempty"`
}

// Base struct is corresponding to table structure file

type Base struct {
	//員工ID
	EmployeesID string `json:"employees_id,omitempty"`
	//員工名稱
	EmployeesName string `json:"employees_name,omitempty"`
	//角色ID
	Roles string `json:"roles,omitempty"`
	//部門
	Department string `json:"department,omitempty"`
	// 創建者
	CreatedBy string ` json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Single return structure file

type Single struct {

	//員工ID
	EmployeesID string `json:"employees_id,omitempty"`
	//員工名稱
	EmployeesName string `json:"employees_name,omitempty"`
	//角色ID
	Roles string `json:"roles,omitempty"`
	//部門
	Department string `json:"department,omitempty"`
	// 創建者
	CreatedBy string ` json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Created struct is used to create

type Created struct {

	//員工名稱
	EmployeesName string `json:"employees_name,omitempty" binding:"required" validate:"required"`
	//角色ID
	Roles string `json:"roles" binding:"required,uuid4" validate:"required"`
	//部門
	Department string `json:"department,omitempty" binding:"required" validate:"required"`
	//創建者
	CreatedBy string `json:"created_by" swaggerignore:"true"`
}

// Field is structure file for search

type Field struct {
	//員工ID
	EmployeesID string `json:"employees_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//員工名稱
	EmployeesName *string `json:"employees_name,omitempty" form:"employees_name"`
	//角色ID
	Roles *string `json:"roles,omitempty" form:"roles" binding:"omitempty,uuid4"`
	//部門
	Department *string `json:"department,omitempty" form:"department"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files

type List struct {
	Employees []*struct {

		//員工ID
		EmployeesID string `json:"employees_id,omitempty"`

		//員工名稱
		EmployeesName *string `json:"employees_name,omitempty" `
		//角色ID
		Roles *string `json:"roles,omitempty" `
		//部門
		Department *string `json:"department,omitempty" `
	} `json:"requisitions"`
	model.OutPage
}

// Updated struct is used to update

type Updated struct {

	//員工ID
	EmployeesID string `json:"employees_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//員工名稱
	EmployeesName *string `json:"employees_name,omitempty" `
	//角色ID
	Roles *string `json:"roles,omitempty" binding:"omitempty,uuid4"`
	//部門
	Department *string `json:"department,omitempty" `
}

// TableName sets the insert table name for this struct type

func (a *Table) TableName() string {
	return "employees"
}
