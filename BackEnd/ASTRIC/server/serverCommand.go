package server

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

// ClearConsole borra la consola VALIDO PARA WINDOWS Y LINUX
func ClearConsole() {

	if !ifbool("PRODUCCION") && !ifbool("DISABLE_CONSOLA") {
		colorRed := "\033[31m"
		value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
		if ok {
			value()
		} else {
			fmt.Println(string(colorRed), "No se puede borrar la consola, sistema operativo no compatible")
		}
	}
}

// VerificarSO Verifica el sistem operativo VALIDO PARA WINDOWS Y LINUX
func VerificarSO() error {
	so := runtime.GOOS
	if so == "linux" || so == "windows" {
		return nil
	}
	return errors.New("los sistemas operativos soportados son Linux o Windows")
}

func ifbool(target string) bool {
	retult, _ := strconv.ParseBool(os.Getenv(target))
	return retult
}
