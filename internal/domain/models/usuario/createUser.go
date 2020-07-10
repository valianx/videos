package usuarios

type CreateUserInput struct {
	Nombre   string `json:"nombre" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
