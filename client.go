package api

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"
)

// Client represents a user client.
type Client struct {
	Comment    string   `json:"comment"`
	Email      string   `json:"email"`
	Enable     bool     `json:"enable"`
	ExpiryTime UnixTime `json:"expiryTime"`
	Flow       string   `json:"flow"`
	UUID       string   `json:"id"`
	LimitIP    int      `json:"limitIp"`
	Reset      int16    `json:"reset"`
	SubId      string   `json:"subId"`
	TgId       string   `json:"tgId"`
	TotalGB    int64    `json:"totalGB"`
	InboundId  uint     `json:"inboundId"`
	Up         int64    `json:"up"`
	Down       int64    `json:"down"`
	AllTime    int64    `json:"allTime"`
}

// ClientSettings represents settings for a client in inbound.
type ClientSettings struct {
	ID         string   `json:"id"`
	Email      string   `json:"email"`
	LimitIP    int      `json:"limitIp"`
	Total      int      `json:"total"`
	ExpiryTime UnixTime `json:"expiryTime"`
}

func NewClient() (*Client, error) {
	uuidBytes := make([]byte, 16)
	if _, err := rand.Read(uuidBytes); err != nil {
		return nil, errors.New("failed to generate UUID")
	}

	return &Client{
		UUID:       hex.EncodeToString(uuidBytes),
		Enable:     true,
		ExpiryTime: UnixTime{time.Unix(0, 0).UTC()},
	}, nil
}
