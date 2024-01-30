package socket

import "github.com/olahol/melody"

// Socket estructura general de socket
type Socket struct {
	conn     *melody.Melody
	sessones []sessiones
}

type sessiones struct {
	databases []string
	session   *melody.Session
	usuario   string
}

type setSession struct {
	Databases []string
	Usuario   string
}
