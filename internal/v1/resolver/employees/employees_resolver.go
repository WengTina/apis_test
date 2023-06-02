package employees

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	employeesModel "eirc.app/internal/v1/structure/employees"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *employeesModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	employees, err := r.EmployeesService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, employees.EmployeesID)
}

func (r *resolver) List(input *employeesModel.Fields) interface{} {

	output := &employeesModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, employees, err := r.EmployeesService.List(input)
	employeesByte, err := json.Marshal(employees)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(employeesByte, &output.Employees)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *employeesModel.Field) interface{} {

	base, err := r.EmployeesService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontEmployees := &employeesModel.Single{}
	employeesByte, _ := json.Marshal(base)
	err = json.Unmarshal(employeesByte, &frontEmployees)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontEmployees)
}

func (r *resolver) Deleted(input *employeesModel.Updated) interface{} {
	_, err := r.EmployeesService.GetByID(&employeesModel.Field{EmployeesID: input.EmployeesID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.EmployeesService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *employeesModel.Updated) interface{} {
	employees, err := r.EmployeesService.GetByID(&employeesModel.Field{EmployeesID: input.EmployeesID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.EmployeesService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, employees.EmployeesID)
}
