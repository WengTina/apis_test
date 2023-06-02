package purchases

import (
	model "eirc.app/internal/v1/structure/purchases"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})
	if input.Quantity != nil {
		db.Where("quantity = ?", input.Quantity)
	}

	if input.Product != nil {
		db.Where("product = ?", input.Product)
	}

	if input.ApplicantName != nil {
		db.Where("applicant_name like %?%", *input.ApplicantName)
	}

	if input.PurchasesReason != nil {
		db.Where("purchases_reason = ?", input.PurchasesReason)
	}

	if input.PurchasesDate != nil {
		db.Where("purchases_date = ?", input.PurchasesDate)
	}

	if input.DemandDate != nil {
		db.Where("demand_date = ?", input.DemandDate)
	}

	if input.Remark != nil {
		db.Where("remark = ?", input.Remark)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_at desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("purchases_id = ?", input.PurchasesID)

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
