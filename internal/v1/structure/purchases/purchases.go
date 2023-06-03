package purchases

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {

	// 請購單ID
	PurchasesID string `gorm:"primaryKey;column:purchases_id;uuid_generate_v4()type:UUID;" json:"purchases_id,omitempty"`
	//申請人
	ApplicantName string `gorm:"column:applicant_name;type:TEXT;not null;" json:"applicant_name,omitempty"`
	//產品
	Product string `gorm:"column:product;type:TEXT;not null;" json:"product,omitempty"`
	//數量
	Quantity int `gorm:"column:quantity;type:INTEGER;not null;" json:"quantity,omitempty"`
	//請購事由
	PurchasesReason string `gorm:"column:purchases_reason;type:TEXT;not null;" json:"purchases_reason,omitempty"`
	//請購日期
	PurchasesDate string `gorm:"column:purchases_date;type:DATE;not null;" json:"purchases_date,omitempty"`
	//需求日期
	DemandDate string `gorm:"column:demand_date;type:DATE;not null;" json:"demand_date,omitempty"`
	//備註
	Remark string `gorm:"column:remark;type:TEXT;not null;" json:"remark,omitempty"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at,omitempty"`
}

// Base struct is corresponding to table structure file

type Base struct {
	// 請購單ID
	PurchasesID string ` json:"purchases_id,omitempty"`
	//申請人
	ApplicantName string `json:"applicant_name,omitempty"`
	//產品
	Product string `json:"product,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//請購事由
	PurchasesReason string ` json:"purchases_reason,omitempty"`
	//請購日期
	PurchasesDate string ` json:"purchases_date,omitempty"`
	//需求日期
	DemandDate string ` json:"demand_date,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Single return structure file

type Single struct {

	// 請購單ID
	PurchasesID string ` json:"purchases_id,omitempty"`
	//申請人
	ApplicantName string `json:"applicant_name,omitempty"`
	//產品
	Product string `json:"product,omitempty"`
	//數量
	Quantity int `json:"quantity,omitempty"`
	//請購事由
	PurchasesReason string ` json:"purchases_reason,omitempty"`
	//請購日期
	PurchasesDate string ` json:"purchases_date,omitempty"`
	//需求日期
	DemandDate string ` json:"demand_date,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Created struct is used to create

type Created struct {
	//申請人
	ApplicantName string `json:"applicant_name,omitempty" binding:"required" validate:"required"`
	//產品
	Product string `json:"product,omitempty" binding:"required" validate:"required"`
	//數量
	Quantity int `json:"quantity,omitempty" binding:"required" validate:"required"`
	//請購事由
	PurchasesReason string `json:"purchases_reason,omitempty" binding:"required" validate:"required"`
	//請購日期
	PurchasesDate string `json:"purchases_date,omitempty" binding:"required" validate:"required"`
	//需求日期
	DemandDate string `json:"demand_date,omitempty" binding:"required" validate:"required"`
	//備註
	Remark string `json:"remark,omitempty" binding:"required" validate:"required"`
	//創建者
	CreatedBy string `json:"created_by" swaggerignore:"true"`
}

// Field is structure file for search

type Field struct {
	//請購單ID
	PurchasesID string `json:"purchases_id,omitempty" binding:"omitempty,uuid4" `
	//申請人
	ApplicantName *string `json:"applicant_name,omitempty" form:"applicant_name"`
	//產品
	Product *string `json:"product,omitempty" form:"product"`
	//數量
	Quantity *int `json:"quantity,omitempty" form:"quantity"`
	//請購事由
	PurchasesReason *string `json:"purchases_reason,omitempty" form:"purchases_reason"`
	//請購日期
	PurchasesDate *string `json:"purchases_date,omitempty" form:"purchases_date"`
	//需求日期
	DemandDate *string `json:"demand_date,omitempty" form:"demand_date"`
	//備註
	Remark *string `json:"remark,omitempty" form:"remark"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files

type List struct {
	Purchases []*struct {

		//請購單ID
		PurchasesID string `json:"purchases_id,omitempty"`
		//申請人
		ApplicantName *string `json:"applicant_name,omitempty" `
		//產品
		Product *string `json:"product,omitempty"`
		//數量
		Quantity *int `json:"quantity,omitempty" `
		//請購事由
		PurchasesReason *string `json:"purchases_reason,omitempty"`
		//請購日期
		PurchasesDate *string `json:"purchases_date,omitempty"`
		//需求日期
		DemandDate *string `json:"demand_date,omitempty"`
		//備註
		Remark *string `json:"remark,omitempty" `
	} `json:"purchases"`
	model.OutPage
}

// Updated struct is used to update

type Updated struct {

	//請購單ID
	PurchasesID string `json:"purchases_id,omitempty" binding:"omitempty,uuid4"`
	//申請人
	ApplicantName *string `json:"applicant_name,omitempty" `
	//產品
	Product *string `json:"product,omitempty"`
	//數量
	Quantity *int `json:"quantity,omitempty" `
	//請購事由
	PurchasesReason *string `json:"purchases_reason,omitempty"`
	//請購日期
	PurchasesDate *string `json:"purchases_date,omitempty"`
	//需求日期
	DemandDate *string `json:"demand_date,omitempty"`
	//備註
	Remark *string `json:"remark,omitempty" `
}

// TableName sets the insert table name for this struct type

func (a *Table) TableName() string {
	return "purchases"
}
