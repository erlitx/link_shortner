package v1

import (
	"encoding/json"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/dto"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/render"
	"net/http"
)

func (h *Handlers) CreateProfile(w http.ResponseWriter, r *http.Request) {
	input := dto.CreateProfileInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "json decode error")

		return
	}

	output, err := h.usecase.CreateProfile(input)
	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")

		return
	}

	render.JSON(w, output, http.StatusOK)
}
