package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"todolist_server/users"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	userStorage := users.NewInMemoryUserStorage()
	userService := users.NewUserService(userStorage)
	jwtService, err := users.NewJWTService("pubkey.rsa", "privkey.rsa")
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/user/signup", userService.Register).Methods(http.MethodPost)
	r.HandleFunc("/user/signin", users.WrapJwt(jwtService, userService.JWT)).Methods(http.MethodPost)

	r.HandleFunc("/todo/lists", jwtService.JWTAuth(userStorage, users.AddList)).Methods(http.MethodPost)
	r.HandleFunc("/todo/lists/{list_id:[0-9]+}", jwtService.JWTAuth(userStorage, users.UpdateList)).Methods(http.MethodPut)
	r.HandleFunc("/todo/lists/{list_id:[0-9]+}", jwtService.JWTAuth(userStorage, users.DeleteList)).Methods(http.MethodDelete)
	r.HandleFunc("/todo/lists", jwtService.JWTAuth(userStorage, users.GetLists)).Methods(http.MethodGet)

	r.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks", jwtService.JWTAuth(userStorage, users.AddTask)).Methods(http.MethodPost)
	r.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks/{task_id:[0-9]+}", jwtService.JWTAuth(userStorage, users.UpdateTask)).Methods(http.MethodPut)
	r.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks/{task_id:[0-9]+}", jwtService.JWTAuth(userStorage, users.DeleteTask)).Methods(http.MethodDelete)
	r.HandleFunc("/todo/lists/{list_id:[0-9]+}/tasks", jwtService.JWTAuth(userStorage, users.GetTasks)).Methods(http.MethodGet)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()

	log.Println("Server started")
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Error: ", err)
	}
}
