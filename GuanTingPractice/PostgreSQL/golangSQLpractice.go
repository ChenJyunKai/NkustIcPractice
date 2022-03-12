package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=17787 dbname=test sslmode=disable")
    checkErr(err)

	//---------------------DB SELECT--------------------------
	// rows, err := db.Query("SELECT * FROM company")
    // checkErr(err)
	// for rows.Next() {
    //     var empid int
    //     var empname string
    //     var brithday string
    //     err = rows.Scan(&empid, &empname, &brithday)
    //     checkErr(err)

    //     fmt.Println(empid)
    //     fmt.Println(empname)
    //     fmt.Println(brithday)
    // } 

	//---------------------DB ADD--------------------------
	// stmt, err := db.Prepare("INSERT INTO company(empname,brithday) VALUES($1,$2)")
    // checkErr(err)
    // res, err := stmt.Exec("Cool man", "2012-12-09")
    // checkErr(err)
	// affect, err := res.RowsAffected() // database affect return 1 else 0
    // checkErr(err)

    // fmt.Println(affect)

	//---------------------DB UPDATE--------------------------
	// stmt, err := db.Prepare("update company set empname=$1 where empid=$2")
	// checkErr(err)
	// res, err := stmt.Exec("Guanting",2)
	// affect, err := res.RowsAffected()
    // checkErr(err)

    // fmt.Println(affect)

	//---------------------DB DELETE--------------------------
	stmt, err :=db.Prepare("delete from company where empid=$1")
	checkErr(err)
	res, err := stmt.Exec(4)
	checkErr(err)
	affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    db.Close()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}