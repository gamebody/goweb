package main

import (
	"github.com/gamebody/goweb/routes"
)

func main() {

	r := routes.Init()

	r.Run(":3001")
}
