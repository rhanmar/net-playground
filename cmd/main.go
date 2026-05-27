package main

import (
	"context"
	"log"
	"net/http"

	"net-playground/internal/config"
	dummyDB "net-playground/internal/db"
	handlers "net-playground/internal/handlers/http"
	dummyRepo "net-playground/internal/repositories/dummy"
	dummySvc "net-playground/internal/services/dummy"
)

func main() {
	ctx := context.Background()
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := dummyDB.NewDB(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := dummyRepo.NewRepository(db)
	service := dummySvc.NewService(repo)
	httpHandlers := handlers.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/server_error/", httpHandlers.HandlerServerError)
	mux.HandleFunc("/echo/", httpHandlers.HandlerEcho)
	mux.HandleFunc("/long/", httpHandlers.HandlerLong)
	mux.HandleFunc("/user/", httpHandlers.SaveUser)
	mux.HandleFunc("/users/", httpHandlers.GetDummyInfo)
	mux.HandleFunc("/", httpHandlers.HandlerHelloWorld)

	log.Printf("Listening on %s...", cfg.GetHTTPServerHostPort())
	log.Fatal(http.ListenAndServe(cfg.GetHTTPServerHostPort(), mux))
}
