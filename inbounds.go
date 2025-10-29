package api

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
