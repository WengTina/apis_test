package accounts

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// 透過gorm聲明模型定義資料表結構，需與資料庫中的資料表欄位一致，資料表關聯的飲用也需宣告在此
// json tag 用在序列化及反序列化的過程
// Table struct is database table struct
type Table struct {
	// 編號
	AccountID string `gorm:"primaryKey;column:account_id;uuid_generate_v4()type:UUID;" json:"account_id,omitempty"` //omitempty 序列化時，欄位為空值則忽略
	// 公司ID
	CompanyID string `gorm:"column:company_id;type:UUID;" json:"company_id,omitempty"`
	// 帳號
	Account string `gorm:"column:account;type:VARCHAR;" json:"account,omitempty"`
	// 中文名稱
	Name string `gorm:"column:name;type:VARCHAR;" json:"name,omitempty"`
	// 密碼
	Password string `gorm:"column:pwd;type:VARCHAR;" json:"password,omitempty"`
	// 角色編號
	RoleID string `gorm:"column:role_id;type:VARCHAR;" json:"role_id,omitempty"`
	// 是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:bool;false" json:"is_deleted,omitempty"`
	// 創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	// 更新時間
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string `gorm:"column:updated_by;type:UUID;" json:"updated_by,omitempty"`
}

// Base struct is corresponding to table structure file
// 在service層用於存放資料轉換後的資料，對應table結構
type Base struct {
	// 編號
	AccountID string `json:"account_id,omitempty"`
	// 公司ID
	CompanyID string `json:"company_id,omitempty"`
	// 帳號
	Account string `json:"account,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 密碼
	Password string `json:"password,omitempty"`
	// 角色編號
	RoleID string `json:"role_id,omitempty"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// 創建時間
	CreatedAt time.Time `json:"created_at"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新時間
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// Single return structure file
// 回傳單比資料用的結構
type Single struct {
	// 編號
	AccountID string `json:"account_id,omitempty"`
	// 公司ID
	CompanyID string `json:"company_id,omitempty"`
	// 帳號
	Account string `json:"account,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 密碼
	Password string `json:"password,omitempty"`
	// 角色編號
	RoleID string `json:"role_id,omitempty"`
	// 創建時間
	CreatedAt time.Time `json:"created_at"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新時間
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// Created struct is used to create
// 創建資料用的結構，存放creat時須輸入的欄位
type Created struct {
	// 公司ID
	CompanyID string `json:"company_id" binding:"required,uuid4" validate:"required"` //validate tag 進行資料庫寫入前對資料進一步校驗
	// 帳號
	Account string `json:"account,omitempty" binding:"required" validate:"required"`
	// 中文名稱
	Name string `json:"name" binding:"required" validate:"required"`
	// 密碼
	Password string `json:"password" binding:"required" validate:"required"`
	// 角色編號
	RoleID string `json:"role_id" binding:"required" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by" swaggerignore:"true"`
}

// Field is structure file for search
// 作為查詢用結構，存放需要查詢的欄位，id為必要
// 包含Field 和分頁結構
// binding tag用於指定驗證規則，可在解析http請求參數時，對該欄位進行自動校驗，確保解析出的資瞭符合預期的格式和類型
// omitempty 表示可為空值，Required 不能為空值
type Field struct {
	// 編號
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"` //swaggerignore:"true"表示該欄位不會自動新增至swagger API
	// 公司ID
	//form tag 標示結構字段在表單中對應的名稱
	//在解析http請求參數時，會從http request中對應欄位獲取並附值到該欄位
	CompanyID *string `json:"company_id,omitempty" form:"company_id" binding:"omitempty,uuid4"`
	// 帳號
	Account *string `json:"account,omitempty" form:"account"`
	// 中文名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 角色編號
	RoleID *string `json:"role_id,omitempty" form:"role_id"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
// 回傳多筆資料用的結構，包含輸出結構及分頁輸出結構
type List struct {
	Accounts []*struct {
		// 編號
		AccountID string `json:"account_id,omitempty"`
		// 公司ID
		CompanyID string `json:"company_id,omitempty"`
		// 帳號
		Account string `json:"account,omitempty"`
		// 中文名稱
		Name string `json:"name,omitempty"`
		// 角色編號
		RoleID string `json:"role_id,omitempty"`
	} `json:"accounts"`
	model.OutPage
}

// Updated struct is used to update
// 更新資料用的結構，存放id和可變動的欄位
type Updated struct {
	// 編號
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 組織ID
	CompanyID *string `json:"company_id,omitempty" binding:"omitempty,uuid4"`
	// 中文名稱
	Name *string `json:"name,omitempty"`
	// 密碼
	Password *string `json:"password,omitempty"`
	// 角色編號
	RoleID *string `json:"role_id,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// TableName sets the insert table name for this struct type
// TableName 要寫入的資料表名稱，需與實際資料表名稱一致
// table 全以小寫命名，以複數's','es'結尾
// 兩字以上以下劃線隔開
func (a *Table) TableName() string {
	return "accounts"
}
