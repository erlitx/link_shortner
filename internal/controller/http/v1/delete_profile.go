package v1

import (
	"github.com/go-chi/chi/v5"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/dto"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/render"
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
