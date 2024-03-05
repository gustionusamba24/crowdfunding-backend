package main

import (
	"crowdfunding_app/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	//dsn := "root:Horse_6007@tcp(127.0.0.1:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//
	//fmt.Println("Connected to database")
	//
	//var users []user.User
	//
	//db.Find(&users)
	//
	//for _, user := range users {
	//	fmt.Println(user.Name)
	//	fmt.Println(user.Email)
	//	fmt.Println("=============")
	//}

	router := gin.Default()
	router.GET("/handler", handler)
	router.Run(":9000")
}

func handler(c *gin.Context) {
	dsn := "root:Horse_6007@tcp(127.0.0.1:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
