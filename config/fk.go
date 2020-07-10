package config

import (
	"github.com/jinzhu/gorm"
	//usuarios "github.com/valianx/videos/internal/domain/models/usuario"
)

func foreingKey(database *gorm.DB)*gorm.DB{
//	database.Model(&usuarios.User{}).AddForeignKey("cargo", "cargos(id)", "RESTRICT", "RESTRICT")

	return database
}