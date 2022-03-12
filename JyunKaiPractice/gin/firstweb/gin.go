package main  
import (
    "github.com/gin-gonic/gin"
    "net/http"
)
func main(){
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.GET("/", test)
	server.Run(":8888")
}

func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "my first web"
	c.HTML(http.StatusOK, "index.html", data)
}

type IndexData struct{
	Title string
	Content string
}
