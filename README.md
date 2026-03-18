# CLI-Application

Aplicação de linha de comando em Go para consulta DNS.

Ela permite, para um `host` informado, listar:

- os IPs associados ao nome (`ip`)
- os servidores DNS autoritativos que respondem por esse domínio (`server`)

## File Tree

```txt
.
├── README.md
├── go.mod
├── go.sum
├── main.go
└── app/
    └── app.go
```


## Pré-requisitos

- Go (conforme `go.mod`)
- Acesso à Internet/DNS para que as consultas `net.LookupIP` e `net.LookupNS` funcionem

## Como executar

Rodando diretamente via `go run`:

```sh
go run . ip --host google.com
go run . server --host google.com
```

Construindo um binário:

```sh
go build -o cli-application .
./cli-application ip --host google.com
./cli-application server --host google.com
```

## Comandos

### `ip`

Busca IPs de endereços na Internet (resolve `host` usando DNS).

- Flag: `--host` (string)
- Valor padrão: `google.com`

Exemplo:

```sh
go run . ip --host example.com
```

Saída: imprime cada IP encontrado em uma linha.

### `server`

Busca o nome do servidor na Internet (resolve servidores NS do `host`).

- Flag: `--host` (string)
- Valor padrão: `google.com`

Exemplo:

```sh
go run . server --host example.com
```

Saída: imprime `server.Host` de cada registro NS encontrado.

## Estrutura do projeto

- `main.go`: inicia a aplicação chamando `app.Generate()` e executa `app.Run(os.Args)`.
- `app/app.go`: define a CLI com `github.com/urfave/cli` e implementa as ações dos comandos `ip` e `server`.


## Dependências

- `github.com/urfave/cli` (framework para construir a CLI)
