package main

import (
	"svc-myfood-echo/db"
	"svc-myfood-echo/routes"
)

func main(){
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":9080"))
}
