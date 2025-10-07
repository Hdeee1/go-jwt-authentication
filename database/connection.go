package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(){
	_, err := gorm.Open(mysql.Open("root:dbmysql@/db_jwt_auth"), &gorm.Config{})
	
	if err != nil {
		panic("could not connect to the database")
	}
}