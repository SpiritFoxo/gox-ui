
# gox-ui

**gox-ui** is a lightweight, idiomatic Go client library for interacting with the **3X-UI** panel. It offers a clean, type-safe API to manage inbounds, clients, traffic statistics, and subscriptions—perfect for automating user provisioning, monitoring usage, generating configs, or building custom management tools and bots.

## Installation

Install the library using:

```bash
go get github.com/SpiritFoxo/gox-ui
```

## Quick Start

### Initialization and Login

Start by importing the library and creating an `Api` instance with your configuration. Then, perform a login to authenticate.

```go
package main

import (
	"context"
	"log"

	api "github.com/SpiritFoxo/gox-ui"
)

func main() {
	cfg := api.Config{
		BaseURL:          "https://your-host.com:2053/your-panel-path",
		Username:         "admin",
		Password:         "admin",
		SubscriptionURI:  "your-subscription-path", // Optional: Leave blank if subscriptions are disabled
		SubscriptionPort: 2096,                     // Optional: Defaults to 2096
		// HTTPClient:    &http.Client{},           // Optional: Custom HTTP client
		// Timeout:       30 * time.Second,         // Optional: Defaults to 30s
	}
	apiClient, err := api.NewApi(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = apiClient.Login(ctx) // Use apiClient.Login(ctx, "123456") if 2FA is enabled
	if err != nil {
		log.Fatal(err)
	}

	// Now you're ready to use the API!
}
```

### Example: Listing Inbounds

Retrieve a list of all inbounds:

```go
inbounds, err := apiClient.ListInbounds(ctx)
if err != nil {
	log.Fatal(err)
}
for _, inbound := range *inbounds {
	log.Printf("Inbound ID: %d, Remark: %s", inbound.ID, inbound.Remark)
}
```

### Example: Adding a New Client

Add a client to an existing inbound:

```go
inbound, err := apiClient.GetInbound(ctx, 1) // Replace 1 with your inbound ID
if err != nil {
	log.Fatal(err)
}

client := &api.Client{
	Email:      "new-client@example.com",
	Enable:     true,
	ExpiryTime: api.UnixTime{Time: time.Now().AddDate(0, 1, 0)}, // Expires in 1 month
	TotalGB:    10 * 1024 * 1024 * 1024,                         // 10 GB limit
}
err = client.GenerateUUID()
if err != nil {
	log.Fatal(err)
}

resp, err := apiClient.AddClient(ctx, inbound, client)
if err != nil {
	log.Fatal(err)
}
log.Printf("Client added: %s", resp.Msg)
```

## Documentation

For full API documentation, including all methods, structs, and examples, see [docs](https://github.com/SpiritFoxo/gox-ui/tree/main/docs).
