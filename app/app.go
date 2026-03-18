package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Generate vai retornar a aplicação de linha de comando pronta para ser executada
func Generate() *cli.App {
	app := cli.NewApp()                                     // Cria uma nova aplicação CLI
	app.Name = "Aplicação de Linha de Comando"              // Atribui nome a aplicação
	app.Usage = "Busca IPs e Nomes de servidor na internet" // Atribui uma descrição básica do que a aplicação faz

	flags := []cli.Flag{ // Atribui flags como parametros que é do tipo array/slice de cli.Flag
		cli.StringFlag{ // Define o primeiro parametro como string
			Name:  "host",       // Nome do parametro
			Value: "google.com", // Valor default desse parametro caso omitido pelo usuário
		},
	}

	app.Commands = []cli.Command{ // Especifica os comandos que o sistema vai possuir
		{
			Name:   "ip",                                 // Define o Nome do primeiro comando
			Usage:  "Busca IPS de endereços na Internet", // Explica basicamente o que faz esse comando
			Flags:  flags,
			Action: searchIps, // Recebe uma função declarada normalmente ou de forma externa como nesse caso. Essa função define o que o comando configurado vai fazer.
		},
		{
			Name:   "server",
			Usage:  "Busca o nome do servidor na internet",
			Flags:  flags,
			Action: searchServers,
		},
	}
	return app // Retorna a aplicação construida para a função Generate
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func searchIps(c *cli.Context) { // Função que recebe um ponteiro de cli.Context, e executa o código específicado
	host := c.String("host")        // Define a variavel host, que recebe de c.String o valor do host que é definido na linha 24
	ips, erro := net.LookupIP(host) // Cria uma variavel ips que é do tipo array de ips, que é importado pelo pacote net com a função Lookup() recebendo o host como parametro

	if erro != nil { // Valida se o erro existe, se sim, cria um log fatal
		log.Fatal(erro)
	}
	for _, ip := range ips { // Percorre o array de ips e imprime na tela
		fmt.Println(ip)
	}
}
