package wbadapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/dto"
)

func (wb *WBConn) GetWBOrders(ctx context.Context, input dto.GetWBOrdersInput) ([]dto.GetWBordersOutput, error) {
	base, err := url.Parse(wb.API.Host)
	base.Path = path.Join(base.Path, wb.API.Orders)

	// Format the date correctly
	dateString := input.DateFrom.Format("2006-01-02T00:00:00")

	// Add query parameters
	params := url.Values{}
	params.Add("dateFrom", dateString)
	params.Add("flag", strconv.Itoa(input.Flag))

	// Attach query params to URL
	base.RawQuery = params.Encode()

	finalURL := base.String() // Get the final formatted URL

	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		fmt.Errorf("Error creating request:", err)
		return []dto.GetWBordersOutput{}, err
	}
	req.Header.Set("Authorization", "Bearer "+wb.Client.WBTokenStats)
	resp, err := wb.Client.Client.Do(req)
	if err != nil {
		fmt.Errorf("Request failed:", err)
		return []dto.GetWBordersOutput{}, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error reading response:, %w", err)
		return []dto.GetWBordersOutput{}, err
	}

	var orders []dto.GetWBordersOutput

	err = json.Unmarshal(body, &orders)
	if err != nil {
		fmt.Errorf("Error decoding JSON:", err)
		return []dto.GetWBordersOutput{}, err
	}

	return orders, nil

}
