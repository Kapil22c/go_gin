package main

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
	//"fmt"
	//"reflect"
)

////

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/golang?sslmode=disable")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(User{}, "User2").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create table failed")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetUsers(c *gin.Context) {
	var users []User
	_, err := dbmap.Select(&users, "SELECT * FROM user2")
	if err == nil {
		c.JSON(200, users)
	} else {
		c.JSON(404, gin.H{"error": "no user(s) into the table"})
	}
}

func GetUser(c *gin.Context) {
	param_id := c.Params.ByName("id")
	id, _ := strconv.Atoi(param_id)
	//original//
	//var user User
	//err := dbmap.SelectOne(&user, "SELECT * FROM user2 WHERE id=?", id)
	//new//
	obj, err := dbmap.Get(User{}, id)
	//c.JSON(404, gin.H{"obj" : obj})
	user := obj.(*User)
	if err == nil {
		user_id, _ := strconv.ParseInt(param_id, 0, 64)
		content := &User{
			Id:        user_id,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
		//c.JSON(404, gin.H{"error" : err})
	}
}

func PostUser(c *gin.Context) {
	var params User
	c.Bind(&params)

	if params.Firstname != "" && params.Lastname != "" {
		user := &User{0, params.Firstname, params.Lastname}
		err := dbmap.Insert(user)
		if err == nil {
			content := &User{
				Id:        user.Id,
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
			}
			c.JSON(201, content)
		} else {
			c.JSON(422, gin.H{"error": err})
		}
	} else {
		c.JSON(422, gin.H{"error": "fields are empty"})
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	err := dbmap.SelectOne(&user, "SELECT * FROM user2 WHERE id=?", id)
	if err == nil {
		var json User
		c.Bind(&json)
		user_id, _ := strconv.ParseInt(id, 0, 64)
		user := User{
			Id:        user_id,
			Firstname: json.Firstname,
			Lastname:  json.Lastname,
		}
		if user.Firstname != "" && user.Lastname != "" {
			_, err = dbmap.Update(&user)
			if err == nil {
				c.JSON(200, user)
			} else {
				checkErr(err, "Updated failed")
			}
		} else {
			c.JSON(422, gin.H{"error": "fields are empty"})
		}
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	err := dbmap.SelectOne(&user, "SELECT id FROM user2 WHERE id=?", id)
	if err == nil {
		_, err = dbmap.Delete(&user)
		if err == nil {
			c.JSON(200, gin.H{"id #" + id: " deleted"})
		} else {
			checkErr(err, "Delete failed")
		}
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}

type User struct {
	Id        int64  `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.POST("/users", PostUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
	}
	r.GET("/", index)
	r.Run(":8081")
}
