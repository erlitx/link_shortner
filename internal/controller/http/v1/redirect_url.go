package v1

import (
	"context"
	"net/http"

	"github.com/erlitx/link_shortner/internal/dto"
	"github.com/erlitx/link_shortner/pkg/render"
	"github.com/go-chi/chi/v5"
)

func (h *Handlers) RedirectByShortURL(w http.ResponseWriter, r *http.Request) {
	short := chi.URLParam(r, "short_url")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	output, err := h.usecase.ResolveShortURL(ctx, dto.GetURLInput{ShortUrl: short})
	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")
		return
	}

	http.Redirect(w, r, output.RedirectURL, http.StatusFound)
}
