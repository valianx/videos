package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

type DBdata struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1346"
	dbname   = "gorm"
)

func ConnectDataBaseProduction() {

	prod := DBdata{
		Host:     "ec2-35-172-73-125.compute-1.amazonaws.com",
		Port:     5432,
		User:     "zxaesyjblfuogt",
		Password: "bf80be2a90e6e9a676451b507fd8b7df8349f2c094527b6da9d9baf7873f7549",
		DBname:   "d8764sj6q4a3pu",
	}

	connect := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		prod.Host, prod.Port, prod.User, prod.Password, prod.DBname)

	database, err := gorm.Open("postgres", connect)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}
	//migrar(database)
	DB = foreingKey(database)
}

