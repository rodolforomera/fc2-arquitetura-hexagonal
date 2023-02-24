package main_old

import (
	"database/sql"
	db2 "github.com/rodolforomera/fc2-arquitetura-hexagonal/adapters/db"
	"github.com/rodolforomera/fc2-arquitetura-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)


func main() {

	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdpater := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdpater)
	product, _ := productService.Create("Product Exemplo", 30)

	productService.Enable(product)

}