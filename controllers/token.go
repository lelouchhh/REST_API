package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type user struct {
	login string
	password string
}
// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *sql.DB
func InitiateDB(db *sql.DB) {
	dbConnect = db
}
func Token(c *gin.Context){
	acc := &user{login: c.Query("login"), password: c.Query("password")}
	var suc bool
	rows, err := dbConnect.Exec("call auth.set_reg ($1,$2);",acc.login, acc.password )
	fmt.Println(rows)
	if err != nil {
		panic(err)
	}
	fmt.Println(suc)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Todo",
		"data": acc.password,
	})
	return
}