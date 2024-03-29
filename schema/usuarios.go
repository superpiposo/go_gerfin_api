package schema

import (
	"gorm.io/gorm"
)

type Usuarios struct {
	gorm.Model
	Nome      string
	Sobrenome string
	Cpf       string
	Senha     string
}
