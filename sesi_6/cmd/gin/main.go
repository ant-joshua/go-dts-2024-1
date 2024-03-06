package main

import (
	"sesi_6/pkg/database"
	"sesi_6/pkg/routers"
)

func main() {
	port := ":5000"

	db := database.SQLInit()

	defer db.Close()

	start := routers.StartServer(db)
	start.Run(port)

}
