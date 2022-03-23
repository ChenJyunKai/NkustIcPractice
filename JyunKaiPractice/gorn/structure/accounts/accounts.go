package accounts

import "time"

type Table struct {
	//帳號ID
	AccountID string `gorm:"primaryKey;column:account_id;uuid_generate_v4();type:UUID;" json:"account_id,omitempty"`
	//帳號: email
	Account string `gorm:"column:account;type:VARCHAR" json:"account,omitempty"`
	//姓名
	Name string `gorm:"column:name;type:VARCHAR" json:"name,omitempty"`
	//密碼
	Password string `gorm:"column:pwd;type:VARCHAR" json:"pwd,omitempty"`
	//創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	//是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:BOOL;false" json:"is_deleted,omitempty"`
}

type Base struct {
	//帳號ID
	AccountID string `json:"account_id,omitempty"`
	//帳號: email
	Account string `json:"account,omitempty"`
	//姓名
	Name string `json:"name,omitempty"`
	//密碼
	Password string `json:"pwd,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at"`
	//是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
}

type Created struct {
	//帳號: email
	Account string `json:"account,omitempty" binding:"required,email" validate:"required"`
	//姓名
	Name string `json:"name" binding:"required" validate:"required"`
	//密碼
	Password string `json:"pwd" binding:"required" validate:"required"`
}

//read one 運作用的struct
type Single struct {
}

//通常是前端進來(非基礎CRUD)的struct
type Field struct {
}

//送往前端的讀全部struct
type List struct {
}

//更新時須送進來的內部結構
type Updated struct {
}

func (a *Table) TableName() string {
	return "accounts"
}
