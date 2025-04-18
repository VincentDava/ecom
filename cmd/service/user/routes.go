package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	//ini maksudnya RegisterRoutes nerima parameter tipedata mux.Router yang bernama router
	//(h *Handler) RegisterRoutes is a method that belongs to the struct Handler.
	//jadinya hanya bs lngsung panggil function kalau udh punya struct handler
	//misal jdnya hand.RegisterRoutes(muxrouter)
	router.HandleFunc("/login", h.handlelogin).Methods("POST")
	router.HandleFunc("/register", h.handleregister).Methods("POST")
}

func (h *Handler) handlelogin(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) handleregister(w http.ResponseWriter, r *http.Request) {

}
