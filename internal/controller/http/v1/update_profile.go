package v1

import (
	"encoding/json"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/dto"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/render"
	"net/http"
)

// UpdateProfile обновляет существующий или создаёт новый профиль
func (h *Handlers) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	input := dto.UpdateProfileInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "json decode error")

		return
	}

	err = h.usecase.UpdateProfile(input)
	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
