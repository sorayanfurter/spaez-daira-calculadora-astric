package model

// TemplateUsuario Modelo genaral
/*
swagger:model Usuarios
Atributos:

	Admin  bool   `json:"admin,omitempty" gorm:"column:admin" validate:"required"`
	Activo bool   `json:"activo,omitempty" gorm:"column:activo"`
	Reset  bool   `json:"reset,omitempty" gorm:"column:reset"`
	Hash   string `json:"hash,omitempty" gorm:"column:hash"`
*/
type TemplateUsuario struct {
	IDInterno int    `json:"id_interno,omitempty" gorm:"column:id_interno"`
	Admin     bool   `json:"admin,omitempty" gorm:"column:admin"`
	Activo    bool   `json:"activo,omitempty" gorm:"column:activo"`
	Reset     bool   `json:"reset,omitempty" gorm:"column:reset"`
	Hash      string `json:"hash,omitempty" gorm:"column:hash"`
}
