package main

import (
	"log"
	"os"

	"github.com/RanyelCerracena/CLI-Application/app"
)

// main e main.go apenas executa a aplicação. A lógica ficará em outro arquivo e será importado para cá.
func main() {
	application := app.Generate()    // Importa a função Generate e trás sua lógica
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
