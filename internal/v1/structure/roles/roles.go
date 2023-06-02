package roles

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {
	//角色ID
	RolesID string `gorm:"primaryKey;column:roles_id;uuid_generate_v4()type:UUID;" json:"roles_id,omitempty"`
	//角色名稱
	RoleName string `gorm:"column:role_name;type:TEXT;not null;" json:"role_name,omitempty"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at,omitempty"`
}

// Base struct is corresponding to table structure file

type Base struct {
	//角色ID
	RolesID string `json:"roles_id,omitempty"`
	//角色名稱
	RoleName string `json:"role_name,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Single return structure file

type Single struct {

	//角色ID
	RolesID string `json:"roles_id,omitempty"`
	//角色名稱
	RoleName string `json:"role_name,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Created struct is used to create

type Created struct {
	//角色名稱
	RoleName string `json:"role_name,omitempty" binding:"required" validate:"required"`
	//創建者
	CreatedBy string `json:"created_by" swaggerignore:"true"`
}

// Field is structure file for search

type Field struct {
	//角色ID
	RolesID string `json:"roles_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//角色名稱
	RoleName *string `json:"role_name,omitempty" form:"role_name"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files

type List struct {
	Roles []*struct {

		//角色ID
		RolesID string `json:"roles_id,omitempty"`
		//角色名稱
		RoleName *string `json:"role_name,omitempty"`
	} `json:"requisitions"`
	model.OutPage
}

// Updated struct is used to update

type Updated struct {

	//角色ID
	RolesID string `json:"roles_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//角色名稱
	RoleName *string `json:"role_name,omitempty"`
}

// TableName sets the insert table name for this struct type

func (a *Table) TableName() string {
	return "roles"
}
