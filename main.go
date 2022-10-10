package main

import (
	"jwt-go/database"
	"jwt-go/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()

	r.Run(":8080")
}
