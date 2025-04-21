package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/httpclient/dto"
)




func (c *Client) Get(id string) (dto.Profile, error) {
	const get_profile = "mbelogortsev/my-app/api/v1/profile"
	path := fmt.Sprintf("http://%s/%s/%s", c.host, get_profile, id)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return dto.Profile{}, fmt.Errorf("https req error")
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return dto.Profile{}, fmt.Errorf("http req error: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	switch resp.StatusCode {
	case http.StatusBadRequest:
	}

	var profile dto.Profile

	if err = json.Unmarshal(body, &profile); err != nil {
		return dto.Profile{}, fmt.Errorf("JSON %w", err)
	}

	return profile, nil

} 