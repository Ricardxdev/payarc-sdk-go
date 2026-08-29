package outputs

import "github.com/Ricardxdev/payarc-sdk-go/pkg/extra"

// Plan represents a PayArc plan as returned by the list, retrieve and
// update endpoints.
type Plan struct {
	Object              string        `json:"object"`
	PlanID              string        `json:"plan_id"`
	Amount              int64         `json:"amount"`
	Interval            string        `json:"interval"`
	IntervalCount       int64         `json:"interval_count"`
	Name                string        `json:"name"`
	PlanDescription     *string       `json:"plan_description"`
	StatementDescriptor *string       `json:"statement_descriptor"`
	TrialPeriodDays     extra.FlexInt `json:"trial_period_days"`
	Currency            string        `json:"currency"`
	CreatedAt           int64         `json:"created_at"`
	UpdatedAt           int64         `json:"updated_at"`
}

type PlanResponse struct {
	Data Plan `json:"data"`
	Meta Meta `json:"meta"`
}

type PlansResponse struct {
	Data     []Plan   `json:"data"`
	Metadata Metadata `json:"meta"`
}

// PlanTimestamp represents the date object returned by the create plan
// endpoint instead of the usual unix timestamp.
type PlanTimestamp struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}

// CreatedPlan represents a PayArc plan as returned by the create endpoint.
type CreatedPlan struct {
	Object              string        `json:"object"`
	ID                  string        `json:"id"`
	RealID              int64         `json:"real_id"`
	Amount              extra.FlexInt `json:"amount"`
	Interval            string        `json:"interval"`
	IntervalCount       extra.FlexInt `json:"interval_count"`
	Name                string        `json:"name"`
	Description         *string       `json:"description"`
	StatementDescriptor string        `json:"statement_descriptor"`
	TrialPeriodDays     extra.FlexInt `json:"trial_period_days"`
	Currency            string        `json:"currency"`
	CreatedAt           PlanTimestamp `json:"created_at"`
	UpdatedAt           PlanTimestamp `json:"updated_at"`
}

type CreatePlanResponse struct {
	Data CreatedPlan `json:"data"`
	Meta Meta        `json:"meta"`
}
