package main

import (
	"net/http"
)

func main() {
	api := &api{
		address: ":8080",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	srv := &http.Server{
		Addr:    api.address,
		Handler: mux,
	}
	srv.ListenAndServe()
	/*
		 	if err := http.ListenAndServe(s.address, s); err != nil {
				log.Fatal(err)
			}
	*/
}
