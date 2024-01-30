package pmodels

// Response estructura de respuestas de la API
//
//swagger:model Response
type Response struct {
	//El Msg solo se enviara si hay mensajes de los desarrolladores y
	//si estan activados de las variables de entorno
	Msg map[string]string `json:"MSG,omitempty"`
	//Er parametro Error se enviara junto con los errores si no hay
	//errores este no se envia
	Error interface{} `json:"error,omitempty"`
	//Datos solicitados o enviados
	Datos interface{} `json:"datos,omitempty"`
	//Si es True la transaccion se ralizo si es False fallo
	Status bool `json:"status" validate:"required"`
	//Accion o transaccion que se realizo
	Tag string `json:"tag" validate:"required"`
	//Token renovado
	Tkn string `json:"tkn,omitempty" validate:"required"`
}

// ResUnautorize estructura interna de responce para desautorizaciones
//
//swagger:model ResUnautorize
type ResUnautorize struct {
	//Codigo sirve como identificacion para el front
	Code int `json:"code"`
	//Mensaje este se puede setear de las varibles de entorno
	Msg map[string]string `json:"MSG,omitempty"`
	//Er parametro Error se enviara junto con los errores si no hay
	//errores este no se envia
	Error interface{} `json:"error,omitempty"`
}

// ResponseWS response de websocket
//
//swagger:model ResponseWS
type ResponseWS struct {
	//App sirve como identificador de app
	App string
	//EndPoint sirve como identificador del endpoint
	EndPoint string
	//Tag accion o transaccion realizada
	Tag string
	//BD base de datos en la cual se produjo el cambio
	BD string
	//Tabla de la base de datos en la cual se produjo el cambio
	Tabla string
	//Usuario usuario que produjo el evento
	Usuario string
	//Data datos enviados si es nesesario
	Data interface{}
	//Msj mensaje si es nesesario
	Msj string
	//Error si es nesessario
	Error interface{}
}
