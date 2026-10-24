package main

import (
	"log"
	"net/http"
)

func main() {

	// ******method 1 to create a server****
	// server := &api{addr: ":9092"}
	// if err := http.ListenAndServe(server.addr, server); err != nil {
	// 	log.Fatal(err)
	// }

	// ******method 2 to create a server****
	// api := &api{addr: ":9092"}
	// server := &http.Server{
	// 	Addr:    ":9092",
	// 	Handler: api,
	// }

	// log.Println("Server is running on", server.Addr)
	// server.ListenAndServe()

	// ******method 3 to create a server******
	api := &api{addr: ":9092"}

	//router
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	log.Println("Server is running on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
