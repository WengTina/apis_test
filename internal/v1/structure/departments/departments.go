package departments

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {
	//部門ID
	DepartmentsID string `gorm:"primaryKey;column:departments_id;uuid_generate_v4()type:UUID;" json:"departments_id,omitempty"`
	//部門代號
	DepartmentsCode int `gorm:"->;column:departments_code;type:TEXT;not null;" json:"departments_code,omitempty"`
	//部門名稱
	DepartmentName string `gorm:"column:department_name;type:TEXT;not null;" json:"department_name,omitempty"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at,omitempty"`
}

// Base struct is corresponding to table structure file

type Base struct {

	//部門ID
	DepartmentsID string `json:"departments_id,omitempty"`
	//部門代號
	DepartmentsCode int `json:"departments_code,omitempty"`
	//部門名稱
	DepartmentName string `json:"department_name,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Single return structure file

type Single struct {

	//部門ID
	DepartmentsID string `json:"departments_id,omitempty"`
	//部門代號
	DepartmentsCode int `json:"departments_code,omitempty"`
	//部門名稱
	DepartmentName string `json:"department_name,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Created struct is used to create

type Created struct {
	//部門名稱
	DepartmentName string `json:"department_name,omitempty" binding:"required" validate:"required"`
	//創建者
	CreatedBy string `json:"created_by" swaggerignore:"true"`
}

// Field is structure file for search

type Field struct {
	//部門ID
	DepartmentsID string `json:"departments_id,omitempty" binding:"omitempty,uuid4"`
	//部門代號
	DepartmentsCode string `json:"departments_code,omitempty" form:"departments_code"`
	//部門名稱
	DepartmentName *string `json:"department_name,omitempty" form:"department_name"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files

type List struct {
	Departments []*struct {

		//部門ID
		DepartmentsID string `json:"departments_id,omitempty"`
		//部門代號
		DepartmentsCode string `json:"departments_code,omitempty"`
		//部門名稱
		DepartmentName *string `json:"department_name,omitempty" `
	} `json:"requisitions"`
	model.OutPage
}

// Updated struct is used to update

type Updated struct {

	//部門ID
	DepartmentsID string `json:"departments_id,omitempty" binding:"omitempty,uuid4"`
	//部門代號
	DepartmentsCode string `json:"departments_code,omitempty" `
	//部門名稱
	DepartmentName *string `json:"department_name,omitempty"`
}

// TableName sets the insert table name for this struct type

func (a *Table) TableName() string {
	return "departments"
}
