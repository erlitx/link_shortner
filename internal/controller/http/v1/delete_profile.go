package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/erlitx/link_shortner/internal/dto"
	"github.com/erlitx/link_shortner/pkg/render"
	"net/http"
)

func (h *Handlers) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	input := dto.DeleteProfileInput{
		ID: chi.URLParam(r, "id"),
	}

	err := h.usecase.DeleteProfile(input)
	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
