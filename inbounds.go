package api

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

type StreamSettings struct {
	Network  string `json:"network"`
	Security string `json:"security"`
}

// GetClientByEmail gets most of the client info
func (i *Inbound) GetClientByEmail(api *Api, email string) *Client {
	for _, client := range i.Settings.Value.Clients {
		if client.Value.Email == email {
			return &Client{
				Comment:    client.Value.Comment,
				Email:      client.Value.Email,
				Enable:     client.Value.Enable,
				ExpiryTime: client.Value.ExpiryTime,
				Flow:       client.Value.Flow,
				UUID:       client.Value.UUID,
				LimitIP:    client.Value.LimitIP,
				Reset:      client.Value.Reset,
				SubId:      client.Value.SubId,
				TgId:       client.Value.TgId,
				TotalGB:    client.Value.TotalGB,
				InboundId:  client.Value.InboundId,
				Up:         client.Value.Up,
				Down:       client.Value.Down,
				AllTime:    client.Value.AllTime,
			}
		}
	}
	return nil
}
