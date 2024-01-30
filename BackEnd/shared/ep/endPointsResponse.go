package ep

import (
	"ASTRIC/BackEnd/ASTRIC/helper/res"
)

// NewResponse Crea un nuevo responce
/*
Ej: res := NewResponse(tag string, w http.ResponseWriter)
Parametros:
	@Tag: Coresponde a la accion a realizar o solicitada por el request
	@w: La variable del response donde se ralizara la respuesta
*/
var NewResponse = res.NewResponse

/*
Response funcion que genera la respuesta de la API

	@tag: Es la tarea o funcion que se solicito ej: Nuevo Usuario
	@datos: Son los datos que se utilizaron para realizar el tag solisitado
	@w: Corresponde al ResponseWriter "w http.ResponseWriter"
*/
var Response = res.Response

/*
ResponseUnauthorized funcion que genera la respuesta para acciones no autorizadas

	@tag: Es la tarea o funcion que se solicito ej: Nuevo Usuario
	@datos: Son los datos que se utilizaron para realizar el tag solisitado
	@w: Corresponde al ResponseWriter "w http.ResponseWriter"
*/
var ResponseUnauthorized = res.ResponseUnauthorized

/*
ResponseMSG funcion que genera la respuesta de la API incorporando un mensaje especial
solo funcona con la variab de entrono NOTIFICCIONES en true

	@tag: Es la tarea o funcion que se solicito ej: Nuevo Usuario
	@datos: Son los datos que se utilizaron para realizar el tag solisitado
	@w: Corresponde al ResponseWriter "w http.ResponseWriter"
	@msg: Corresponde al mensaje el mismo aprecesra con 1 - !ATENCION¡ al comienzo
*/
var ResponseMSG = res.ResponseMSG
