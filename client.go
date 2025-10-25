package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// GetClientTrafficByEmail returns amount of downloaded and uploaded data by a client
func (a *Api) GetClientTrafficByEmail(ctx context.Context, email string) (*GetClientTrafficResponse, error) {
	var resp GetClientTrafficResponse
	endpoint := fmt.Sprintf("/inbounds/getClientTraffics/%s", email)
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

// GetClientTrafficByUUID returns amount of downloaded and uploaded data by a client
func (a *Api) GetClientTrafficByUUID(ctx context.Context, uuid string) (*GetClientTrafficResponse, error) {
	var resp GetClientTrafficResponse
	endpoint := fmt.Sprintf("/inbounds/getClientTrafficsById/%s", uuid)
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

// AddClient allows to create new client inside selected inbound
func (a *Api) AddClient(ctx context.Context, inboundId uint, client Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := "/inbounds/addClient"
	settings := Settings{
		Clients: []JSONString[Client]{{
			Value: client,
		}},
	}
	settingsJSON, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal settings: %w", err)
	}

	formData := map[string]string{
		"id":       strconv.FormatUint(uint64(inboundId), 10),
		"settings": string(settingsJSON),
	}

	return &resp, a.DoFormDataRequest(ctx, "POST", endpoint, formData, &resp)
}

// GetClientIpAdress returns clients`s IP adress
func (a *Api) GetClientIpAdress(ctx context.Context, email string) (*GetClientIpAdressResponse, error) {
	var resp GetClientIpAdressResponse
	endpoint := fmt.Sprintf("/inbounds/clientIps/%s", email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}
