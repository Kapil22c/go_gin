package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	Id      int    `json:id`
	Name    string `json:name`
	price   string `json:price`
	quality string `json:quality`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "order_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func main() {
	router := gin.Default()

	router.POST("/add", func(c *gin.Context) {

		name := c.Query("name")
		price := c.Query("price")
		quality := c.Query("quality")

		c.JSON(200, gin.H{
			"name":    name,
			"price":   price,
			"quality": quality,
		})
		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO product(name, price, quality) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, price, quality)
		fmt.Printf("name: %s; price: %s; quality: %s", name, price, quality)
	})

	router.PUT("/update", func(c *gin.Context) {
		id1 := c.Query("id")
		name := c.Query("name")
		price := c.Query("price")
		quality := c.Query("quality")

		c.JSON(200, gin.H{
			"name":    name,
			"price":   price,
			"quality": quality,
		})
		db := dbConn()
		upForm, err := db.Prepare("UPDATE product SET name=?, price=?, quality=? Where id=?")
		if err != nil {
			panic(err.Error())
		}
		upForm.Exec(name, price, quality, id1)
		fmt.Printf("name: %s; price: %s; quality: %s", name, price, quality)
	})

	router.GET("/GET", func(c *gin.Context) {
		id := c.Query("id")
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM product WHERE id=?", id)
		if err != nil {
			panic(err.Error())
		}

		var name, price, quality string
		for selDB.Next() {

			err = selDB.Scan(&id, &name, &price, &quality)
			if err != nil {
				panic(err.Error())
			}
		}
		fmt.Printf("name: %s; price: %s; quality: %s", name, price, quality)

		c.JSON(200, gin.H{
			"id":      id,
			"name":    name,
			"price":   price,
			"quality": quality,
		})

	})

	router.DELETE("/delete", func(c *gin.Context) {
		var p product
		if c.BindJSON(&p) == nil {
			db := dbConn()
			delForm, err := db.Prepare("DELETE FROM product WHERE name=?")
			if err != nil {
				panic(err.Error())
			}
			delForm.Exec(p.Name)
			log.Println("DELETE")
			defer db.Close()
		}

	})

	router.Run(":8080")
}
