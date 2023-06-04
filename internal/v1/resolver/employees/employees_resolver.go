package employees

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	employeesModel "eirc.app/internal/v1/structure/employees"
	"github.com/docker/docker/libcontainerd/supervisor"
	"google.golang.org/genproto/googleapis/storage/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
// func (s *storage)GetBySingle(input *model.Base)(output *model.Table, err error){
// 	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
// 	if input.EmployeesID != nil{
// 		query.Where(query:"employees_id = ?",input.EmployeesID)
// 	}

// 	err =query.First(&output).Error
// 	if err != nil{
// 		log.Error(err)
// 		return output nil,err
// 	}
// 	return output,err:nil
// }

func (s *storage) GetBySingle(input *model.Base) (*model.Table, error) {
    query := s.db.Model(&model.Table{}).Preload(clause.Associations)
    if input.EmployeesID != nil {
        query = query.Where("employees_id = ?", *input.EmployeesID)
    }

    output := &model.Table{}
    err := query.First(output).Error
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return output, nil
}

func (m *manager) GetBySingle(input *employeesModel.Field)(int,interface{})  {
	employeesBase, err :=m.EmployeesService.GetBySingle(input)
	if err != nil{
		if errors.Is(err,ErrRecordNotFound) :code.DoesNotExist,code.GetCodeMessage(code.InternalServerError, err)

		log.Error(err)
		return code.InternalServerError,code.GetCodeMessage(code.InternalServerError, err.Error)
	}
	
	output := &employeesModel.Single{}
	employeesByte,_ :=json.Marshal(employeesByte)
	err = json.Unmarshal(employeesByte,&output)
	if err != nil{
		log.Error(err)
		return code.InternalServerError,code.GetCodeMessage(code.InternalServerError, err.Error)
	}

	output.createBy = *employeesByte.createByUsers.Name
	output.DepartmentsName = *employeesBase.Departments.Name
	
}