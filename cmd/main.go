package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/server_error/", handlerServerError)
	mux.HandleFunc("/echo/", handlerEcho)
	mux.HandleFunc("/long/", handlerLong)
	mux.HandleFunc("/user/", handlerUserInfo)
	mux.HandleFunc("/", handlerHelloWorld)

	socket := os.Getenv("HTTP_SERVER_HOST_PORT")

	log.Printf("Listening on %s...", socket)
	log.Fatal(http.ListenAndServe(socket, mux))
}

func handlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

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
