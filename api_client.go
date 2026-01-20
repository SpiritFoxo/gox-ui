package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetClientTrafficByEmail returns amount of downloaded and uploaded data by a client.
func (a *Api) GetClientTrafficByEmail(ctx context.Context, client *Client) (*GetClientTrafficResponse, error) {
	var resp GetClientTrafficResponse
	endpoint := fmt.Sprintf("/inbounds/getClientTraffics/%s", client.Email)
	client.Up = resp.Obj.Up
	client.Down = resp.Obj.Down
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

// GetClientTrafficByUUID returns amount of downloaded and uploaded data by a client.
func (a *Api) GetClientTrafficByUUID(ctx context.Context, client *Client) (*GetClientTrafficResponse, error) {
	var resp GetClientTrafficResponse
	endpoint := fmt.Sprintf("/inbounds/getClientTrafficsById/%s", client.UUID)
	client.Up = resp.Obj.Up
	client.Down = resp.Obj.Down
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

// AddClient allows to create new client inside selected inbound.
func (a *Api) AddClient(ctx context.Context, inbound *Inbound, client *Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := "/inbounds/addClient"
	settings := Settings{
		Clients: []JSONString[Client]{{
			Value: *client,
		}},
	}
	settingsJSON, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal settings: %w", err)
	}

	formData := map[string]string{
		"id":       strconv.FormatUint(uint64(inbound.ID), 10),
		"settings": string(settingsJSON),
	}

	return &resp, a.DoFormDataRequest(ctx, "POST", endpoint, formData, &resp)
}

// ResetClientTraffic resets selected client traffic.
func (a *Api) ResetClientTraffic(ctx context.Context, client *Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/%d/resetClientTraffic/%s", client.InboundId, client.Email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// UpdateClientInfo updates information about client.
func (a *Api) UpdateClientInfo(ctx context.Context, client *Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/updateClient/%s", client.UUID)
	settings := Settings{
		Clients: []JSONString[Client]{{
			Value: *client,
		}},
	}
	settingsJSON, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal settings: %w", err)
	}
	formData := map[string]string{
		"id":       strconv.FormatUint(uint64(client.InboundId), 10),
		"settings": string(settingsJSON),
	}

	return &resp, a.DoFormDataRequest(ctx, "POST", endpoint, formData, &resp)
}

// GetClientIpAdress returns clients`s IP adress.
func (a *Api) GetClientIpAdress(ctx context.Context, client *Client) (string, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/clientIps/%s", client.Email)
	return resp.Obj, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// ClearClientIps clears IP assigned to a client.
func (a *Api) ClearClientIps(ctx context.Context, client *Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/clearClientIps/%s", client.Email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// DeleteClient deletes client from inbound.
func (a *Api) DeleteClient(ctx context.Context, client *Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/%d/delClient/%s", client.InboundId, client.UUID)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// GetKey generates a key that can be used in suitable software.
func (a *Api) GetKey(ctx context.Context, client *Client) (string, error) {
	inbound, err := a.GetInbound(ctx, client.InboundId)
	if err != nil {
		return "", fmt.Errorf("inboud doesnt exist")
	}
	u, err := url.Parse(a.config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("error parsing domain info")
	}
	key := fmt.Sprintf("%s://%s@%s:%d?type=%s&security=%s&pbk=%s&fp=%s&sni=%s&sid=%s&spx=%s&flow=%s#%s-%s", inbound.Protocol, client.UUID, u.Host, inbound.Port, inbound.StreamSettings.Value.Network, inbound.StreamSettings.Value.Security, inbound.StreamSettings.Value.RealitySettings.Settings.PublicKey, inbound.StreamSettings.Value.RealitySettings.Settings.Fingerprint, inbound.StreamSettings.Value.RealitySettings.ServerNames[0], inbound.StreamSettings.Value.RealitySettings.ShortIds[0], inbound.StreamSettings.Value.RealitySettings.Settings.SpiderX, client.Flow, inbound.Remark, client.Email)
	return key, nil
}

// GetSubscriptionLink generates a link to client`s subscription info.
func (a *Api) GetSubscriptionLink(ctx context.Context, client *Client) (string, error) {
	if client.SubId == "" {
		return "", fmt.Errorf("client doesnt have valid subscription id")
	}
	u, err := url.Parse(a.config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("error parsing domain info")
	}
	link := fmt.Sprintf("%s:%d/sub/%s/%s", u.Scheme+"://"+u.Host, a.config.SubscriptionPort, a.config.SubscriptionURI, client.SubId)
	return link, nil
}
