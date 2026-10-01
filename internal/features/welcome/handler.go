package welcome

import (
	"encoding/json"
	"net/http"
	"time"
)

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func CreateNewHandler() *handler {
	return newHandler()
}

func (h *handler) GETWelcomeAPI() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"service":   "go-auth-microservice",
			"status":    "healthy",
			"uptime":    "operational",
			"timestamp": time.Now(),
		})
	})
}

func (h *handler) OPTIONSWelcome() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OPTIONS / ---> Up and Running"))
	})
}

func (h *handler) HEADWelcome() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HEAD / ---> Up and Running"))
	})
}
