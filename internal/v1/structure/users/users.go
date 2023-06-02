package users

import (
	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {

	// 使用者ID
	UserID int `gorm:"column:user_id;type:serial;not null;" json:"user_id,omitempty"`
	// 中文名稱
	UserName string `gorm:"column:username;type:VARCHAR;not null;" json:"username,omitempty"`
	// 密碼
	Password string `gorm:"column:pwd;type:VARCHAR;not null;" json:"password,omitempty"`
	// 角色編號
	Email string `gorm:"column:email;type:VARCHAR;not null;" json:"email,omitempty"`
}

// Base struct is corresponding to table structure file

type Base struct {

	// 使用者ID
	UserID int `json:"user_id,omitempty"`

	// 中文名稱
	UserName string `json:"username,omitempty"`
	// 密碼
	Password string `json:"password,omitempty"`
	// email
	Email string `json:"email,omitempty"`
}

// Single return structure file

type Single struct {

	// 使用者ID
	UserID string `json:"user_id,omitempty"`
	// 名稱
	UserName string `json:"username,omitempty"`
	// 密碼
	Password string `json:"password,omitempty"`
	// email
	Email string `json:"email,omitempty"`
}

// Created struct is used to create

type Created struct {
	// 使用者名稱
	UserName string `json:"userame,omitempty" binding:"required" validate:"required"`
	// 密碼
	Password string `json:"password,omitempty" binding:"required" validate:"required"`
	// 角色編號
	Email string `json:"email,omitempty" binding:"required" validate:"required"`
}

// Field is structure file for search

type Field struct {

	// 使用者ID
	UserID int `json:"user_id,omitempty" form:"user_id" binding:"omitempty,serial"`
	// 中文名稱
	UserName *string `json:"username,omitempty" form:"username"`
	// mail
	Email *string `json:"email,omitempty" form:"email"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files

type List struct {
	Users []*struct {

		// 使用者ID
		UserID int `json:"user_id,omitempty"`
		// 中文名稱
		UserName string `json:"username,omitempty"`
		// mail
		Email string `json:"email,omitempty"`
	} `json:"users"`
	model.OutPage
}

// Updated struct is used to update

type Updated struct {

	// 使用者ID
	UserID int `json:"user_id,omitempty" binding:"omitempty,serial"`
	// 中文名稱
	UserName *string `json:"username,omitempty"`
	// 密碼
	Password *string `json:"password,omitempty"`
	// mail
	Email *string `json:"email,omitempty"`
}

// TableName sets the insert table name for this struct type

func (a *Table) TableName() string {
	return "users"
}
