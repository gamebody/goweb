package main

import (
	"routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := routes.Init()

	r.Run(":8080")
}
