package res

import (
	"ASTRIC/BackEnd/ASTRIC/global"
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/ASTRIC/helper/errors"
	ASTRICmodels "ASTRIC/BackEnd/ASTRIC/pModels"
	messages "ASTRIC/BackEnd/Messages"
	"encoding/json"
	"net/http"
)

/*
Response funcion que genera la respuesta de la API

	@tag: Es la tarea o funcion que se solicito ej: Nuevo Usuario
	@datos: Son los datos que se utilizaron para realizar el tag solisitado
	@w: Corresponde al ResponseWriter "w http.ResponseWriter"
*/
func Response(tag string, errores interface{}, datos interface{}, w http.ResponseWriter) {
	defer errors.ErrorControl("shared/helpers/helpersResponse")
	if errores != nil {
		createRespnse(tag, errores, datos, w, http.StatusBadRequest, false, "")
	} else {
		createRespnse(tag, errores, datos, w, http.StatusOK, true, "")
	}
}

/*
ResponseExpire funcion que genera la respuesta de la API

	@tag: Es la tarea o funcion que se solicito ej: Nuevo Usuario
	@w: Corresponde al ResponseWriter "w http.ResponseWriter"
*/
func ResponseExpire(tag string, w http.ResponseWriter, msg string) {
	defer errors.ErrorControl("shared/helpers/helpersResponse")
	createRespnse(tag, nil, nil, w, http.StatusNetworkAuthenticationRequired, true, msg)
}

/*
ResponseLogin funcion que genera la respuesta de la API

	@tag: Es la tarea o funcion que se solicito ej: Nuevo Usuario
	@w: Corresponde al ResponseWriter "w http.ResponseWriter"
*/
func ResponseLogin(tag string, w http.ResponseWriter) {
	defer errors.ErrorControl("shared/helpers/helpersResponse")
	createRespnse(tag, nil, nil, w, http.StatusOK, true, "")
}

/*
ResponseMSG funcion que genera la respuesta de la API incorporando un mensaje especial
solo funcona con la variab de entrono NOTIFICCIONES en true

	@tag: Es la tarea o funcion que se solicito ej: Nuevo Usuario
	@datos: Son los datos que se utilizaron para realizar el tag solisitado
	@w: Corresponde al ResponseWriter "w http.ResponseWriter"
	@msg: Corresponde al mensaje el mismo aprecesra con 1 - !ATENCION¡ al comienzo
*/
func ResponseMSG(tag string, errores interface{}, datos interface{}, w http.ResponseWriter, msg string) {
	defer errors.ErrorControl("shared/helpers/helpersResponse")
	if errores != nil {
		createRespnse(tag, errores, datos, w, http.StatusBadRequest, false, msg)
	} else {
		createRespnse(tag, errores, datos, w, http.StatusOK, true, msg)
	}
}

/*
ResponseUnauthorized funcion que genera la respuesta para acciones no autorizadas

	@tag: Es la tarea o funcion que se solicito ej: Nuevo Usuario
	@datos: Son los datos que se utilizaron para realizar el tag solisitado
	@w: Corresponde al ResponseWriter "w http.ResponseWriter"
*/
func ResponseUnauthorized(tag string, errores interface{}, datos interface{}, w http.ResponseWriter) {
	defer errors.ErrorControl("shared/helpers/helpersResponseLogin")
	createRespnse(tag, errores, datos, w, http.StatusUnauthorized, false, "")
}

/*
ResponseUnauthorizedMSG funcion que genera la respuesta para acciones no autorizadas incorporando un mensaje especial
solo funcona con la variab de entrono NOTIFICCIONES en true y este desabilitado el tocken

	@tag: Es la tarea o funcion que se solicito ej: Nuevo Usuario
	@datos: Son los datos que se utilizaron para realizar el tag solisitado
	@w: Corresponde al ResponseWriter "w http.ResponseWriter"
	@msg: Corresponde al mensaje el mismo aprecesra con 1 - !ATENCION¡ al comienzo
*/
func ResponseUnauthorizedMSG(tag string, errores interface{}, datos interface{}, w http.ResponseWriter, msg string) {
	defer errors.ErrorControl("shared/helpers/helpersResponseLogin")
	if env.GetEnvGeneral().DisableToken {
		createRespnse(tag, errores, datos, w, http.StatusUnauthorized, false, msg)
		return
	}
	createRespnse(tag, errores, datos, w, http.StatusUnauthorized, false, "")
}

func createRespnse(tag string, errores interface{}, datos interface{}, w http.ResponseWriter, stausHead int, status bool, msg string) {

	defer errors.ErrorControl("shared/helpers/helpersCreateResponse")
	var response interface{}
	var msgs map[string]string

	if !env.GetEnvGeneral().Notificaciones {
		if messages.GetMessagesResp() != nil {
			msgs = messages.GetMessagesResp()
		}
		if msg != "" {
			msgs["1 - !ATENCION¡"] = msg
		}
	}

	switch stausHead {
	case http.StatusUnauthorized:
		response = ASTRICmodels.ResUnautorize{
			Error: errores,
			Code:  http.StatusUnauthorized,
			Msg:   msgs,
		}
	case http.StatusNetworkAuthenticationRequired:
		response = ASTRICmodels.ResUnautorize{
			Error: errores,
			Code:  http.StatusNetworkAuthenticationRequired,
			Msg:   msgs,
		}
	case http.StatusOK:
		response = ASTRICmodels.Response{
			Datos:  datos,
			Tag:    tag,
			Status: status,
			Msg:    msgs,
			Tkn:    global.GetTKN(),
		}
	case http.StatusBadRequest:
		response = ASTRICmodels.Response{
			Datos:  datos,
			Tag:    tag,
			Error:  errores,
			Msg:    msgs,
			Status: status,
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(stausHead)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic("Error al decodificar json para crear response " + err.Error())
	}

}
