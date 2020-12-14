package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Player struct {
	Id      int    `json:id`
	Name    string `json:name`
	Role    string `json:role`
	Matches int    `json:matches`
	Age     int    `json:age`
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
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		//render only file, must full name with extension
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM player ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		player := Player{}
		res := []Player{}
		for selDB.Next() {
			var id, matches, age int
			var name, role string
			err = selDB.Scan(&id, &name, &role, &matches, &age)
			if err != nil {
				panic(err.Error())
			}
			player.Id = id
			player.Name = name
			player.Role = role
			player.Matches = matches
			player.Age = age
			res = append(res, player)
		}
		//var a = "hello words"
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Home Page!!", "a": res})
	})

	r.GET("/add", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "add.html", gin.H{"title": "Add player!!"})
	})

	r.POST("/insert", func(ctx *gin.Context) {
		//render only file, must full name with extension
		var name, role string
		var matches, age string
		name = ctx.Request.FormValue("name")
		role = ctx.Request.FormValue("role")
		matches = ctx.Request.FormValue("matches")
		age = ctx.Request.FormValue("age")

		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO player (name, role, matches, age) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, role, matches, age)
		// ctx.HTML(http.StatusOK, "updated.html", gin.H{"title": "Player"})

		selDB, err := db.Query("SELECT * FROM player ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		player := Player{}
		res := []Player{}
		for selDB.Next() {
			var id, matches, age int
			var name, role string
			err = selDB.Scan(&id, &name, &role, &matches, &age)
			if err != nil {
				panic(err.Error())
			}
			player.Id = id
			player.Name = name
			player.Role = role
			player.Matches = matches
			player.Age = age
			res = append(res, player)
		}
		//var a = "hello words"
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Home Page!!", "a": res})
	})

	r.Run(":8080")

}
