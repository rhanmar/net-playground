package http

import (
	"context"
	"io"
	"net/http"
	"time"
)

type dummyService interface {
	Save(ctx context.Context, data string) error
}

type Handler struct {
	service dummyService
}

func NewHandler(svc dummyService) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) SaveUser(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	osName := r.Header.Get("Sec-Ch-Ua-Platform")
	data := userAgent + "\t|\t" + osName

	err := h.service.Save(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(data))
}

func (h *Handler) HandlerEcho(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) HandlerLong(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	w.Write([]byte("Long answer"))
}

func (h *Handler) HandlerServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal server error"))
}

func (h *Handler) HandlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
