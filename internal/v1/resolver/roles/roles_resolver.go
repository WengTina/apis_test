package roles

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	rolesModel "eirc.app/internal/v1/structure/roles"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *rolesModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	roles, err := r.RolesService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, roles.RolesID)
}

func (r *resolver) List(input *rolesModel.Fields) interface{} {

	output := &rolesModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, roles, err := r.RolesService.List(input)
	rolesByte, err := json.Marshal(roles)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(rolesByte, &output.Roles)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *rolesModel.Field) interface{} {

	base, err := r.RolesService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontRoles := &rolesModel.Single{}
	rolesByte, _ := json.Marshal(base)
	err = json.Unmarshal(rolesByte, &frontRoles)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontRoles)
}

func (r *resolver) Deleted(input *rolesModel.Updated) interface{} {
	_, err := r.RolesService.GetByID(&rolesModel.Field{RolesID: input.RolesID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RolesService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *rolesModel.Updated) interface{} {
	roles, err := r.RolesService.GetByID(&rolesModel.Field{RolesID: input.RolesID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RolesService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, roles.RolesID)
}
