package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"github.com/valianx/videos/config"
	usuarios "github.com/valianx/videos/internal/domain/models/usuario"
	"github.com/valianx/videos/utils"
	_ "golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Mostrar Users godoc
// @tags CRUD users
// @Summary Mostrar todos los users
// @Description Muestra todos los users que pueden tener los usuarios
// @Accept  json
// @Produce  json
// @Success 200 {object} usuarios.User
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users [get]
func FindUsers(c *gin.Context) {
	var users []usuarios.User
	config.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// Crear Users godoc
// @ID post-user
// @tags CRUD users
// @Summary Crear nuevo user
// @Description Crea un nuevo user para los usuarios
// @Accept  json
// @Produce  json
// @Param User body usuarios.CreateUserInput true "Modelo user"
// @Success 200 {object} usuarios.User
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users [post]
func CreateUser(c *gin.Context) {
	input := usuarios.CreateUserInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pass, _ := utils.HashPassword(input.Password)
	// Create User
	user := usuarios.User{
		Nombre:    input.Nombre,
		Email:     input.Email,
		Password:  pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Mostrar Users godoc
// @ID get-user
// @tags CRUD users
// @Summary Muestra un user
// @Description Con su id se muestra un user seleccionado
// @Accept  json
// @Produce  json
// @Param id path int true "user"
// @Success 200 {object} usuarios.User
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/{id} [get]
func FindUser(c *gin.Context) { // Get model if exist
	var user usuarios.User

	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @tags CRUD users
// @Summary editar user
// @Description Edita un user con su id
// @ID patch-user
// @Accept  json
// @Produce  json
// @Param id path int true "user"
// @Param User body usuarios.UpdateUserInput true "Modelo user"
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/{id} [patch]
func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user usuarios.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input usuarios.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @tags CRUD users
// @Summary Eliminar user
// @Description Elimina un user con su id
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param id path int true "User"
// @Success 200 {object} usuarios.User
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user usuarios.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
