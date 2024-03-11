package main

import (
	"encoding/json"
	"net/http"
)

const (
	USERNAME = "admin"
	PASSWORD = "admin"
)

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		http.Error(w, "Someting Went Wrong / Header not Set", http.StatusBadRequest)
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)

	if !isValid {
		http.Error(w, "Invalid Username or Password", http.StatusUnauthorized)
		return false
	}

	return true
}

func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		http.Error(w, "Only GET is supported", http.StatusNotFound)
		return false
	}

	return true
}

func OutputJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)

	// json.NewEncoder(w).Encode(data)

}

func OutputJsonEncode(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// res, err := json.Marshal(data)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(res)

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
