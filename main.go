package main

import (
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"time"
)

func handlerServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal server error"))
}

func handlerLong(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	w.Write([]byte("Long answer"))
}

func handlerUserInfo(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	osName := r.Header.Get("Sec-Ch-Ua-Platform")
	w.Write([]byte(userAgent + "\n" + osName))
}

func handlerEcho(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func main() {
	r := chi.NewRouter()
	r.Get("/server_error/", handlerServerError)
	r.Post("/echo/", handlerEcho)
	r.Get("/long/", handlerLong)
	r.Get("/user/", handlerUserInfo)
	log.Fatal(http.ListenAndServe(":80", r))
}
