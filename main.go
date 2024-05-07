package main

import (
	"EniQilo/db"
	"EniQilo/routes"
)

func main() {
	dbPool := db.InitDB()
	routes.New(dbPool)
	db.ClosePool(dbPool)
}
