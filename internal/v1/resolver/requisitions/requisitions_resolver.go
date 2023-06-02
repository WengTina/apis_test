package requisitions

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	requisitionsModel "eirc.app/internal/v1/structure/requisitions"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *requisitionsModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	requisitions, err := r.RequisitionsService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, requisitions.RequisitionsID)
}

func (r *resolver) List(input *requisitionsModel.Fields) interface{} {

	output := &requisitionsModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, requisitions, err := r.RequisitionsService.List(input)
	requisitionsByte, err := json.Marshal(requisitions)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(requisitionsByte, &output.Requisitions)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *requisitionsModel.Field) interface{} {

	base, err := r.RequisitionsService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontRequisitions := &requisitionsModel.Single{}
	requisitionsByte, _ := json.Marshal(base)
	err = json.Unmarshal(requisitionsByte, &frontRequisitions)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontRequisitions)
}

func (r *resolver) Deleted(input *requisitionsModel.Updated) interface{} {
	_, err := r.RequisitionsService.GetByID(&requisitionsModel.Field{RequisitionsID: input.RequisitionsID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RequisitionsService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *requisitionsModel.Updated) interface{} {
	requisitions, err := r.RequisitionsService.GetByID(&requisitionsModel.Field{RequisitionsID: input.RequisitionsID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RequisitionsService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, requisitions.RequisitionsID)
}
