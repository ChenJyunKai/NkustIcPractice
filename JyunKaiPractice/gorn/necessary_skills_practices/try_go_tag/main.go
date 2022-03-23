package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validate struct {
	Phone   string `form:"phone" json:"phone" validate:"required,len=10"`
	Pwd     string `form:"pwd" json:"pwd" validate:"required,max=20,min=6"`
	SamePwd string `form:"same_pwd" json:"same_pwd" validate:"required,eqfield=Pwd"`
	Code1   string `form:"code1" json:"code1" validate:"required"`
	Code2   string `form:"code2" json:"code2" validate:"required,excludesall=0x20"`
	Email   string `form:"email" json:"email" validate:"required,email"`
}

func main() {
	//請在此填入可以正確被驗證的值
	validates := &Validate{
		Phone:   "",
		Pwd:     "",
		SamePwd: "",
		Code1:   "",
		Code2:   "",
		Email:   "",
	}

	validate := validator.New()
	err := validate.Struct(validates)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
	}
}
