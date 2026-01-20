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
	defaultPort    = 2053
	apiPrefix      = "/panel/api"
	defaultSubPort = 2096
)

type Config struct {
	BaseURL          string
	IsSecure         bool
	Username         string
	Password         string
	Port             int
	SubscriptionURI  string
	SubscriptionPort int
	HTTPClient       *http.Client
	Timeout          time.Duration
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
	if cfg.SubscriptionPort == 0 {
		cfg.SubscriptionPort = defaultSubPort
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
	cfg.IsSecure = false
	if u.Scheme == "https" {
		cfg.IsSecure = true
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

// SendBackupTelegram triggers the creation of a system backup and initiates the delivery of the backup file to designated administrators via a configured Telegram bot
func (a *Api) SendBackupTelegram(ctx context.Context) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := "/backuptotgbot"
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
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

// DoFormDataRequest is a universal method for API requests that require form data.
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
