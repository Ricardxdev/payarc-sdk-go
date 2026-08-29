package inputs

import "github.com/Ricardxdev/payarc-sdk-go/pkg/extra"

// CreatePlanDTO represents the parameters required to create a plan.
type CreatePlanDTO struct {
	Amount              int64              `json:"amount" form:"amount"`                                       //  - Amount: A positive integer in cents representing how much to charge (50 - 10000000).
	PlanType            extra.PlanType     `json:"plan_type" form:"plan_type"`                                 //  - PlanType: The type of plan (physical, digital).
	Name                string             `json:"name" form:"name"`                                           //  - Name: Name of the plan (3 - 50 characters).
	Interval            extra.PlanInterval `json:"interval" form:"interval"`                                   //  - Interval: The unit of interval for billing a plan (day, month, week, year).
	StatementDescriptor string             `json:"statement_descriptor" form:"statement_descriptor"`           //  - StatementDescriptor: An arbitrary string to be displayed on a plan (5 - 25 characters).
	IntervalCount       *int               `json:"interval_count,omitempty" form:"interval_count"`             //  - IntervalCount: The frequency of billing (days: 1-365, months: 1-12, weeks: 1-52, years: 1).
	TrialPeriodDays     *int               `json:"trial_period_days,omitempty" form:"trial_period_days"`       //  - TrialPeriodDays: Number of days of the trial period that a customer is not charged for (0 - 365).
	PlanID              string             `json:"plan_id,omitempty" form:"plan_id"`                           //  - PlanID: A custom ID for the plan. If empty, PAYARC generates one automatically.
	PlanDescription     string             `json:"plan_description,omitempty" form:"plan_description"`         //  - PlanDescription: An arbitrary string that describes a plan.
	Currency            extra.Currency     `json:"currency,omitempty" form:"currency"`                         //  - Currency: Three-letter ISO currency code, in lowercase.
	SurchargeApplicable *int               `json:"surcharge_applicable,omitempty" form:"surcharge_applicable"` //  - SurchargeApplicable: Enabling this option applies surcharge to the plan amount.
}

// UpdatePlanDTO represents the parameters required to update a plan.
type UpdatePlanDTO struct {
	Name                string `json:"name" form:"name"`                                           //  - Name: Name of the plan (3 - 50 characters).
	StatementDescriptor string `json:"statement_descriptor,omitempty" form:"statement_descriptor"` //  - StatementDescriptor: An arbitrary string to be displayed on a plan (5 - 25 characters).
	TrialPeriodDays     *int   `json:"trial_period_days,omitempty" form:"trial_period_days"`       //  - TrialPeriodDays: Number of days of the trial period that a customer is not charged for (0 - 365).
}

// ExportPlansInput represents the query parameters used to export plans.
type ExportPlansInput struct {
	SelectedColumns []string `query:"selected_columns"`          //  - SelectedColumns: Columns to include in the export. Ex. plan_id,name,trial_period_days,amount (required).
	Limit           int      `query:"limit,omitempty"`           //  - Limit: Number of plans to be listed per page (default 25).
	Page            int      `query:"page,omitempty"`            //  - Page: A number that represents a current page number (default 1).
	Excel           *bool    `query:"excel,omitempty"`           //  - Excel: Represents the export as excel (default true).
	CreatedAtGte    string   `query:"created_at[gte],omitempty"` //  - CreatedAtGte: Plans created on or after the date. Format: YYYY/MM/DD.
	CreatedAtLte    string   `query:"created_at[lte],omitempty"` //  - CreatedAtLte: Plans created before or equal to the date. Format: YYYY/MM/DD.
}
