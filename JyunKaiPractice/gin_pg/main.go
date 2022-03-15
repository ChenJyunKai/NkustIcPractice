package main

import(
	"fmt"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/lib/pq"
)

type Student struct{
	id int
	name string
	age int
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
	server.LoadHTMLGlob("template/*")

	api := server.Group("api")
	{
		api.GET("/", func(c *gin.Context){
			students := selectstudent(db)
			c.HTML(http.StatusOK,"index.html",gin.H{
				"students" : students,
			})
		})
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
		err = rows.Scan(&student.id ,&student.name ,&student.age)
		students = append(students, student)
	}
	return
}