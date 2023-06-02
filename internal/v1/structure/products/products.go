package products

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {
	//產品ID
	ProductsID string `gorm:"primaryKey;column:products_id;uuid_generate_v4()type:UUID;" json:"products_id,omitempty"`
	//品名
	Product string `gorm:"column:product;type:TEXT;not null;" json:"product,omitempty"`
	//用途
	ProductUse string `gorm:"column:product_use;type:TEXT;not null;" json:"product_use,omitempty"`
	//單位
	Unit string `gorm:"column:unit;type:TEXT;not null;" json:"unit,omitempty"`
	//價格
	Price int `gorm:"column:price;type:INTEGER;not null;" json:"price,omitempty"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at,omitempty"`
}

// Base struct is corresponding to table structure file

type Base struct {
	//產品ID
	ProductsID string `json:"products_id,omitempty"`
	//品名
	Product string ` json:"product,omitempty"`
	//用途
	ProductUse string `json:"product_use,omitempty"`
	//單位
	Unit string `json:"unit,omitempty"`
	//價格
	Price int `json:"price,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Single return structure file

type Single struct {

	//產品ID
	ProductsID string `json:"products_id,omitempty"`
	//品名
	Product string ` json:"product,omitempty"`
	//用途
	ProductUse string `json:"product_use,omitempty"`
	//單位
	Unit string `json:"unit,omitempty"`
	//價格
	Price int `json:"price,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Created struct is used to create

type Created struct {
	//品名
	Product string `json:"product,omitempty" binding:"required" validate:"required"`
	//用途
	ProductUse string `json:"product_use,omitempty" binding:"required" validate:"required"`
	//單位
	Unit string `json:"unit,omitempty" binding:"required" validate:"required"`
	//價格
	Price int `json:"price,omitempty" binding:"required" validate:"required"`
	//創建者
	CreatedBy string `json:"created_by" swaggerignore:"true"`
}

// Field is structure file for search

type Field struct {
	//產品ID
	ProductsID string `json:"products_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//品名
	Product *string `json:"product,omitempty" form:"product"`
	//用途
	ProductUse *string `json:"product_use,omitempty" form:"product_use"`
	//單位
	Unit *string `json:"unit,omitempty" form:"unit"`
	//價格
	Price *int `json:"price,omitempty" form:"price"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files

type List struct {
	Products []*struct {
		//產品ID
		ProductsID string `json:"products_id,omitempty"`
		//品名
		Product *string `json:"product,omitempty" `
		//用途
		ProductUse *string `json:"product_use,omitempty" `
		//單位
		Unit *string `json:"unit,omitempty" `
		//價格
		Price *int `json:"price,omitempty" `
	} `json:"requisitions"`
	model.OutPage
}

// Updated struct is used to update

type Updated struct {

	//產品ID
	ProductsID string `json:"products_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//品名
	Product *string `json:"product,omitempty" `
	//用途
	ProductUse *string `json:"product_use,omitempty" `
	//單位
	Unit *string `json:"unit,omitempty" `
	//價格
	Price *int `json:"price,omitempty" `
}

// TableName sets the insert table name for this struct type

func (a *Table) TableName() string {
	return "products"
}
