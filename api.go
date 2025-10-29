package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	defaultPort = 2053
	apiPrefix   = "/panel/api"
)

type Config struct {
	BaseURL    string
	Username   string
	Password   string
	Port       int
	HTTPClient *http.Client
	Timeout    time.Duration
}

type Api struct {
	config     Config
	httpClient *http.Client
	baseURL    string
}

func NewApi(cfg Config) (*Api, error) {
	if cfg.Port == 0 {
		cfg.Port = defaultPort
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = 30 * time.Second
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %w", err)
	}
	if cfg.HTTPClient == nil {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		transport := &http.Transport{
			TLSClientConfig: tlsConfig,
		}
		cfg.HTTPClient = &http.Client{
			Transport: transport,
			Timeout:   cfg.Timeout,
			Jar:       jar,
		}
	}

	u, err := url.Parse(cfg.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid baseURL: %w", err)
	}
	base := u.Scheme + "://" + u.Host + ":" + strconv.Itoa(cfg.Port) + u.Path

	c := &Api{
		config:     cfg,
		httpClient: cfg.HTTPClient,
		baseURL:    base,
	}

	if err := c.Login(context.Background()); err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	return c, nil
}

// Login performs login to the server and stores response cookie for further access.
func (a *Api) Login(ctx context.Context) error {
	loginURL := a.baseURL + "/login"

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("username", a.config.Username)
	_ = writer.WriteField("password", a.config.Password)

	_ = writer.WriteField("twoFactorCode", "")

	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, loginURL, &body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("login failed: %s", string(responseBody))
	}

	return nil
}

// ListInbounds gets array of all available inbounds from the server.
func (a *Api) ListInbounds(ctx context.Context) (*ListInboundsResponse, error) {
	var resp ListInboundsResponse
	endpoint := "inbounds/list"
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

// GetInbound gets inbound by provided id.
func (a *Api) GetInbound(ctx context.Context, inboundId uint) (*GetInboundResponse, error) {
	var resp GetInboundResponse
	endpoint := fmt.Sprintf("/inbounds/get/%d", inboundId)
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

func (a *Api) ResetAllTraffic(ctx context.Context) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := "/inbounds/resetAllTraffics"
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

func (a *Api) ResetTraffic(ctx context.Context, inboundId uint) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/resetAllClientTraffics/%d", inboundId)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// DeleteInbound deletes inbound from server
func (a *Api) DeleteInbound(ctx context.Context, inboundId uint) (*MessageResponse, error) {
	var resp MessageResponse
	enpoint := fmt.Sprintf("/inbounds/del/%d", inboundId)
	return &resp, a.DoRequest(ctx, "POST", enpoint, nil, &resp)
}

// DeleteDepletedClients delets all depleted clients from inbound
func (a *Api) DeleteDepletedClients(ctx context.Context, inboundId uint) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("inbounds/delDepletedClients/%d", inboundId)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

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

func (a *Api) ResetClientTraffic(ctx context.Context, inboundId uint, email string) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/%d/resetClientTraffic/%s", inboundId, email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

func (a *Api) UpdateClientinfo(ctx context.Context, inboundId uint, client Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/updateClient/%s", client.UUID)
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
func (a *Api) GetClientIpAdress(ctx context.Context, email string) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/clientIps/%s", email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// ClearClientIps clears IP assigned to a client
func (a *Api) ClearClientIps(ctx context.Context, email string) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/clearClientIps/%s", email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// DeleteClient deletes client from inbound
func (a *Api) DeleteClient(ctx context.Context, inboundId uint, uuid string) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/%d/delClient/%s", inboundId, uuid)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// DoRequest is a universal method for API requests.
func (a *Api) DoRequest(ctx context.Context, method, endpoint string, body, out interface{}) error {
	reqURL := a.baseURL + apiPrefix + endpoint
	var req *http.Request
	var err error

	fmt.Println(reqURL)

	switch method {
	case "GET":
		req, err = http.NewRequestWithContext(ctx, method, reqURL, nil)
	case "POST":
		b, errMarshal := json.Marshal(body)
		if errMarshal != nil {
			return errMarshal
		}
		req, err = http.NewRequestWithContext(ctx, method, reqURL, bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
	default:
		return fmt.Errorf("unsupported method: %s", method)
	}
	if err != nil {
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return &APIError{StatusCode: resp.StatusCode, Message: string(body)}
	}

	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}
	return nil
}

func (a *Api) DoFormDataRequest(ctx context.Context, method, endpoint string, formData map[string]string, out interface{}) error {
	reqURL := a.baseURL + apiPrefix + endpoint
	var req *http.Request
	var err error

	fmt.Println(reqURL)

	switch method {
	case "GET":
		req, err = http.NewRequestWithContext(ctx, method, reqURL, nil)
	case "POST":
		data := url.Values{}
		for key, value := range formData {
			data.Set(key, value)
		}
		req, err = http.NewRequestWithContext(ctx, method, reqURL, strings.NewReader(data.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	default:
		return fmt.Errorf("unsupported method: %s", method)
	}
	if err != nil {
		return err
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return &APIError{StatusCode: resp.StatusCode, Message: string(body)}
	}

	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}
	return nil
}
