package main

import (
	"EniQilo/db"
	"EniQilo/routes"
)

func main() {
	dbPool := db.InitDB()

	r := routes.New(&routes.Routes{
		Db: dbPool,
	})

	r.Mount()
	//db.ClosePool(dbPool)
}
