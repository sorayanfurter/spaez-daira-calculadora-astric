package lib

import (
	"container/list"
	"fmt"
)

// ArrayCompDeleteStr Toma dos arrays o slices los compara y borra del target los datos que estan en el source
/*
@target: datos a procesar
@source: fuente da datos a borrar del target
Return:
	Arrays
*/
func ArrayCompDeleteStr(target []string, sourse []string) []string {

	var resoult []string
	lista := list.New()

	for _, v := range target {
		lista.PushBack(v)
	}

	for v := lista.Front(); v != nil; v = v.Next() {
		for _, p := range sourse {
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

	var cont int

	for v := lista.Front(); v != nil; v = v.Next() {
		resoult[cont] = fmt.Sprintf("%s", v.Value)
		cont++
	}

	return resoult

}
