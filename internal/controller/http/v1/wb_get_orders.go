package v1

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/dto"
)

func (h *Handlers) GetOrders(w http.ResponseWriter, r *http.Request) {
	// Get `dttm` and `flag` from query parameters
	dateFromStr := r.URL.Query().Get("dttm")
	flagStr := r.URL.Query().Get("flag")

	// Parse `dttm` into time.Time
	dateFrom, err := time.Parse(time.RFC3339, dateFromStr)
	if err != nil {
		http.Error(w, "Invalid dttm format", http.StatusBadRequest)
		return
	}

	// Parse `flag` (optional, default = 0)
	flag := 0
	if flagStr != "" {
		fmt.Sscanf(flagStr, "%d", &flag)
	}

	// Create input DTO
	input := dto.GetWBOrdersInput{
		DateFrom: dateFrom,
		Flag:     flag,
	}

	// Call usecase
	
	err = h.usecase.GetWBOrders(r.Context(), input)
	if err != nil {
		http.Error(w, "Error fetching orders", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
