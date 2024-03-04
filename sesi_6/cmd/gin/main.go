package main

import "sesi_6/pkg/routers"

func main() {
	port := ":5000"

	routers.StartServer().Run(port)
}
