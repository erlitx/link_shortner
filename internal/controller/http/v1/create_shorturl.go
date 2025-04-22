package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/erlitx/link_shortner/internal/dto"
	"github.com/erlitx/link_shortner/pkg/render"
)

const timeout = time.Second * 5

func (h *Handlers) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	input := dto.CreateShortUrlInput{Host: r.Host}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "json decode error")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	output, err := h.usecase.CreateShortURL(ctx, input)

	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")
		return
	}

	render.JSON(w, output, http.StatusOK)

}
