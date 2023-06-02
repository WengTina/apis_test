package requisitions

import (
	model "eirc.app/internal/v1/structure/requisitions"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	if input.Company != nil {
		db.Where("company = ?", input.Company)
	}

	if input.ApplicantName != nil {
		db.Where("applicantname like %?%", *input.ApplicantName)
	}

	if input.Department != nil {
		db.Where("department = ?", input.Department)
	}

	if input.Quantity != nil {
		db.Where("quantity = ?", input.Quantity)
	}

	if input.Product != nil {
		db.Where("product = ?", input.Product)
	}

	if input.Price != nil {
		db.Where("price = ?", input.Price)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_at desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("requisitions_id = ?", input.RequisitionsID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("requisitions_id = ?", input.RequisitionsID).Save(&input).Error

	return err
}
