package usuarios

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Nombre    string `json:"nombre"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
