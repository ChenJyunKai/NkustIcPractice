package practices

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	entity "bitbucket.org/henry111666888/hs_gorm/entity/practices"
	"bitbucket.org/henry111666888/hs_gorm/pkg/log"
	"bitbucket.org/henry111666888/hs_gorm/pkg/util"

	practices "bitbucket.org/henry111666888/hs_gorm/structure/practices"
)

type Food struct {
	Name  string
	Price int
}

func Created(input practices.Table) (err error) {
	easy_json, _ := os.Open("E:/Golang/hs_gorm/easy_json.json")
	easybyteValue, _ := ioutil.ReadAll(easy_json)
	json.Unmarshal(easybyteValue, &input.Easy_json)

	complex_json, _ := os.Open("E:/Golang/hs_gorm/complex_json.json")
	complexbyteValue, _ := ioutil.ReadAll(complex_json)
	json.Unmarshal(complexbyteValue, &input.Complex_json)

	input.Uuid4_field = util.GenerateUUID()
	input.Str_field = "test"
	input.Int_field = 1
	input.Is_bool = false
	input.Time_field = time.Now()
	input.Date_field = time.Now()

	if err != entity.Created(input) {
		log.Error(err)
	}
	return err
}
