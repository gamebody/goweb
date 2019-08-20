package main

import (
	"github.com/gamebody/goweb/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := routes.Init()

	r.Run(":8080")
}
