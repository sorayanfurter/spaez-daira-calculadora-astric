package migration

import (
	"ASTRIC/BackEnd/config"
	"errors"

	"gorm.io/gorm"
)

func createUserDev(con *gorm.DB) error {

	User := config.LoadConfigUser()

	User.IDInterno = 0

	result := con.Delete(&User, "id_interno", 0)

	if result.Error != nil {
		return errors.New("usuario dev: " + result.Error.Error())
	}

	result = con.Create(&User)

	if result.Error != nil {
		return errors.New("usuario dev: " + result.Error.Error())
	}

	if result.RowsAffected < 1 {
		return errors.New("usuario dev: El usuario no pudo ser creado")
	}

	return nil
}

func createUserStructBD(con *gorm.DB) error {

	var usuarios config.Usuarios
	err := con.AutoMigrate(
		&usuarios,
	)
	if err != nil {
		return errors.New("creacion de struct usuario: " + err.Error())
	}

	return createUserDev(con)
}
