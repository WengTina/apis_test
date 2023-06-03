package employees

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/employees"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增使用者
// @description 新增使用者
// @Tags Employees
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body employees.Created true "新增使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/employees [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	createBy := util.GenerateUUID()
	input := &employees.Created{}
	input.CreatedBy = createBy
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.EmployeesResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 條件搜尋使用者
// @description 條件使用者
// @Tags Employees
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param organizationID query string false "組織ID"
// @param employees query string false "帳號"
// @param chineseName query string false "中文名稱"
// @param roleName query string false "角色名稱"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=employees.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/employees [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &employees.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.EmployeesResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一使用者
// @description 取得單一使用者
// @Tags Employees
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param employeesID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=employees.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/employees/{employeesID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	employeesID := ctx.Param("employeesID")
	input := &employees.Field{}
	input.EmployeesID = employeesID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.EmployeesResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一使用者
// @description 刪除單一使用者
// @Tags Employees
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param employeesID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/employees/{employeesID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者

	employeesID := ctx.Param("employeesID")
	input := &employees.Updated{}
	input.EmployeesID = employeesID

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.EmployeesResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一使用者
// @description 更新單一使用者
// @Tags Employees
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param employeesID path string true "使用者ID"
// @param * body employees.Updated true "更新使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/employees/{employeesID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者

	employeesID := ctx.Param("employeesID")
	input := &employees.Updated{}
	input.EmployeesID = employeesID

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.EmployeesResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
