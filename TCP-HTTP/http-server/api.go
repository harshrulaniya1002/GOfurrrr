package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (s *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting users
	w.Header().Set("Content-Type", "application/json")

	// encode users slice to json
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating users
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	u := User{
		ID:    payload.ID,
		Name:  payload.Name,
		Email: payload.Email,
	}

	if err := s.InsertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *api) InsertUser(u User) error {
	//input validation
	if u.ID == 0 || u.Name == "" || u.Email == "" {
		log.Println("Invalid user data")
		return errors.New("invalid user data")
	}

	//storage validation
	for _, user := range users {
		if user.ID == u.ID {
			log.Printf("User with ID %d already exists", u.ID)
			return errors.New("user already exists")
		}
	}

	users = append(users, u)
	log.Printf("User with ID %d inserted successfully", u.ID)
	return nil
}

// mux handles the routing of HTTP requests to the appropriate handler functions based on the request method and path.
// It acts as a multiplexer, directing incoming requests to the correct handler based on the defined routes.
// so not needed to implement ServeHTTP method for the api struct, as the mux will handle the routing of requests to the appropriate handler functions.
// ********func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) ******{
// 	// Implementation for serving HTTP requests
// 	switch r.Method {
// 		case http.MethodGet:
// 			switch r.URL.Path {
// 			case "/users":
// 				s.getUsersHandler(w, r)
// 			default:
// 				http.NotFound(w, r)
// 			}
// 		case http.MethodPost:
// 			switch r.URL.Path {
// 			case "/users":
// 				s.createUsersHandler(w, r)
// 			default:
// 				http.NotFound(w, r)
// 			}
// 		default:
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }
