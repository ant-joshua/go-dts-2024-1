package main

import (
	"sesi_6/pkg/database"
	"sesi_6/pkg/routers"
)

func main() {
	port := ":5000"

	db := database.SQLInit()

	gorm := database.GormInit(db)

	defer db.Close()

	start := routers.StartServer(db, gorm)
	start.Run(port)

}
