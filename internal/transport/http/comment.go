package http

import (
	"context"
	"encoding/json"
	"go-rest-api/internal/comment"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CommentService interface {
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

}
