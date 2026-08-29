# PayArc SDK

A Go package for working with PayArc. This package provides a client to interact with charges, customers, cards, and tokens.

## Features

- Create, retrieve, and list charges
- Create and fetch customers
- Create cards and tokens
- Create, retrieve, list, update, delete, and export plans
- Create, retrieve, list, and delete coupons
- Create, list, update, cancel, pause, resume, delete, and export subscriptions

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

### Plans

```go
// Create a plan
plan, err := payArcClient.CreatePlan(inputs.CreatePlanDTO{
    Amount:              2900,
    PlanType:            extra.PlanTypeDigital,
    Name:                "Daily yoga",
    Interval:            extra.PlanIntervalDay,
    StatementDescriptor: "A1 Company 888-888-8888",
})

// List, retrieve, update and delete plans
plans, err := payArcClient.GetPlans(1, 10)
plan, err := payArcClient.GetPlan("plan_7012f470")
updated, err := payArcClient.UpdatePlan("plan_7012f470", inputs.UpdatePlanDTO{Name: "New name"})
err = payArcClient.DeletePlan("plan_7012f470")

// Export plans to Excel (returns the file contents)
file, err := payArcClient.ExportPlans(inputs.ExportPlansInput{
    SelectedColumns: []string{"plan_id", "name", "trial_period_days", "amount"},
})
```

### Coupons

```go
// Create a coupon
amountOff := int64(1000)
coupon, err := payArcClient.CreateCoupon(inputs.CreateCouponDTO{
    Name:      "my coupon",
    Duration:  extra.CouponDurationOnce,
    Currency:  extra.CurrencyUSD,
    AmountOff: &amountOff,
})

// List, retrieve and delete coupons
coupons, err := payArcClient.GetCoupons(1, 10)
coupon, err := payArcClient.GetCoupon("dc_75d302b4")
err = payArcClient.DeleteCoupon("dc_75d302b4")
```

### Subscriptions

```go
// Create a subscription
subscription, err := payArcClient.CreateSubscription(inputs.CreateSubscriptionDTO{
    CustomerID: "jDPAnVxA4xPDVpKM",
    PlanID:     "plan_d2b50d00",
})

// List subscriptions with filters
subscriptions, err := payArcClient.GetSubscriptions(inputs.ListSubscriptionsInput{
    Limit:    10,
    Page:     1,
    StatusEq: extra.SubscriptionStatusActive,
})

// Update, cancel, pause, resume and delete subscriptions
updated, err := payArcClient.UpdateSubscription("Vg0rxj00XXljPAoX", inputs.UpdateSubscriptionDTO{
    Description: "New description",
})
canceled, err := payArcClient.CancelSubscription("Vg0rxj00XXljPAoX")
paused, err := payArcClient.PauseSubscription("Vg0rxj00XXljPAoX", inputs.PauseSubscriptionDTO{
    Duration: 2,
    Unit:     extra.PauseUnitWeeks,
})
resumed, err := payArcClient.ResumeSubscription("Vg0rxj00XXljPAoX")
err = payArcClient.DeleteSubscription("Vg0rxj00XXljPAoX")

// Export subscriptions to Excel (returns the file contents)
file, err := payArcClient.ExportSubscriptions(inputs.ExportSubscriptionsInput{
    SelectedColumns: []string{"customer", "status", "billing_type", "plan", "created_at"},
})
```
