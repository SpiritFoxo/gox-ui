package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Inbound represents inbound object with all its parameters.
type Inbound struct {
	ID             uint                       `json:"id"`
	Remark         string                     `json:"remark"`
	Listen         string                     `json:"listen"`
	Port           int                        `json:"port"`
	Protocol       string                     `json:"protocol"`
	Settings       JSONString[Settings]       `json:"settings"`
	StreamSettings JSONString[StreamSettings] `json:"streamSettings"`
	Enable         bool                       `json:"enable"`
	ExpiryTime     UnixTime                   `json:"expiryTime"`
	Total          int64                      `json:"total"`
	Up             int64                      `json:"up"`
	Down           int64                      `json:"down"`
}

type Settings struct {
	Clients []JSONString[Client] `json:"clients"`
}

// ClientSettings represents settings for a client in inbound.
type ClientSettings struct {
	ID         string   `json:"id"`
	Email      string   `json:"email"`
	LimitIP    int      `json:"limitIp"`
	Total      int      `json:"total"`
	ExpiryTime UnixTime `json:"expiryTime"`
}

type StreamSettings struct {
	Network  string `json:"network"`
	Security string `json:"security"`
}

type ListInboundsResponse struct {
	Obj     []Inbound `json:"obj"`
	Success bool      `json:"success"`
}

type GetInboundResponse struct {
	Obj     Inbound `json:"obj"`
	Success bool    `json:"success"`
}

type GetClientTrafficResponse struct {
	Obj     Client `json:"obj"`
	Success bool   `json:"success"`
}

type GetClientResponse struct {
	Obj     Client `json:"obj"`
	Success bool   `json:"success"`
}

type MessageResponse struct {
	Obj     string `json:"obj"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

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
	InboundId  *uint    `json:"inboundId"`
	Up         int64    `json:"up"`
	Down       int64    `json:"down"`
	AllTime    int64    `json:"allTime"`
}

type UnixTime struct {
	time.Time
}

type JSONString[T any] struct {
	Value T
}

func (ut *UnixTime) UnmarshalJSON(b []byte) error {
	var v int64
	if err := json.Unmarshal(b, &v); err != nil {
		return fmt.Errorf("UnixTime: %w", err)
	}
	if v == 0 {
		ut.Time = time.Time{}
		return nil
	}
	ut.Time = time.Unix(v, 0)
	return nil
}

func (ut UnixTime) MarshalJSON() ([]byte, error) {
	if ut.Time.IsZero() {
		return []byte("0"), nil
	}
	return []byte(strconv.FormatInt(ut.Time.Unix(), 10)), nil
}

func (j *JSONString[T]) UnmarshalJSON(b []byte) error {
	if len(b) > 0 && b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}

		return json.Unmarshal([]byte(s), &j.Value)
	}
	return json.Unmarshal(b, &j.Value)
}

func (j JSONString[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Value)
}
