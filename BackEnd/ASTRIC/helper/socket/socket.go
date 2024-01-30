package socket

import (
	"encoding/json"
	"strings"

	"github.com/olahol/melody"
)

// WS websocket
var WS Socket

// NewSocket nueva instancia de conexion
func NewSocket(conn *melody.Melody) *Socket {
	WS.conn = conn
	WS.start()
	return &WS
}

// GetSocket retorna el Websocket
func GetSocket() *Socket {
	return &WS
}

func (s *Socket) start() {

	s.conn.HandleConnect(s.newSession)
	s.conn.HandleDisconnect(s.delSession)
	s.conn.HandleMessage(s.msg)

}

func (s *Socket) msg(Session *melody.Session, msg []byte) {
	s.setSession(Session, msg)
}

// Send envia el websocket a las secciones de la bd
func (s *Socket) Send(res *Responce) {

	data, _ := json.Marshal(res.res)
	for _, ss := range s.sessones {

		if existe(ss.databases, res.res.BD) {
			err := ss.session.Write(data)
			if err != nil {
				panic("Error al escribir en la session socket: " + err.Error())
			}
		}
	}
}

func existe(cont []string, find string) bool {
	for _, findCont := range cont {
		if strings.Contains(strings.ToLower(findCont), strings.ToLower(find)) {
			return true
		}
	}
	return false
}

func (s *Socket) setSession(Session *melody.Session, datos []byte) {

	var datosJSON setSession
	err := json.Unmarshal(datos, &datosJSON)
	if err != nil {
		msg := "Los datos enviados para setear la session son incorrectos: " + err.Error()
		err := Session.Write([]byte(msg))
		if err != nil {
			panic("Error al escribir en la session socket: " + err.Error())
		}
		return
	}

	if len(datosJSON.Databases) < 1 {
		msg := "Debe envar al menos una base de datos con el paramtro databases"
		err := Session.Write([]byte(msg))
		if err != nil {
			panic("Error al escribir en la session socket: " + err.Error())
		}
		return
	}

	if datosJSON.Usuario == "" {
		msg := "Debe envar el usauario en el parametro usuario"
		err := Session.Write([]byte(msg))
		if err != nil {
			panic("Error al escribir en la session socket: " + err.Error())
		}
		return
	}

	for index, ss := range s.sessones {
		if ss.session == Session {
			s.sessones[index].databases = datosJSON.Databases
			s.sessones[index].usuario = datosJSON.Usuario
		}
	}
}

func (s *Socket) newSession(Session *melody.Session) {
	ss := sessiones{
		session: Session,
	}
	s.sessones = append(s.sessones, ss)
}

func (s *Socket) delSession(Session *melody.Session) {
	for index, ss := range s.sessones {
		if ss.session == Session {
			s.sessones = append(s.sessones[:index], s.sessones[index+1:]...)
		}
	}
}
