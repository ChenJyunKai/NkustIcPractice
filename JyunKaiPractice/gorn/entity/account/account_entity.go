package accounts

import (
	"bitbucket.org/henry111666888/hs_gorm/pkg/dao"
	model "bitbucket.org/henry111666888/hs_gorm/structure/accounts"
)

func Created(input model.Table) (err error) {
	err = dao.DB.Create(&input).Error

	return err
}

func List(input model.Table) (err error) {
	err = dao.DB.Find(&input).Error

	return err
}

func GetByAccount(input model.Table, account_id string) (err error) {
	err = dao.DB.Find(&input, "account_id = ?", account_id).Error

	return err
}

func Updated(input model.Table, account_id string, account string) (err error) {
	err = dao.DB.Model(&input).Where("account_id = ?", account_id).Update("account", account).Error

	return err
}

func Deleted(input model.Table, account_id string, deleted bool) (err error) {
	if deleted {
		err = dao.DB.Delete(&input, &account_id).Error
	} else {
		err = dao.DB.Model(&input).Where("account_id = ?", account_id).Update("is_deleted", true).Error
	}

	return err
}
