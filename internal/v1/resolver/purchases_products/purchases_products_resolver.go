package purchases_products

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	purchases_productsModel "eirc.app/internal/v1/structure/purchases_products"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *purchases_productsModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	purchases_products, err := r.PurchasesProductsService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, purchases_products.PurchasesProductsID)
}

func (r *resolver) List(input *purchases_productsModel.Fields) interface{} {

	output := &purchases_productsModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, purchases_products, err := r.PurchasesProductsService.List(input)
	purchases_productsByte, err := json.Marshal(purchases_products)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(purchases_productsByte, &output.PurchasesProducts)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *purchases_productsModel.Field) interface{} {

	base, err := r.PurchasesProductsService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontPurchasesProducts := &purchases_productsModel.Single{}
	purchases_productsByte, _ := json.Marshal(base)
	err = json.Unmarshal(purchases_productsByte, &frontPurchasesProducts)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontPurchasesProducts)
}

func (r *resolver) Deleted(input *purchases_productsModel.Updated) interface{} {
	_, err := r.PurchasesProductsService.GetByID(&purchases_productsModel.Field{PurchasesProductsID: input.PurchasesProductsID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.PurchasesProductsService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *purchases_productsModel.Updated) interface{} {
	purchases_products, err := r.PurchasesProductsService.GetByID(&purchases_productsModel.Field{PurchasesProductsID: input.PurchasesProductsID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.PurchasesProductsService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, purchases_products.PurchasesProductsID)
}
