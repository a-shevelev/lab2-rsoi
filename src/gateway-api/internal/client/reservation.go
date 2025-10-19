package client

import (
	"encoding/json"
	"fmt"
	"gateway-api/internal/dto"
	"net/http"
)

type Reservation struct {
	BaseURL    string `envconfig:"BASE_URL"`
	HTTPClient *http.Client
}

// NewReservationClient создаёт новый клиент для ReservationService
func NewReservation(baseURL string) *Reservation {
	return &Reservation{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
	}
}

func (c *Reservation) Get(username string) ([]dto.ReservationResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/reservation", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-User-Name", username)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
	fmt.Println(resp.Body)

	var result []dto.ReservationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
