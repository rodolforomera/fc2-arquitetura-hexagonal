# Arquitetura Hexagonal

Exemplo simples de arquitetura hexagonal utilizando *Go, SQLite3 e Cobra (CLI)* conforme o curso **Full Cycle**

Dois adapters para executar os comandos (Utilizando comandos via cli ou iniciando um web server (http))

## Tests

```
mockgen -destination=application/mocks/application.go -source=application/product.go application
go test ./...
```

## SQLite

```
sqlite
touch sqlite.db
sqlite3 sqlite.db
create table products(id string, name string, price float, status string);
```

## Cobra

```
cobra-cli init
cobra-cli add cli
cobra-cli add http

go run main.go cli -a=create -n=Product CLI -p=25.0 
go run main.go cli -a=get --id=58b88dd0-a83d-48f4-8ae1-4aa353f5017c
go run main.go http
````
