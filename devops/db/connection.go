package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func GetConnection(engine, username, password, name string) *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, name)
	connection, err := gorm.Open(engine, dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return connection
}
