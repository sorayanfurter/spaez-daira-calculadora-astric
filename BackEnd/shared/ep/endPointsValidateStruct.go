package ep

import (
	"reflect"

	"ASTRIC/BackEnd/ASTRIC/helper/env"
	messages "ASTRIC/BackEnd/Messages"

	"github.com/go-playground/validator/v10"
)

type info struct {
	Tag   string       `json:"tags"`
	Tipe  reflect.Type `json:"type"`
	Val   interface{}  `json:"val"`
	Param string       `json:"param"`
}

/*
ValidateStruct Valida las estructuras
Ej:

	type MyStruct struct {
			String string `validate:"is-awesome,required"`
	}

https://github.com/go-playground/validator

@s: Requiere la estructura a validar

Returns:

	Interface: con objete del detalle de los errore
	Bool: True si ubo error False si no hay error
*/
func ValidateStruct(s interface{}) (interface{}, bool) {
	errores := make(map[string]interface{})
	notificacion := !env.GetEnvGeneral().Notificaciones
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		if notificacion {
			errores["MSG"] = messages.GetMessagesValid()
			for _, err := range err.(validator.ValidationErrors) {
				i := info{}
				i.Tag = err.Tag()
				i.Param = err.Param()
				i.Tipe = err.Type()
				i.Val = err.Value()
				errores[err.Namespace()] = i
			}
			return errores, true
		}
		errores["DATOS"] = "Los datos enviados no son validos"
		return errores, true
	}
	return nil, false
}

/*
ValidateStructs Valida las estructuras
Ej:

	type MyStruct struct {
			String string `validate:"is-awesome,required"`
	}

https://github.com/go-playground/validator

@s: Requiere la estructura a validar

Returns:

	Interface: con objete del detalle de los errore, si no hay errores es Nil
*/
func ValidateStructs(s interface{}) interface{} {
	errores := make(map[string]interface{})
	notificacion := !env.GetEnvGeneral().Notificaciones
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		if notificacion {
			errores["MSG"] = messages.GetMessagesValid()
			for _, err := range err.(validator.ValidationErrors) {
				i := info{}
				i.Tag = err.Tag()
				i.Param = err.Param()
				i.Tipe = err.Type()
				i.Val = err.Value()
				errores[err.Namespace()] = i
			}
			return errores
		}
		errores["DATOS"] = "Los datos enviados no son validos"
		return errores
	}
	return nil
}
