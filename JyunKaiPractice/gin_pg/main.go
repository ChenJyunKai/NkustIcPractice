package main

import(
	"fmt"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
	_ "github.com/lib/pq"
	_ "strconv"
)

type Student struct{
	Id int `json "id"`
	Name string `json "name"`
	Age int `json "age"`
}

func main(){
	db,err := sql.Open("postgres","user=postgres password=usa960092 dbname=corn sslmode=disable")
	if err != nil{
		fmt.Println("鏈結error",err.Error())
		return
	}
	defer db.Close()
	fmt.Println("鏈結success")
	err = db.Ping()
	if err != nil{
		fmt.Println("錯誤",err.Error())
		return
	}

	server := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	server.Use(CORSMiddleware())
	server.LoadHTMLGlob("template/*")

	api := server.Group("api")
	{
		api.GET("/select", func(c *gin.Context){
			students := selectstudent(db)
			c.JSON(http.StatusOK,gin.H{
				"students" : students,
			})
		})

		// api.POST("/insert",func(c *gin.Context){
		// 	id,_:= strconv.
		// })
	}
	server.Run()
}

func selectstudent(db *sql.DB)(students []Student){
	var rows *sql.Rows
	rows,err := db.Query("Select * from student")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next(){
		student := Student{}
		err = rows.Scan(&student.Id ,&student.Name ,&student.Age)
		students = append(students, student)
	}
	return
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}