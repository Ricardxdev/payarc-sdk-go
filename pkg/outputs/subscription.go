package outputs

// Subscription represents a PayArc subscription.
type Subscription struct {
	Object                string   `json:"object"`
	ID                    string   `json:"id"`
	CustomerID            string   `json:"customer_id"`
	ApplicationFeePercent *float64 `json:"application_fee_percent"`
	BillingType           int      `json:"billing_type"`
	PaymentDueDays        *int64   `json:"payment_due_days"`
	CancelAtPeriodEnd     int      `json:"cancel_at_period_end"`
	CanceledAt            *int64   `json:"canceled_at"`
	CurrentPeriodEnd      *int64   `json:"current_period_end"`
	CurrentPeriodStart    *int64   `json:"current_period_start"`
	DaysUntilDue          *int64   `json:"days_until_due"`
	PlanRef               string   `json:"plan_ref"`
	StartAt               *int64   `json:"start_at"`
	EndAt                 *int64   `json:"end_at"`
	Status                string   `json:"status"`
	Message               string   `json:"message"`
	TaxPercent            float64  `json:"tax_percent"`
	TrialEnd              *int64   `json:"trial_end"`
	TrialDays             int64    `json:"trial_days"`
	TrialStart            *int64   `json:"trial_start"`
	CreatedAt             *int64   `json:"created_at"`
	UpdatedAt             *int64   `json:"updated_at"`
	Description           *string  `json:"description"`
	PaymentType           *string  `json:"payment_type"`
	PausedAt              *int64   `json:"paused_at"`
	PauseEndAt            *int64   `json:"pause_end_at"`
	ResumedAt             *int64   `json:"resumed_at"`
	PauseUnitsAllowed     []string `json:"pause_units_allowed"`
	Plan                  struct {
		Data Plan `json:"data"`
	} `json:"plan"`
	Discount *struct {
		Data Coupon `json:"data"`
	} `json:"discount,omitempty"`
	Customer *struct {
		Data Customer `json:"data"`
	} `json:"customer,omitempty"`
}

type SubscriptionResponse struct {
	Data Subscription `json:"data"`
	Meta Meta         `json:"meta"`
}

type SubscriptionsResponse struct {
	Data     []Subscription `json:"data"`
	Metadata Metadata       `json:"meta"`
}
