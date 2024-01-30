package config

import (
	"gorm.io/gorm"
)

// Migration funcion para realizar migrado de la BD nesecita configuracion
func Migration(con *gorm.DB) error {

	//var tabla my.Tabla

	if err := con.AutoMigrate(

	// Modelos a migrar
	//&tabla,
	); err != nil {
		return err
	}

	return nil

}
