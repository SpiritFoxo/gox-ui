
# Gox-ui

**gox-ui** — a lightweight, idiomatic Go client library for interacting with **3X-UI** panel.

It provides a clean, type-safe API to manage inbounds, clients, traffic statistics, and subscriptions - everything you need to automate user provisioning, monitor usage, generate configs, or build your own management tools and bots.


## Installation

Install library with

```go
go get github.com/SpiritFoxo/gox-ui
```
    
## Usage/Examples

To use API you need to initialize api object

```go
package main

import (
	"context"
	"log"

	api "github.com/SpiritFoxo/gox-ui"
)

func main() {
	cfg := api.Config{
		BaseURL:         "https://your-host.com/your-panel-path",
		Username:        "admin",
		Password:        "admin",
		SubscriptionURI: "your-subscription-path", //optional field. can be left blank if subscription service disabled in panel settings
		Port:            12345, //optional field, if not provided defaults to 2053
        SubscriptionPort: 65432, //optional field, if not provided defaults to 2096
	}

	apiClient, err := api.NewApi(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
}
```

Then you will be able to call any other method provided by gox-ui

### Getting existing inbound
```go
package main

import (
	"context"
	"log"

	api "github.com/SpiritFoxo/gox-ui"
)

func main() {
	cfg := api.Config{
		BaseURL:         "https://your-host.com/your-panel-path",
		Username:        "admin",
		Password:        "admin",
		SubscriptionURI: "your-subscription-path",
		Port:            12345,
	}

	apiClient, err := api.NewApi(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

    inbound, err := apiClient.GetInbound(ctx, 1) // 1 - needed inbound id. If inbound exists you`ll get full model
	if err != nil {
		log.Fatal(err)
	}
}
```

### Getting existing client
```go
package main

import (
	"context"
	"log"

	api "github.com/SpiritFoxo/gox-ui"
)

func main() {
	cfg := api.Config{
		BaseURL:         "https://your-host.com/your-panel-path",
		Username:        "admin",
		Password:        "admin",
		SubscriptionURI: "your-subscription-path",
		Port:            12345,
	}

	apiClient, err := api.NewApi(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

    //in order to get client you need to get inbound first
    inbound, err := apiClient.GetInbound(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

    client := inbound.GetClientByEmail(apiClient, "your-email-string") //client object is containing every field except for traffic info
    //if you also need to get its traffic use:
    apiClient.GetClientTrafficByEmail(ctx, client) //or GetClientTrafficByUUID

}
```
### Creating new client

```go
func main() {
	cfg := api.Config{
		BaseURL:         "https://your-host.com/your-panel-path",
		Username:        "admin",
		Password:        "admin",
		SubscriptionURI: "your-subscription-path",
		Port:            12345,
	}

	apiClient, err := api.NewApi(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

    inbound, err := apiClient.GetInbound(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
    
    //creating new client object
    client := &api.Client{
		InboundId: inbound.ID,
		Email:     "your-new-client@mail.mail",
		Enable:    true,
        ExpiryTime: api.UnixTime{Time: time.Now().AddDate(0, 1, 0)} //will be expired after 1 month
	}
	client.GenerateUUID() //generates unique uuid and subscription id
	apiClient.AddClient(ctx, inbound, client) //adds client to the provided inbound
}
```

### Getting key or subscription link
```go
package main

import (
	"context"
    "fmt"
	"log"

	api "github.com/SpiritFoxo/gox-ui"
)

func main() {
	cfg := api.Config{
		BaseURL:         "https://your-host.com/your-panel-path",
		Username:        "admin",
		Password:        "admin",
		SubscriptionURI: "your-subscription-path",
		Port:            12345,
	}

	apiClient, err := api.NewApi(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

    inbound, err := apiClient.GetInbound(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

    client := inbound.GetClientByEmail(apiClient, "your-email-string")

    //getting subscription link
    subscriptionLink, err := apiClient.GetSubscriptionLink(ctx, client)
    if err != nil {
        log.Fatal("error")
    }
    fmt.Println(subscriptionLink)

    //getting client`s key
    key, err := apiClient.GetKey(ctx, client)
    if err != nil {
        log.Fatal("error")
    }
    fmt.Println(key)

}
```

