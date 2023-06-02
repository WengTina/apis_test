package requisitions

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {

	// 請購單ID
	RequisitionsID string `gorm:"primaryKey;column:requisitions_id;uuid_generate_v4()type:UUID;" json:"requisitions_id,omitempty"`
	//請購碼
	RequisitionsCode int `gorm:"->;column:requisitions_code;type:TEXT;not null;" json:"requisitions_code,omitempty"`
	// 申請者名稱
	ApplicantName string `gorm:"column:applicantname;type:VARCHAR;not null;" json:"applicantname,omitempty"`
	// 公司
	Company string `gorm:"column:company;type:VARCHAR;not null;" json:"company,omitempty"`
	// 部門
	Department string `gorm:"column:department;type:VARCHAR;not null;" json:"department,omitempty"`
	//產品
	Product string `gorm:"column:product;type:VARCHAR;not null;" json:"product,omitempty"`
	//數量
	Quantity int `gorm:"column:quantity;type:INTEGER;not null;" json:"quantity,omitempty"`
	//價格
	Price int `gorm:"column:price;type:INTEGER;not null;" json:"price,omitempty"`
	//創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at,omitempty"`
}

// Base struct is corresponding to table structure file

type Base struct {

	// 請購單ID
	RequisitionsID string ` json:"requisitions_id,omitempty"`
	// 請購碼
	RequisitionsCode string ` json:"requisitions_code,omitempty"`
	// 申請者名稱
	ApplicantName string `json:"applicantname,omitempty"`
	// 公司
	Company string `json:"company,omitempty"`
	// 部門
	Department string `json:"department,omitempty"`
	//產品
	Product string `json:"product,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//價格
	Price int `json:"price,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Single return structure file

type Single struct {

	// 請購單ID
	RequisitionsID string ` json:"requisitions_id,omitempty"`
	// 請購碼
	RequisitionsCode string ` json:"requisitions_code,omitempty"`
	// 申請者名稱
	ApplicantName string `json:"applicantname,omitempty"`
	// 公司
	Company string `json:"company,omitempty"`
	// 部門
	Department string `json:"department,omitempty"`
	//產品
	Product string `json:"product,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//價格
	Price int `json:"price,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Created struct is used to create

type Created struct {
	//申請者名稱
	ApplicantName string `json:"applicantname,omitempty" binding:"required" validate:"required"`
	//公司
	Company string `json:"company,omitempty" binding:"required" validate:"required"`
	//部門
	Department string `json:"department,omitempty" binding:"required" validate:"required"`
	//產品
	Product string `json:"product,omitempty" binding:"required" validate:"required"`
	//數量
	Quantity int `json:"quantity,omitempty" binding:"required" validate:"required"`
	//價錢
	Price int `json:"price,omitempty" binding:"required" validate:"required"`
}

// Field is structure file for search

type Field struct {

	// 請購單ID
	RequisitionsID string `json:"requisitions_id,omitempty" binding:"omitempty,uuid4"`
	// 請購碼
	RequisitionsCode string `json:"requisitions_code,omitempty" form:"requisitions_code"`
	// 申請者名稱
	ApplicantName *string `json:"applicantname,omitempty" form:"applicantname"`
	// 公司
	Company *string `json:"company,omitempty" form:"company"`
	//部門
	Department *string `json:"department,omitempty" form:"department"`
	//產品
	Product *string `json:"product,omitempty" form:"product"`
	//數量
	Quantity *int `json:"quantity,omitempty" form:"quantity"`
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
	Requisitions []*struct {

		// 請購單ID
		RequisitionsID string `json:"requisitions_id,omitempty"`
		// 請購碼
		RequisitionsCode string `json:"requisitions_code,omitempty"`
		// 申請者名稱
		ApplicantName *string `json:"applicantname,omitempty" `
		// 公司
		Company *string `json:"company,omitempty" `
		//部門
		Department *string `json:"department,omitempty" `
		//產品
		Product *string `json:"product,omitempty" `
		//數量
		Quantity *int `json:"quantity,omitempty"`
		//價格
		Price *int `json:"price,omitempty"`
	} `json:"requisitions"`
	model.OutPage
}

// Updated struct is used to update

type Updated struct {

	// 請購單ID
	RequisitionsID string `json:"requisitions_id,omitempty" binding:"omitempty,uuid4"`
	// 請購碼
	RequisitionsCode string `json:"requisitions_code,omitempty" `
	// 申請者名稱
	ApplicantName *string `json:"applicantname,omitempty"`
	// 公司
	Company *string `json:"company,omitempty" `
	//部門
	Department *string `json:"department,omitempty" `
	//產品
	Product *string `json:"product,omitempty"`
	//數量
	Quantity *int `json:"quantity,omitempty" `
	//價格
	Price *int `json:"price,omitempty" `
}

// TableName sets the insert table name for this struct type

func (a *Table) TableName() string {
	return "requisitions"
}
