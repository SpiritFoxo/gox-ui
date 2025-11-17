package api

import (
	"encoding/json"
	"strconv"
	"time"
)

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

type UnixTime struct {
	time.Time
}

type JSONString[T any] struct {
	Value T
}

func (ut *UnixTime) UnmarshalJSON(b []byte) error {
	var ms float64
	if err := json.Unmarshal(b, &ms); err == nil {
		if ms == 0 {
			ut.Time = time.Time{}
			return nil
		}
	}
	secs := int64(ms) / 1000
	nsecs := int64(ms*1_000_000) % 1_000_000_000
	ut.Time = time.Unix(secs, nsecs).UTC()
	return nil
}

func (ut UnixTime) MarshalJSON() ([]byte, error) {
	if ut.Time.IsZero() {
		return []byte("0"), nil
	}
	ms := ut.Time.Unix()*1000 + int64(ut.Time.Nanosecond())/1_000_000
	return []byte(strconv.FormatInt(ms, 10)), nil
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
