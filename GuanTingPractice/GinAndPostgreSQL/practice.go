package main
import (
	"github.com/gin-gonic/gin"
	"strconv"
	"database/sql"
    "fmt"
    _ "github.com/lib/pq"
	"github.com/gin-contrib/cors"
)
//Lowercase will not be converted to JSON
type Employee struct {
	Empid int `json:"empid"`
	Empname string `json:"empname"`
	Brithday string `json:"brithday"`
}

func main(){
	db, err := sql.Open("postgres", "user=postgres password=17787 dbname=test sslmode=disable")
	checkErr(err)
	defer db.Close()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	company := router.Group("/company")
	{
		company.GET("/select",func(c *gin.Context){
			employees := selectemployee(db)
			c.JSON(200, gin.H{
				"message":"查詢成功",
				"result":employees,
			})
		})
		company.POST("/insert",func(c *gin.Context){
			json := Employee{}
			c.BindJSON(&json)
			employee :=  Employee{0,json.Empname,json.Brithday}
			insertemployee(db, employee)
			c.JSON(200, gin.H{
				"message":"新增成功",
			})
		})
		company.POST("/update",func(c *gin.Context){
			empid,_ := strconv.Atoi(c.PostForm("empid"))
			empname := c.PostForm("empname")
			brithday := c.PostForm("brithday")
			employee :=  Employee{empid,empname,brithday}
			fmt.Println(employee)
			updateemployee(db, employee)
			c.JSON(200, gin.H{
				"message":"修改成功",
			})
		})

		company.POST("/delete",func(c *gin.Context){
			empid,_ := strconv.Atoi(c.PostForm("empid"))
			deleteemployee(db,empid)
			c.JSON(200, gin.H{
				"message":"刪除成功",
			})
		})
	}
 
	router.Run()
}


func selectemployee(db *sql.DB) (employees []Employee){
	rows, err := db.Query("SELECT * FROM company")
	checkErr(err)
	for rows.Next() { 
		employee := Employee{}
		err = rows.Scan(&employee.Empid, &employee.Empname, &employee.Brithday) 
		checkErr(err)
		employees = append(employees, employee) 
	} 
	return
}

func insertemployee(db *sql.DB,employee Employee) {
	stmt, err := db.Prepare("INSERT INTO company(empname,brithday) VALUES($1,$2)")
	checkErr(err)
	stmt.Exec(employee.Empname,employee.Brithday)
	return
}

func updateemployee(db *sql.DB,employee Employee){
	stmt, err := db.Prepare("UPDATE company set empname=$1,brithday=$2 where empid=$3")
	checkErr(err)
	stmt.Exec(employee.Empname,employee.Brithday,employee.Empid)
	return
}

func deleteemployee(db *sql.DB,empid int){
	stmt, err := db.Prepare("DELETE from company where empid=$1")
	checkErr(err)
	stmt.Exec(empid)
	return
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}