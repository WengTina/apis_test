package products

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	productsModel "eirc.app/internal/v1/structure/products"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *productsModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	products, err := r.ProductsService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, products.ProductsID)
}

func (r *resolver) List(input *productsModel.Fields) interface{} {

	output := &productsModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, products, err := r.ProductsService.List(input)
	productsByte, err := json.Marshal(products)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(productsByte, &output.Products)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *productsModel.Field) interface{} {

	base, err := r.ProductsService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontProducts := &productsModel.Single{}
	productsByte, _ := json.Marshal(base)
	err = json.Unmarshal(productsByte, &frontProducts)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontProducts)
}

func (r *resolver) Deleted(input *productsModel.Updated) interface{} {
	_, err := r.ProductsService.GetByID(&productsModel.Field{ProductsID: input.ProductsID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ProductsService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *productsModel.Updated) interface{} {
	products, err := r.ProductsService.GetByID(&productsModel.Field{ProductsID: input.ProductsID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ProductsService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, products.ProductsID)
}
