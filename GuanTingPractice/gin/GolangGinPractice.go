package main
import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)
func main(){
	router := gin.Default()
	router.GET("./ping",func(c*gin.Context){
		c.JSON(200,gin.H{
			"message":"Hello",
		})
	}) 

	router.POST("./ping/:id",func(c*gin.Context){ 
		id := c.Param("id")
		c.JSON(200,gin.H{
			"id":id,
		})
	})

	router.POST("./ping",func(c*gin.Context){
		id := c.Query("id")
		c.JSON(200,gin.H{
			"id":id,
		})
	})
	
	router.POST("/",func(c *gin.Context){
		body:=c.Request.Body
		x, _ := ioutil.ReadAll(body)
		c.JSON(200,gin.H{
			"message":string(x),
		})
	})
	
	router.Run()
}  