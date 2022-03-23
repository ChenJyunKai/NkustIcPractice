package accounts

import (
	"time"

	entity "bitbucket.org/henry111666888/hs_gorm/entity/account"
	"bitbucket.org/henry111666888/hs_gorm/pkg/log"
	"bitbucket.org/henry111666888/hs_gorm/pkg/util"

	model "bitbucket.org/henry111666888/hs_gorm/structure/accounts"
)

func Created(input model.Table) (err error) {
	input.AccountID = util.GenerateUUID()
	input.CreatedAt = time.Now()
	input.IsDeleted = false

	if err != entity.Created(input) {
		log.Error(err)
	}
	return err
}

func List(input model.Table) (err error) {
	if err != entity.List(input) {
		log.Error(err)
	}
	return err
}

func GetByAccount(input model.Table, account_id string) (err error) {
	if err != entity.GetByAccount(input, account_id) {
		log.Error(err)
	}
	return err
}

func Updated(input model.Table, account_id string, account string) (err error) {
	if err != entity.Updated(input, account_id, account) {
		log.Error(err)
	}
	return err
}

func Deleted(input model.Table, account_id string, deleted bool) (err error) {
	if err != entity.Deleted(input, account_id, deleted) {
		log.Error(err)
	}
	return err
}
