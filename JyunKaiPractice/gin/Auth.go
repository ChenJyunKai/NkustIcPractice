// package main  
// import (
// 	"errors"
// )

// var Userdata map[string]string //[key]值

// func init(){
// 	Userdata = map[string]string{
// 		"corn" : "123",
// 	}
// }

// func CheckUserIsExist(username string) bool{
// 	_, isexist := Userdata[username] //判斷有無key值
// 	return isexist
// }

// func CheckPassword(pw string,pw_currect string) error{
// 	if pw == pw_currect{
// 		return nil //空值
// 	}else{
// 		return errors.New("password is not exist")
// 	}
// }

// func Auth(username string,password string) error {
// 	isexist := CheckUserIsExist(username)
// 	if isexist{
// 		return CheckPassword(Userdata[username],password)
// 	}else{
// 		return errors.New("username is not exist")
// 	}
// }