package purchases_products

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {
	//請購單產品ID
	PurchasesProductsID string `gorm:"primaryKey;column:purchases_products_id;uuid_generate_v4()type:UUID;" json:"purchases_products_id,omitempty"`
	//請購單ID
	PurchasesID string `gorm:"column:purchases_id;type:UUID;" json:"purchases_id,omitempty"`
	//品名
	Product string `gorm:"column:product;type:TEXT;not null;" json:"product,omitempty"`
	//小計
	Subtotal int `gorm:"column:subtotal;type:INTEGER;not null;" json:"subtotal,omitempty"`
	//合計
	Total int `gorm:"column:total;type:INTEGER;not null;" json:"total,omitempty"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at,omitempty"`
}

// Base struct is corresponding to table structure file

type Base struct {
	//請購單產品ID
	PurchasesProductsID string `json:"purchases_products_id,omitempty"`
	//請購單ID
	PurchasesID string `json:"purchases_id,omitempty"`
	//品名
	Product string ` json:"product,omitempty"`
	//小計
	Subtotal int `json:"subtotal,omitempty"`
	//合計
	Total int `json:"total,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Single return structure file

type Single struct {

	//請購單產品ID
	PurchasesProductsID string `json:"purchases_products_id,omitempty"`
	//請購單ID
	PurchasesID string `json:"purchases_id,omitempty"`
	//品名
	Product string ` json:"product,omitempty"`
	//小計
	Subtotal int `json:"subtotal,omitempty"`
	//合計
	Total int `json:"total,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Created struct is used to create

type Created struct {
	//請購單ID
	PurchasesID string `json:"purchases_id" binding:"required,uuid4" validate:"required"`
	//品名
	Product string `json:"product,omitempty" binding:"required" validate:"required"`
	//小計
	Subtotal int `json:"subtotal,omitempty" binding:"required" validate:"required"`
	//合計
	Total int `json:"total,omitempty" binding:"required" validate:"required"`
	//創建者
	CreatedBy string `json:"created_by" swaggerignore:"true"`
}

// Field is structure file for search

type Field struct {
	//請購單產品ID
	PurchasesProductsID string `json:"purchases_products_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//請購單ID
	PurchasesID *string `json:"purchases_id,omitempty" form:"purchases_id" binding:"omitempty,uuid4"`
	//品名
	Product *string `json:"product,omitempty" form:"product"`
	//小計
	Subtotal *int `json:"subtotal,omitempty" form:"subtotal"`
	//合計
	Total *int `json:"total,omitempty" form:"total"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files

type List struct {
	PurchasesProducts []*struct {

		//請購單產品ID
		PurchasesProductsID string `json:"purchases_products_id,omitempty"`
		//請購單ID
		PurchasesID *string `json:"purchases_id,omitempty" `
		//品名
		Product *string `json:"product,omitempty" `
		//小計
		Subtotal *int `json:"subtotal,omitempty" `
		//合計
		Total *int `json:"total,omitempty" `
	} `json:"requisitions"`
	model.OutPage
}

// Updated struct is used to update

type Updated struct {

	//請購單ID
	PurchasesProductsID string `json:"purchases_products_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//請購單ID
	PurchasesID *string `json:"purchases_id,omitempty" binding:"omitempty,uuid4"`
	//品名
	Product *string `json:"product,omitempty" `
	//小計
	Subtotal *int `json:"subtotal,omitempty" `
	//合計
	Total *int `json:"total,omitempty" `
}

// TableName sets the insert table name for this struct type

func (a *Table) TableName() string {
	return "purchases_products"
}
