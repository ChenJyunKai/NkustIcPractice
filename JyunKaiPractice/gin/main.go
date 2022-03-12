package main  
import (
    "github.com/gin-gonic/gin"
    "net/http"
	"errors"
)

var Userdata map[string]string //[key]值

func main(){
	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")
	server.Static("/assets", "./template/assets")
	server.GET("/", LoginPage)
	server.POST("/", LoginAuth)
	server.Run(":8888")
}

//Get
func LoginPage(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", nil)
}

//Post
func LoginAuth(c *gin.Context){
	var username,password string
	if in, isexist := c.GetPostForm("username"); isexist && in != ""{
		username = in 
	}else {
		c.HTML(http.StatusBadRequest,"login.html", gin.H{
			"error" : errors.New("必須輸入使用者"),
		})
		return
	}
	if in, isexist := c.GetPostForm("password"); isexist && in != ""{
		password = in 
	}else {
		c.HTML(http.StatusBadRequest,"login.html", gin.H{
			"error" : errors.New("必須輸入密碼"),
		})
		return
	}
	if err := Auth(username,password); err == nil{
		c.HTML(http.StatusOK,"login.html", gin.H{
			"error" : errors.New("登入成功"),
		})
		return
	}else {
		c.HTML(http.StatusUnauthorized,"login.html", gin.H{
			"error" : err,
		})
		return
	} 
}


func init(){
	Userdata = map[string]string{
		"corn" : "123",
	}
}

func CheckUserIsExist(username string) bool{
	_, isexist := Userdata[username] //判斷有無key值
	return isexist
}

func CheckPassword(pw string,pw_currect string) error{
	if pw == pw_currect{
		return nil //空值
	}else{
		return errors.New("password is not exist")
	}
}

func Auth(username string,password string) error {
	isexist := CheckUserIsExist(username)
	if isexist{
		return CheckPassword(Userdata[username],password)
	}else{
		return errors.New("username is not exist")
	}
}