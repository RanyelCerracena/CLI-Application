package app

import "github.com/urfave/cli"

// Generate vai retornar a aplicação de linha de comando pronta para ser executada
func Generate() *cli.App{
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet"

	return app
}