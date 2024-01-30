package usuarios

import "ASTRIC/BackEnd/ASTRIC/users/model"

// Template template de atributos requeridos
/*
Atributos:
	Admin  bool   `json:"admin,omitempty" gorm:"column:admin" validate:"required"`
	Activo bool   `json:"activo,omitempty" gorm:"column:activo"`
	Reset  bool   `json:"reset,omitempty" gorm:"column:reset"`
	Hash   string `json:"hash,omitempty" gorm:"column:hash"`
*/
type Template model.TemplateUsuario
