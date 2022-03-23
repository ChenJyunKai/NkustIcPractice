package practices

import (
	"bitbucket.org/henry111666888/hs_gorm/pkg/dao"
	practices "bitbucket.org/henry111666888/hs_gorm/structure/practices"
)

func Created(input practices.Table) (err error) {
	err = dao.DB.Create(&input).Error

	return err
}
