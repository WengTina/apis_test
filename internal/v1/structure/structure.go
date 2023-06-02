package structure

//存放共用的結構，需要時再進行調用

//查詢多筆資料時需要的「輸入」分頁結構
type InPage struct {
	//頁數(請從1開始帶入)
	Page int64 `json:"page" binding:"required,gt=0" validate:"required" form:"page"`
	//筆數(請從1開始帶入,最高上限20)
	Limit int64 `json:"limit" binding:"required,gt=0" validate:"required" form:"limit"`
}

//查詢多筆資料需要的「輸出」分頁結構
type OutPage struct {
	InPage
	//總頁數
	Pages int64 `json:"pages"`
	Total int64 `json:"total"`
}
