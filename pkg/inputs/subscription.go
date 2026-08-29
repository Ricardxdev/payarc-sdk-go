package inputs

import "github.com/Ricardxdev/payarc-sdk-go/pkg/extra"

// CreateSubscriptionDTO represents the parameters required to create a subscription.
type CreateSubscriptionDTO struct {
	CustomerID     string `json:"customer_id" form:"customer_id"`                     //  - CustomerID: An ID against which a new subscription will be created.
	PlanID         string `json:"plan_id" form:"plan_id"`                             //  - PlanID: A plan ID against which a new subscription is created.
	StartAfterDays *int   `json:"start_after_days,omitempty" form:"start_after_days"` //  - StartAfterDays: Days after which the subscription will start (1 - 365).
	TrialDays      *int   `json:"trial_days,omitempty" form:"trial_days"`             //  - TrialDays: Number of days until the subscription is under the trial period (0 - 365).
	EndAfterCycles *int   `json:"end_after_cycles,omitempty" form:"end_after_cycles"` //  - EndAfterCycles: Number of cycles after which a subscription will end (1 - 999).
	DiscountID     string `json:"discount_id,omitempty" form:"discount_id"`           //  - DiscountID: An ID that refers to a coupon when applying to a subscription.
	BillingType    *int   `json:"billing_type,omitempty" form:"billing_type"`         //  - BillingType: 1 for 'Automatic' (charged with the default card on file), 0 for 'Manual' (payment link over email). Default 1.
	PaymentDueDays *int   `json:"payment_due_days,omitempty" form:"payment_due_days"` //  - PaymentDueDays: Number of days the payment is due after the invoice is sent. Required if BillingType is 0.
	Description    string `json:"description,omitempty" form:"description"`           //  - Description: An arbitrary string that describes a subscription (5 - 200 characters).
}

// UpdateSubscriptionDTO represents the parameters required to update a subscription.
type UpdateSubscriptionDTO struct {
	Description string   `json:"description,omitempty" form:"description"` //  - Description: An arbitrary string that describes the subscription.
	TaxPercent  *float64 `json:"tax_percent,omitempty" form:"tax_percent"` //  - TaxPercent: The percentage of tax applied to a subscription (0 - 100). Decimal values are allowed.
}

// PauseSubscriptionDTO represents the parameters required to pause a subscription.
// The allowed pause unit depends on the subscription plan interval: day plans
// allow days, week plans allow weeks, month plans allow months, and year plans
// allow years.
type PauseSubscriptionDTO struct {
	Duration int             `json:"duration" form:"duration"` //  - Duration: Pause duration value. Min 1. Max: 365 days, 52 weeks, 12 months, 1 year.
	Unit     extra.PauseUnit `json:"unit" form:"unit"`         //  - Unit: Pause duration unit (days, weeks, months, years).
}

// ListSubscriptionsInput represents the query parameters used to list subscriptions.
type ListSubscriptionsInput struct {
	Limit              int                        `query:"limit,omitempty"`                   //  - Limit: Number of subscription records to return per page.
	Page               int                        `query:"page,omitempty"`                    //  - Page: Page number of the result set to return.
	OrderBy            string                     `query:"orderBy,omitempty"`                 //  - OrderBy: Column name to sort the results by (e.g. created_at). Multiple fields separated by a semicolon (;).
	SortedBy           extra.SortOrder            `query:"sortedBy,omitempty"`                //  - SortedBy: Direction that OrderBy is sorted by (asc, desc).
	Include            string                     `query:"include,omitempty"`                 //  - Include: Embeds the related customer or discount object in each subscription (customer, discount).
	StatusIn           []extra.SubscriptionStatus `query:"status[in],omitempty"`              //  - StatusIn: Filter subscriptions by multiple statuses (comma separated).
	StatusEq           extra.SubscriptionStatus   `query:"status[eq],omitempty"`              //  - StatusEq: Filters subscriptions by an exact status.
	CreatedInLastDays  int                        `query:"created_at[intl],omitempty"`        //  - CreatedInLastDays: Returns subscriptions created within the last N days.
	CreatedAtEq        string                     `query:"created_at[eq],omitempty"`          //  - CreatedAtEq: Returns subscriptions created on the specified date. Format: YYYY/MM/DD.
	CreatedAtGte       string                     `query:"created_at[gte],omitempty"`         //  - CreatedAtGte: Returns subscriptions created on or after the specified date. Format: YYYY/MM/DD.
	CreatedAtLte       string                     `query:"created_at[lte],omitempty"`         //  - CreatedAtLte: Returns subscriptions created on or before the specified date. Format: YYYY/MM/DD.
	CreatedAtGt        string                     `query:"created_at[gt],omitempty"`          //  - CreatedAtGt: Returns subscriptions created after the specified date. Format: YYYY/MM/DD.
	CreatedAtLt        string                     `query:"created_at[lt],omitempty"`          //  - CreatedAtLt: Returns subscriptions created before the specified date. Format: YYYY/MM/DD.
	CreatedAtRangeFrom string                     `query:"created_at[range][from],omitempty"` //  - CreatedAtRangeFrom: Start of the creation date range (inclusive). Format: YYYY/MM/DD.
	CreatedAtRangeTo   string                     `query:"created_at[range][to],omitempty"`   //  - CreatedAtRangeTo: End of the creation date range (inclusive). Format: YYYY/MM/DD.
	PaymentType        string                     `query:"payment_type[eq],omitempty"`        //  - PaymentType: Filters subscriptions by payment method (card, ach).
	Plan               string                     `query:"plan,omitempty"`                    //  - Plan: Filters subscriptions by plan, using the plan's code.
	CustomerID         string                     `query:"customer_id,omitempty"`             //  - CustomerID: Filters subscriptions belonging to a specific customer.
	Excel              *bool                      `query:"excel,omitempty"`                   //  - Excel: If true, exports all matching results without pagination.
}

// ExportSubscriptionsInput represents the query parameters used to export
// subscriptions to Excel.
type ExportSubscriptionsInput struct {
	SelectedColumns []string `query:"selected_columns"`          //  - SelectedColumns: The selected columns for export. Ex. customer,status,billing_type,plan,created_at (required).
	Limit           int      `query:"limit,omitempty"`           //  - Limit: Number of records to be listed per page (default 25).
	Page            int      `query:"page,omitempty"`            //  - Page: A number that represents the current page number (default 1).
	Include         string   `query:"include,omitempty"`         //  - Include: A string that represents the included object (default customer).
	CreatedAtGte    string   `query:"created_at[gte],omitempty"` //  - CreatedAtGte: Subscriptions created on or after the date. Format: YYYY/MM/DD.
	CreatedAtLte    string   `query:"created_at[lte],omitempty"` //  - CreatedAtLte: Subscriptions created on or before the date. Format: YYYY/MM/DD.
}
