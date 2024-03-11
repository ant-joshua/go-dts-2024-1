package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/students", ActionStudent)

	// base64.NewEncoder(base64.StdEncoding, w)

	sEnc := base64.StdEncoding.EncodeToString([]byte(USERNAME + ":" + PASSWORD))
	fmt.Println(sEnc)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("Server started at http://localhost:9000")
	server.ListenAndServe()
}
