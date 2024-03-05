package main

import (
	"crowdfunding_app/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:Horse_6007@tcp(127.0.0.1:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to database")

	var users []user.User

	db.Find(&users)

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
		fmt.Println("=============")
	}
}
