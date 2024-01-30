package socket

import (
	"ASTRIC/BackEnd/ASTRIC/global"
	ASTRICmodels "ASTRIC/BackEnd/ASTRIC/pModels"
)

// Responce modelo de responce wensoket
type Responce struct {
	res ASTRICmodels.ResponseWS
}

// NewResponce crea un nuevo responce webssocket
func NewResponce() *Responce {

	return &Responce{res: global.GetResponceWS()}
}

// SetMsj setea un mensaje para enviar al websocket
func (r *Responce) SetMsj(mjs string) *Responce {
	r.res.Msj = mjs
	return r
}

// SetError setea un error para enviar al websocket
func (r *Responce) SetError(err interface{}) *Responce {
	r.res.Error = err
	return r
}

// SetDatos seta datos para enviar el websocket
func (r *Responce) SetDatos(datos interface{}) *Responce {
	r.res.Data = datos
	return r
}

// Send enviar socket
func (r *Responce) Send() {
	WS := GetSocket()
	WS.Send(r)
}
