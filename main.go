package main

import (
	"APP-TOKO/db"
	"APP-TOKO/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":38600"))
}
