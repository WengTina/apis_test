package purchases

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	purchasesModel "eirc.app/internal/v1/structure/purchases"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *purchasesModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	purchases, err := r.PurchasesService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, purchases.PurchasesID)
}

func (r *resolver) List(input *purchasesModel.Fields) interface{} {

	output := &purchasesModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, purchases, err := r.PurchasesService.List(input)
	purchasesByte, err := json.Marshal(purchases)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(purchasesByte, &output.Purchases)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *purchasesModel.Field) interface{} {

	base, err := r.PurchasesService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontPurchases := &purchasesModel.Single{}
	purchasesByte, _ := json.Marshal(base)
	err = json.Unmarshal(purchasesByte, &frontPurchases)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontPurchases)
}

func (r *resolver) Deleted(input *purchasesModel.Updated) interface{} {
	_, err := r.PurchasesService.GetByID(&purchasesModel.Field{PurchasesID: input.PurchasesID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.PurchasesService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *purchasesModel.Updated) interface{} {
	purchases, err := r.PurchasesService.GetByID(&purchasesModel.Field{PurchasesID: input.PurchasesID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.PurchasesService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, purchases.PurchasesID)
}
