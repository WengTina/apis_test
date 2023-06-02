package departments

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	departmentsModel "eirc.app/internal/v1/structure/departments"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *departmentsModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	departments, err := r.DepartmentsService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, departments.DepartmentsID)
}

func (r *resolver) List(input *departmentsModel.Fields) interface{} {

	output := &departmentsModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, departments, err := r.DepartmentsService.List(input)
	departmentsByte, err := json.Marshal(departments)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(departmentsByte, &output.Departments)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *departmentsModel.Field) interface{} {

	base, err := r.DepartmentsService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontDepartments := &departmentsModel.Single{}
	departmentsByte, _ := json.Marshal(base)
	err = json.Unmarshal(departmentsByte, &frontDepartments)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontDepartments)
}

func (r *resolver) Deleted(input *departmentsModel.Updated) interface{} {
	_, err := r.DepartmentsService.GetByID(&departmentsModel.Field{DepartmentsID: input.DepartmentsID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.DepartmentsService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *departmentsModel.Updated) interface{} {
	departments, err := r.DepartmentsService.GetByID(&departmentsModel.Field{DepartmentsID: input.DepartmentsID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.DepartmentsService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, departments.DepartmentsID)
}
