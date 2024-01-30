package main

import (
	BackEnd "ASTRIC/BackEnd"
	migration "ASTRIC/BackEnd/migration"
	"flag"
	"fmt"
)

func main() {

	colorRed := "\033[31m"

	modo := flag.String("modo", "BeckEnd", "Mode de operacion del servidor")
	flag.Parse()

	fmt.Println(modo)

	switch *modo {
	case "BeckEnd":
		BackEnd.StartAPI("BeckEnd")
	case "BeckEndTKN":
		BackEnd.StartAPI("BeckEndTKN")
	case "Migratio":
		migration.StartMigration()
	default:
		fmt.Println(string(colorRed), "No se configuro el inicio o hay errores en los servidores")
	}

}
