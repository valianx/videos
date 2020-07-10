package usuarios


type UpdateUserInput struct {
	Nombre   	string `json:"nombre"`
	Email    	string `json:"email"`
	Password 	string `json:"password"`
}