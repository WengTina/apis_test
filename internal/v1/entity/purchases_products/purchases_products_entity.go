package purchases_products

import (
	model "eirc.app/internal/v1/structure/purchases_products"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})
	if input.PurchasesID != nil {
		db.Where("purchases_id = ?", input.PurchasesID)
	}

	if input.Product != nil {
		db.Where("product = ?", input.Product)
	}

	if input.Subtotal != nil {
		db.Where("subtotal = ?", input.Subtotal)
	}

	if input.Total != nil {
		db.Where("total = ?", input.Total)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_at desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("purchases_products_id = ?", input.PurchasesProductsID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Field) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Save(&input).Error

	return err
}
