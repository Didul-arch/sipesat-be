// Package subjects for subject logics
package subjects

import (
	"log"
	"net/http"

	"github.com/Didul-arch/sipesat/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListSubjects(w http.ResponseWriter, r *http.Request) {
	subjects, err := h.service.ListSubjects(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, subjects)
}
