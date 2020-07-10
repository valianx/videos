package handlers

import (
	"fmt"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"github.com/valianx/videos/config"
	usuarios "github.com/valianx/videos/internal/domain/models/usuario"
	"github.com/valianx/videos/utils"
)

type login struct {
	Nombre   string `form:"nombre" json:"nombre" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func UserLogin(password string, email string) bool { // Get model if exist
	var user usuarios.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	if utils.CheckPasswordHash(password, user.Password) {
		return true
	} else {
		return false
	}
}
