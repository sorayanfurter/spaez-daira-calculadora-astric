package routes

import (
	"container/list"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Rout estructura de la clase
type Rout struct {
	request  *http.Request
	Rutafull string
	Ruta     string
	App      string
	Modulo   string
	Bd       string
}

/*
NewRoutDesc Constructor de la clase Routs
@req: Recibe el request propiamente dicho
Return:

	Retorna un puntero a la clase rout
*/
func NewRoutDesc(req *http.Request) *Rout {

	return &Rout{
		request:  req,
		Rutafull: req.RequestURI,
		Ruta:     getRuta(req),
		App:      getApp(req.RequestURI),
		Modulo:   getModulo(getRuta(req), getApp(req.RequestURI)),
	}

}

func getRuta(r *http.Request) string {
	paramtetros := mux.Vars(r)
	path := strings.Split(r.RequestURI, "/")
	lista := list.New()
	var rutasFull string
	for _, v := range path {
		lista.PushBack(v)
	}
	for v := lista.Front(); v != nil; v = v.Next() {
		for _, p := range paramtetros {
			if b := v.Prev(); b != nil && v.Value == p {
				lista.Remove(v)
				v = b
			} else if v.Value == p {
				b := v.Next()
				lista.Remove(v)
				v = b
			} else if v.Value == "" {
				b := v.Next()
				lista.Remove(v)
				v = b
			}
		}
	}
	for v := lista.Front(); v != nil; v = v.Next() {
		if v.Value != "" {
			rutasFull = rutasFull + fmt.Sprintf("/%s", v.Value)
		}
	}

	ruta := strings.Split(rutasFull, "-")
	return ruta[0]
}

func getApp(r string) string {
	path := strings.Split(r, "/")
	return path[1]
}

func getModulo(r string, app string) string {
	path := strings.Split(r, app+"/")
	return path[1]
}
