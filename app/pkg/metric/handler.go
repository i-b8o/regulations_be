package metric

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
}

// TODO fix dependency on httprouter
func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, URL, h.Heartbeat)
}

func (h *Handler) Heartbeat(w http.ResponseWriter, req *http.Request) {
	log.Print("OK")
	w.WriteHeader(204)
}
