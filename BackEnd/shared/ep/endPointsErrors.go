package ep

import (
	"ASTRIC/BackEnd/ASTRIC/helper/errors"
)

/*
ErrorControlResponse Controla los errores en los controllers de puntos finales
@pila: Corresponde a donde se esta llamando el controlador de errores. Generalmente en el nombre de la funcion contenedora
@w: El ResponseWriter ( w http.ResponseWriter)
@r: El Request (r *http.Request)
*/
var ErrorControlResponse = errors.ErrorControlResponse
