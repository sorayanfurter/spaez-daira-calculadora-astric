package res

import (
	"ASTRIC/BackEnd/ASTRIC/global"
	"ASTRIC/BackEnd/ASTRIC/helper/socket"
	"ASTRIC/BackEnd/shared/ws"
	"fmt"
	"net/http"
)

// Res modelo general de res
type Res struct {
	tag string
	w   http.ResponseWriter
	e   interface{}
	d   interface{}
	m   string
	ws  *socket.Responce
}

// SendNot tipo de SendNotificacion
type SendNot struct {
	SendNot func()
}

/*
NewResponse Crea un nuevo responce
Ej: res := NewResponse(tag string, w http.ResponseWriter)
Parametros:

	@Tag: Coresponde a la accion a realizar o solicitada por el request
	@w: La variable del response donde se ralizara la respuesta
*/
func NewResponse(tag string, w http.ResponseWriter) *Res {

	global.SetTagResponceWS(tag)
	ws := ws.NewResponceWS()
	return &Res{tag: tag, w: w, ws: ws}

}

// sendNot envia una notificacion a los request afectados
func (r *Res) sendNot() {
	r.ws.Send()
}

// GetWS retorna el websocket para su configuracion
func (r *Res) GetWS() *socket.Responce {
	return r.ws
}

// Msg configura un mensage en el response
func (r *Res) Msg(mensaje string) *Res {
	r.m = mensaje
	return r
}

// Err configura un error en el responce
func (r *Res) Err(err interface{}) *Res {
	r.e = err
	return r
}

// Dato configura el dato del responce
func (r *Res) Dato(dato interface{}) *Res {
	r.d = dato
	return r
}

// ErrSend configura y envia el error
func (r *Res) ErrSend(err interface{}) {
	r.e = err
	r.send()
}

// DatoSend configura y envia el dato
func (r *Res) DatoSend(dato interface{}) SendNot {
	r.d = dato
	r.send()
	ws := SendNot{SendNot: r.sendNot}
	return ws
}

// MsgSend configura un mensage en el response y lo envia
func (r *Res) MsgSend(mensaje string) {
	r.m = mensaje
	r.send()
}

// send envia toda la configuracion al request
func (r *Res) send() {
	if r.d != nil || r.e != nil || r.m != "" {
		ResponseMSG(r.tag, r.e, r.d, r.w, r.m)
	} else {
		colorRed := "\033[31m"
		colorReset := "\033[0m"
		fmt.Print(string(colorRed), "ATENCION ERROR EN RESPONCE: ")
		fmt.Println(string(colorReset), "Ningun parametro seteado con informacion a enviar")
	}

}
