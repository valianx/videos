package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	usuarios "github.com/valianx/videos/internal/domain/models/usuario"
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
		Host:     "ec2-54-211-210-149.compute-1.amazonaws.com",
		Port:     5432,
		User:     "udezkmvezvjrjp",
		Password: "b94c54f9b109d00eb05c161bd50b6499c02bffdad640e4f806cd47bd0b40f814",
		DBname:   "d1q9kfaqcvk40b",
	}

	connect := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		prod.Host, prod.Port, prod.User, prod.Password, prod.DBname)

	database, err := gorm.Open("postgres", connect)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}
	migrar(database)
	DB = foreingKey(database)
}
func Drop() {
	connect := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	database, err := gorm.Open("postgres", connect)

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.DropTable(&usuarios.User{})

}

func migrar(database *gorm.DB) {
	database.AutoMigrate(&usuarios.User{})

}
