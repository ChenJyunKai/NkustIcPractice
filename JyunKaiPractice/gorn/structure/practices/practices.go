package practices

import "time"

type Table struct {
	Uuid4_field  string                 `gorm:"column:uuid4_field;uuid_generate_v4();type:UUID;" json:"account_id,omitempty"`
	Str_field    string                 `gorm:"column:str_field;type:VARCHAR" json:"account,omitempty"`
	Int_field    int                    `gorm:"column:int_field;type:INTEGER" json:"name,omitempty"`
	Is_bool      bool                   `gorm:"column:is_bool;type:BOOL;false" json:"is_deleted,omitempty"`
	Time_field   time.Time              `gorm:"column:time_field;type:TIMESTAMP;" json:"time_field"`
	Date_field   time.Time              `gorm:"column:date_field;type:DATE;" json:"date_field"`
	Easy_json    map[string]interface{} `gorm:"column:easy_json;type:JSONB" json:"esay_json,omitempty"`
	Complex_json map[string]interface{} `gorm:"column:complex_json;type:JSONB" json:"complex_json,omitempty"`
}

func (a *Table) TableName() string {
	return "practices"
}
