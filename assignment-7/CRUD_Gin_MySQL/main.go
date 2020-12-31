// this is practiced program
//assignment is crud_gin.go anf crud_gorilla.go

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Rollno string `json:"rollno"`
	Age    string `json:"age"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "golang"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	router := gin.Default()

	router.POST("/create", func(ctx *gin.Context) {
		var s student
		if ctx.BindJSON(&s) == nil {
			ctx.JSON(200, gin.H{
				"name":   s.Name,
				"rollno": s.Rollno,
				"age":    s.Age,
			})
			db := dbConn()
			insert, err := db.Query("INSERT INTO student(name, rollno, age) VALUES(?,?,?)", s.Name, s.Rollno, s.Age)
			if err != nil {
				panic(err.Error())
			}
			defer insert.Close()
			// insert.Exec(s.Name, s.Rollno, s.Age)
			fmt.Printf("name: %s, rollno: %s, age: %s", s.Name, s.Rollno, s.Age)
		}
	})

	router.PUT("/update", func(ctx *gin.Context) {
		var s student
		if ctx.BindJSON(&s) == nil {
			ctx.JSON(200, gin.H{
				"name":   s.Name,
				"rollno": s.Rollno,
				"age":    s.Age,
			})
			db := dbConn()
			update, err := db.Prepare("UPDATE student SET name=?, rollno=?, age=? Where id=?")
			if err != nil {
				panic(err.Error())
			}
			update.Exec(s.Name, s.Rollno, s.Age, s.Id)
		}
	})

	router.GET("/read", func(ctx *gin.Context) {
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM student")
		if err != nil {
			panic(err.Error())
		}
		var id int
		var name, rollno, age string
		for selDB.Next() {
			err = selDB.Scan(&id, &name, &rollno, &age)
			ctx.JSON(200, gin.H{
				"id":     id,
				"name":   name,
				"rollno": rollno,
				"age":    age,
			})
			fmt.Printf("name: %s, rollno: %s, age: %s", name, rollno, age)
			if err != nil {
				panic(err.Error())
			}
		}
	})

	router.DELETE("/delete", func(ctx *gin.Context) {
		var s student
		if ctx.BindJSON(&s) == nil {
			db := dbConn()
			del, err := db.Prepare("DELETE FROM student WHERE name=?")
			if err != nil {
				panic(err.Error())
			}
			del.Exec(s.Name)
			log.Println("DELETE")
			defer db.Close()
		}
	})

	router.Run(":8000")

}
