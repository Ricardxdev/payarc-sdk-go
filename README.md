# PayArc SDK

A Go package for working with PayArc. This package provides a client to interact with charges, customers, cards, and tokens.

## Features

- Create, retrieve, and list charges
- Create and fetch customers
- Create cards and tokens

## Installation

```bash
go get github.com/Ricardxdev/payarc-sdk-go
```

## Usage

```go
package main

import (
    "context"
    "fmt"

    "github.com/Ricardxdev/payarc-sdk-go/payarcsdk"
    "github.com/Ricardxdev/payarc-sdk-go/client"
    "github.com/Ricardxdev/payarc-sdk-go/inputs"
)

func main() {
    // Create options for the PayArc client
    options := payarcsdk.PayarcClientOptions{
        BaseUrl:    "https://api.payarc.com",
        ApiVersion: "v1",
        Token:      "YOUR_API_KEY",
        HTTPClient: client.NewHTTPClient(),
    }

    // Create a new PayArc client
    payArcClient := payarcsdk.NewPayarcClient(context.Background(), options)

    // Example: Create a new charge
    chargeInput := inputs.ChargeInput{
        Amount:   1000,
        Currency: "USD",
        // other fields
    }

    chargeResponse, err := payArcClient.CreateCharge(chargeInput)
    if err != nil {
        panic(err)
    }
    fmt.Println("Charge:", chargeResponse)
}
```