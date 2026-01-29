# Api

  

This module provides the core Api struct to interact with the 3X-UI API.

  

<a  id="api.api.Api"></a>

## Api Objects

  

```go

type  Api  struct {

config  Config

httpClient *http.Client

}

```

  

This struct provides a high-level interface to interact with the 3X-UI API.

Access to inbound and client methods is provided through this struct.

  

**Arguments** (via Config):

-  `BaseURL`  _string_ - The base URL of the 3X-UI panel.

-  `Username`  _string_ - The panel username.

-  `Password`  _string_ - The panel password.

-  `SubscriptionURI`  _string_ - The subscription path (optional).

-  `SubscriptionPort`  _int_ - The subscription port (default: 2096).

-  `HTTPClient`  _*http.Client_ - Custom HTTP client (optional).

-  `Timeout`  _time.Duration_ - Request timeout (default: 30s).

  

Public Methods:

- NewApi: Creates a new Api instance.

- Login: Performs login and stores the session cookie.

- SendBackupTelegram: Triggers a system backup via Telegram.


  

<a  id="api.api.NewApi"></a>

#### NewApi

  

```go

func  NewApi(cfg  Config) (*Api, error)

```

  

This function initializes and returns a new Api instance based on the provided configuration.

  

**Arguments**:

-  `cfg`  _Config_ - The configuration for the API.

  

**Returns**:

-  `*Api` - A pointer to the new Api instance.

-  `error` - An error if initialization fails (e.g., invalid URL).

  

**Examples**:

```go

cfg := api.Config{

BaseURL: "https://your-host.com:2053/your-panel-path",

Username: "admin",

Password: "admin",

}

apiClient, err := api.NewApi(cfg)

if  err != nil {

log.Fatal(err)

}

ctx := context.Background()

err = apiClient.Login(ctx)

if  err != nil {

log.Fatal(err)

}


```

  

<a  id="api.api.Api.Login"></a>

#### Login

  

```go

func (a *Api) Login(ctx  context.Context, twoFactorCode ...string) error

```

  

This method performs login to the server and stores the response cookie for further access.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `twoFactorCode`  _...string_ - Optional 2FA code.

  

**Returns**:

-  `error` - An error if login fails.

  

**Examples**:

```go
ctx := context.Background()

//err = apiClient.Login(ctx) // Without 2FA 

err = apiClient.Login(ctx, "123456") // With 2FA

if  err != nil {

log.Fatal(err)

}

```

  

<a  id="api.api.Api.SendBackupTelegram"></a>

#### SendBackupTelegram

  

```go

func (a *Api) SendBackupTelegram(ctx  context.Context) (*MessageResponse, error)

```

  

This method triggers the creation of a system backup and initiates delivery via Telegram bot.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

  

**Returns**:

-  `*MessageResponse` - A pointer to the response message.

-  `error` - An error if the request fails.

  

**Examples**:

```go

resp, err := apiClient.SendBackupTelegram(ctx)

if  err != nil {

log.Fatal(err)

}

```

  

<a  id="api.api.Api.DoRequest"></a>

  


# Inbounds

  

This module contains methods related to inbound management.

  

<a  id="api.api_inbound.Api"></a>

## Api Methods for Inbounds

  

```go

type  Inbound  struct {

ID  uint  `json:"id"`

Remark  string  `json:"remark"`

Listen  string  `json:"listen"`

Port  int  `json:"port"`

Protocol  string  `json:"protocol"`

Settings  JSONString[Settings] `json:"settings"`

StreamSettings  JSONString[StreamSettings] `json:"streamSettings"`

Enable  bool  `json:"enable"`

ExpiryTime  UnixTime  `json:"expiryTime"`

Total  int64  `json:"total"`

Up  int64  `json:"up"`

Down  int64  `json:"down"`

}

  

type  Settings  struct {

Clients []JSONString[Client] `json:"clients"`

}

  

type  StreamSettings  struct {

Network  string  `json:"network"`

Security  string  `json:"security"`

ExternalProxy []interface{} `json:"externalProxy"`

RealitySettings  RealitySettings  `json:"realitySettings"`

TCPSettings  TCPSettings  `json:"tcpSettings"`

}

  

type  RealitySettings  struct {

Show  bool  `json:"show"`

Xver  int  `json:"xver"`

Dest  string  `json:"dest"`

ServerNames []string  `json:"serverNames"`

PrivateKey  string  `json:"privateKey"`

MinClient  string  `json:"minClient"`

MaxClient  string  `json:"maxClient"`

MaxTimediff  int  `json:"maxTimediff"`

ShortIds []string  `json:"shortIds"`

Settings  RealityInnerSettings  `json:"settings"`

}

  

type  RealityInnerSettings  struct {

PublicKey  string  `json:"publicKey"`

Fingerprint  string  `json:"fingerprint"`

ServerName  string  `json:"serverName"`

SpiderX  string  `json:"spiderX"`

}

  

type  TCPSettings  struct {

AcceptProxyProtocol  bool  `json:"acceptProxyProtocol"`

Header  struct {

Type  string  `json:"type"`

} `json:"header"`

}

```

  

This struct provides methods to interact with inbounds in the 3X-UI API.

  

Public Methods for Inbounds:

- ListInbounds: Retrieves a list of all inbounds.

- GetInbound: Retrieves a specific inbound by ID.

- ResetAllTraffic: Resets traffic statistics for all inbounds.

- ResetTraffic: Resets traffic statistics for a specific inbound.

- DeleteInbound: Deletes a specific inbound.

- DeleteDepletedClients: Deletes depleted clients from a specific inbound.

- GetClientByEmail: Retrieves a client by email from the inbound.



  

<a  id="api.api_inbound.Api.ListInbounds"></a>

#### ListInbounds

  

```go

func (a *Api) ListInbounds(ctx  context.Context) (*[]Inbound, error)

```

  

This method retrieves an array of all available inbounds from the server.

  

**Returns**:

-  `*[]Inbound` - A pointer to a slice of Inbound structs.

-  `error` - An error if the request fails.

  

**Examples**:

```go

inbounds, err := apiClient.ListInbounds(ctx)

if  err != nil {

log.Fatal(err)

}

for  _, inbound := range *inbounds {

log.Printf("Inbound ID: %d, Remark: %s", inbound.ID, inbound.Remark)

}

```

  

<a  id="api.api_inbound.Api.GetInbound"></a>

#### GetInbound

  

```go

func (a *Api) GetInbound(ctx  context.Context, inboundId  uint) (*Inbound, error)

```

  

This method retrieves an inbound by the provided ID.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `inboundId`  _uint_ - The ID of the inbound to retrieve.

  

**Returns**:

-  `*Inbound` - A pointer to the Inbound struct.

-  `error` - An error if the inbound is not found or the request fails.

  

**Examples**:

```go

inbound, err := apiClient.GetInbound(ctx, 1)

if  err != nil {

log.Fatal(err)

}

log.Printf("Inbound Remark: %s", inbound.Remark)

```

  

<a  id="api.api_inbound.Api.ResetAllTraffic"></a>

#### ResetAllTraffic

  

```go

func (a *Api) ResetAllTraffic(ctx  context.Context) (*MessageResponse, error)

```

  

This method resets the traffic statistics for all inbounds within the server.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

  

**Returns**:

-  `*MessageResponse` - A pointer to the response message.

-  `error` - An error if the request fails.

  

**Examples**:

```go

resp, err := apiClient.ResetAllTraffic(ctx)

if  err != nil {

log.Fatal(err)

}

log.Println(resp.Msg)

```

  

<a  id="api.api_inbound.Api.ResetTraffic"></a>

#### ResetTraffic

  

```go

func (a *Api) ResetTraffic(ctx  context.Context, inbound *Inbound) (*MessageResponse, error)

```

  

This method resets the traffic statistics for all clients associated with a specific inbound.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `inbound`  _*Inbound_ - The inbound to reset traffic for.

  

**Returns**:

-  `*MessageResponse` - A pointer to the response message.

-  `error` - An error if the request fails.

  

**Examples**:

```go

inbound, _ := apiClient.GetInbound(ctx, 1)

resp, err := apiClient.ResetTraffic(ctx, inbound)

if  err != nil {

log.Fatal(err)

}

```

  

<a  id="api.api_inbound.Api.DeleteInbound"></a>

#### DeleteInbound

  

```go

func (a *Api) DeleteInbound(ctx  context.Context, inbound *Inbound) (*MessageResponse, error)

```

  

This method deletes an inbound from the server.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `inbound`  _*Inbound_ - The inbound to delete.

  

**Returns**:

-  `*MessageResponse` - A pointer to the response message.

-  `error` - An error if the request fails.

  

**Examples**:

```go

inbound, _ := apiClient.GetInbound(ctx, 1)

resp, err := apiClient.DeleteInbound(ctx, inbound)

if  err != nil {

log.Fatal(err)

}

```

  

<a  id="api.api_inbound.Api.DeleteDepletedClients"></a>

#### DeleteDepletedClients

  

```go

func (a *Api) DeleteDepletedClients(ctx  context.Context, inbound *Inbound) (*MessageResponse, error)

```

  

This method deletes all depleted clients from an inbound.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `inbound`  _*Inbound_ - The inbound to delete depleted clients from.

  

**Returns**:

-  `*MessageResponse` - A pointer to the response message.

-  `error` - An error if the request fails.

  

**Examples**:

```go

inbound, _ := apiClient.GetInbound(ctx, 1)

resp, err := apiClient.DeleteDepletedClients(ctx, inbound)

if  err != nil {

log.Fatal(err)

}

```


  

<a  id="api.inbound.Inbound.GetClientByEmail"></a>

#### GetClientByEmail

  

```go

func (i *Inbound) GetClientByEmail(api *Api, email  string) *Client

```

  

This method retrieves most of the client info by email.

  

**Arguments**:

-  `api`  _*Api_ - The API client (unused in method, but for compatibility).

-  `email`  _string_ - The email to search for.

  

**Returns**:

-  `*Client` - A pointer to the Client if found, else nil.

**Examples**:

```go

client := inbound.GetClientByEmail(apiClient, "your-email-string")

if client != nil {

	fmt.Println(client.UUID)
	
}
```

  
  

# Clients

  

This module contains methods related to client management.

  

<a  id="api.api_client.Api"></a>

## Api Methods for Clients

  

```go

type  Client  struct {

Comment  string  `json:"comment"`

Email  string  `json:"email"`

Enable  bool  `json:"enable"`

ExpiryTime  UnixTime  `json:"expiryTime"`

Flow  string  `json:"flow"`

UUID  string  `json:"id"`

LimitIP  int  `json:"limitIp"`

Reset  int16  `json:"reset"`

SubId  string  `json:"subId"`

TgId  string  `json:"tgId"`

TotalGB  int64  `json:"totalGB"`

InboundId  uint  `json:"inboundId"`

Up  int64  `json:"up"`

Down  int64  `json:"down"`

AllTime  int64  `json:"allTime"`

}


type  ClientSettings  struct {

ID  string  `json:"id"`

Email  string  `json:"email"`

LimitIP  int  `json:"limitIp"`

Total  int  `json:"total"`

ExpiryTime  UnixTime  `json:"expiryTime"`

}

```

  

This struct provides methods to interact with clients in the 3X-UI API.

  

Public Methods for Clients:

- GetClientTrafficByEmail: Retrieves traffic for a client by email.

- GetClientTrafficByUUID: Retrieves traffic for a client by UUID.

- AddClient: Adds a new client to an inbound.

- ResetClientTraffic: Resets traffic for a client.

- UpdateClientInfo: Updates client information.

- GetClientIpAdress: Retrieves client's IP address.

- ClearClientIps: Clears IPs assigned to a client.

- DeleteClient: Deletes a client.

- GetKey: Generates a client key.

- GetSubscriptionLink: Generates a subscription link.

- GenerateUUID: Generates a unique UUID and subscription ID.

  

**Examples**:

```go

// See previous example for initialization

inbound, _ := apiClient.GetInbound(ctx, 1)

client := inbound.GetClientByEmail(apiClient, "your-email-string")

```

  

<a  id="api.api_client.Api.GetClientTrafficByEmail"></a>

#### GetClientTrafficByEmail

  

```go

func (a *Api) GetClientTrafficByEmail(ctx  context.Context, client *Client) (*GetClientTrafficResponse, error)

```

  

This method returns the amount of downloaded and uploaded data by a client identified by email.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `client`  _*Client_ - The client to retrieve traffic for.

  

**Returns**:

-  `*GetClientTrafficResponse` - A pointer to the traffic response.

-  `error` - An error if the request fails.

  

**Examples**:

```go

resp, err := apiClient.GetClientTrafficByEmail(ctx, &api.Client{Email: "your-email"})

if  err != nil {

log.Fatal(err)

}

log.Printf("Up: %d, Down: %d", client.Up, client.Down)

```

  

<a  id="api.api_client.Api.GetClientTrafficByUUID"></a>

#### GetClientTrafficByUUID

  

```go

func (a *Api) GetClientTrafficByUUID(ctx  context.Context, client *Client) (*GetClientTrafficResponse, error)

```

  

This method returns the amount of downloaded and uploaded data by a client identified by UUID.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `client`  _*Client_ - The client to retrieve traffic for.

  

**Returns**:

-  `*GetClientTrafficResponse` - A pointer to the traffic response.

-  `error` - An error if the request fails.

  

**Examples**:

```go

resp, err := apiClient.GetClientTrafficByUUID(ctx, &api.Client{UUID: "your-uuid"})

if  err != nil {

log.Fatal(err)

}

```

  

<a  id="api.api_client.Api.AddClient"></a>

#### AddClient

  

```go

func (a *Api) AddClient(ctx  context.Context, inbound *Inbound, client *Client) (*MessageResponse, error)

```

  

This method creates a new client inside the selected inbound.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `inbound`  _*Inbound_ - The inbound to add the client to.

-  `client`  _*Client_ - The client to add.

  

**Returns**:

-  `*MessageResponse` - A pointer to the response message.

-  `error` - An error if the request fails.

  

**Examples**:

```go

client := &api.Client{


Email: "your-email",

Enable: true,

ExpiryTime: api.UnixTime{Time: time.Now().AddDate(0, 1, 0)},

}

client.GenerateUUID()

resp, err := apiClient.AddClient(ctx, inbound, client)

if  err != nil {

log.Fatal(err)

}

```

  

<a  id="api.api_client.Api.ResetClientTraffic"></a>

#### ResetClientTraffic

  

```go

func (a *Api) ResetClientTraffic(ctx  context.Context, client *Client) (*MessageResponse, error)

```

  

This method resets the selected client's traffic.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `client`  _*Client_ - The client to reset traffic for.

  

**Returns**:

-  `*MessageResponse` - A pointer to the response message.

-  `error` - An error if the request fails.

  

**Examples**:

```go

resp, err := apiClient.ResetClientTraffic(ctx, client)

if  err != nil {

log.Fatal(err)

}

```

  

<a  id="api.api_client.Api.UpdateClientInfo"></a>

#### UpdateClientInfo

  

```go

func (a *Api) UpdateClientInfo(ctx  context.Context, client *Client) (*MessageResponse, error)

```

  

This method updates information about a client.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `client`  _*Client_ - The client to update.

  

**Returns**:

-  `*MessageResponse` - A pointer to the response message.

-  `error` - An error if the request fails.

  

**Examples**:

```go

resp, err := apiClient.UpdateClientInfo(ctx, client)

if  err != nil {

log.Fatal(err)

}

```

  

<a  id="api.api_client.Api.GetClientIpAdress"></a>

#### GetClientIpAdress

  

```go

func (a *Api) GetClientIpAdress(ctx  context.Context, client *Client) (string, error)

```

  

This method returns the client's IP address.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `client`  _*Client_ - The client to get IP for.

  

**Returns**:

-  `string` - The IP address.

-  `error` - An error if the request fails.

  

**Examples**:

```go

ip, err := apiClient.GetClientIpAdress(ctx, client)

if  err != nil {

log.Fatal(err)

}

log.Println(ip)

```

  

<a  id="api.api_client.Api.ClearClientIps"></a>

#### ClearClientIps

  

```go

func (a *Api) ClearClientIps(ctx  context.Context, client *Client) (*MessageResponse, error)

```

  

This method clears the IP assigned to a client.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `client`  _*Client_ - The client to clear IPs for.

  

**Returns**:

-  `*MessageResponse` - A pointer to the response message.

-  `error` - An error if the request fails.

  

**Examples**:

```go

resp, err := apiClient.ClearClientIps(ctx, client)

if  err != nil {

log.Fatal(err)

}

```

  

<a  id="api.api_client.Api.DeleteClient"></a>

#### DeleteClient

  

```go

func (a *Api) DeleteClient(ctx  context.Context, client *Client) (*MessageResponse, error)

```

  

This method deletes a client from an inbound.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `client`  _*Client_ - The client to delete.

  

**Returns**:

-  `*MessageResponse` - A pointer to the response message.

-  `error` - An error if the request fails.

  

**Examples**:

```go

resp, err := apiClient.DeleteClient(ctx, client)

if  err != nil {

log.Fatal(err)

}

```

  

<a  id="api.api_client.Api.GetKey"></a>

#### GetKey

  

```go

func (a *Api) GetKey(ctx  context.Context, client *Client) (string, error)

```

  

This method generates a key that can be used in suitable software.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `client`  _*Client_ - The client to generate key for.

  

**Returns**:

-  `string` - The generated key.

-  `error` - An error if the request fails or inbound doesn't exist.

  

**Examples**:

```go

key, err := apiClient.GetKey(ctx, client)

if  err != nil {

log.Fatal(err)

}

log.Println(key)

```

  

<a  id="api.api_client.Api.GetSubscriptionLink"></a>

#### GetSubscriptionLink

  

```go

func (a *Api) GetSubscriptionLink(ctx  context.Context, client *Client) (string, error)

```

  

This method generates a link to the client's subscription info.

  

**Arguments**:

-  `ctx`  _context.Context_ - The context for the request.

-  `client`  _*Client_ - The client to generate link for.

  

**Returns**:

-  `string` - The subscription link.

-  `error` - An error if SubId is invalid or parsing fails.

  

**Examples**:

```go

link, err := apiClient.GetSubscriptionLink(ctx, client)

if  err != nil {

log.Fatal(err)

}

log.Println(link)

```

<a  id="api.client.Client.GenerateUUID"></a>

#### GenerateUUID

  

```go

func (c *Client) GenerateUUID() error

```

  

This method generates a unique UUID and SubId for the client.

  

**Returns**:

-  `error` - Always nil (for future-proofing).

  

**Examples**:

```go

client := &api.Client{

Email: "example@mail.com",

}

err := client.GenerateUUID()

if  err != nil {

log.Fatal(err)

}

log.Println(client.UUID, client.SubId)

```