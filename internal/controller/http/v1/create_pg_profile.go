package v1

import (
	"context"
	"fmt"
	"net/http"
)

func (h *Handlers) CreatePGProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ROUTER")
	ctx := context.Background()
	h.usecase.CreatePGProfile(ctx)
}
