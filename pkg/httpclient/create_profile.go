package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/httpclient/dto"
	"bytes"

)


func (c *Client) CreateProfile(name string, age int) (dto.Profile, error) {
	const createProfileEndpoint = "mbelogortsev/my-app/api/v1/profile"
	path := fmt.Sprintf("http://%s/%s", c.host, createProfileEndpoint)

	data := map[string]interface{}{
		"name": name,
		"age":  age,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return dto.Profile{}, fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
	if err != nil {
		return dto.Profile{}, fmt.Errorf("https req error")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return dto.Profile{}, fmt.Errorf("http req error: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return dto.Profile{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var profile dto.Profile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return dto.Profile{}, fmt.Errorf("error decoding response: %w", err)
	}

	return profile, nil

} 