package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var newLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	// ...

	w.Write([]byte("hello world"))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newLogger.Info("log", "host", r.URL.Host, "path", r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware2")

		next.ServeHTTP(w, r)
	})
}

func main() {

	r := mux.NewRouter()

	productHandler := http.HandlerFunc(ProductsHandler)

	r.Handle("/products", loggingMiddleware(middleware2(productHandler)))

	fmt.Println("Server is running at :8080")

	http.ListenAndServe(":8080", r)
}
