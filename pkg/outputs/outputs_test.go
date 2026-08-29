package outputs

import (
	"encoding/json"
	"testing"
)

func TestPlansResponseUnmarshal(t *testing.T) {
	payload := `{
		"data": [
			{
				"object": "Plan",
				"plan_id": "plan_7012f470",
				"amount": 99,
				"interval": "day",
				"interval_count": 1,
				"name": "Test",
				"plan_description": null,
				"statement_descriptor": "dfghdfhg",
				"trial_period_days": "0",
				"currency": "usd",
				"created_at": 1724503328,
				"updated_at": 1724503328
			}
		],
		"meta": {
			"include": [],
			"custom": [],
			"pagination": {
				"total": 6,
				"count": 6,
				"per_page": 10,
				"current_page": 1,
				"total_pages": 1,
				"links": {}
			}
		}
	}`

	res := &PlansResponse{}
	if err := json.Unmarshal([]byte(payload), res); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(res.Data) != 1 {
		t.Fatalf("expected 1 plan, got %d", len(res.Data))
	}

	plan := res.Data[0]
	if plan.PlanID != "plan_7012f470" {
		t.Errorf("expected plan_id plan_7012f470, got %s", plan.PlanID)
	}
	if plan.Amount != 99 {
		t.Errorf("expected amount 99, got %d", plan.Amount)
	}
	if plan.TrialPeriodDays.Int64() != 0 {
		t.Errorf("expected trial_period_days 0, got %d", plan.TrialPeriodDays.Int64())
	}
	if res.Metadata.Pagination.Total != 6 {
		t.Errorf("expected pagination total 6, got %d", res.Metadata.Pagination.Total)
	}
}

func TestCreatePlanResponseUnmarshal(t *testing.T) {
	payload := `{
		"data": {
			"object": "Plan",
			"id": "daily_yoga",
			"real_id": 9,
			"amount": "2900",
			"interval": "day",
			"interval_count": "1",
			"name": "Daily yoga",
			"description": null,
			"statement_descriptor": "sample statement descriptor",
			"trial_period_days": "3",
			"currency": "USD",
			"created_at": {
				"date": "2018-11-27 09:15:59.000000",
				"timezone_type": 3,
				"timezone": "UTC"
			},
			"updated_at": {
				"date": "2018-11-27 09:15:59.000000",
				"timezone_type": 3,
				"timezone": "UTC"
			}
		},
		"meta": {
			"include": [],
			"custom": []
		}
	}`

	res := &CreatePlanResponse{}
	if err := json.Unmarshal([]byte(payload), res); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.Data.ID != "daily_yoga" {
		t.Errorf("expected id daily_yoga, got %s", res.Data.ID)
	}
	if res.Data.Amount.Int64() != 2900 {
		t.Errorf("expected amount 2900, got %d", res.Data.Amount.Int64())
	}
	if res.Data.TrialPeriodDays.Int64() != 3 {
		t.Errorf("expected trial_period_days 3, got %d", res.Data.TrialPeriodDays.Int64())
	}
	if res.Data.CreatedAt.Timezone != "UTC" {
		t.Errorf("expected timezone UTC, got %s", res.Data.CreatedAt.Timezone)
	}
}

func TestCouponsResponseUnmarshal(t *testing.T) {
	payload := `{
		"data": [
			{
				"object": "Discount",
				"discount_id": "dc_75d302b4",
				"amount_off": 10,
				"percent_off": null,
				"duration": "repeating",
				"duration_in_months": null,
				"max_redemptions": null,
				"redeem_by": null,
				"times_redeemed": null,
				"status": "active",
				"name": "testers",
				"created_at": 1724584599,
				"updated_at": 1724584599,
				"id": "WzA1lYjlWMWYDg5M"
			}
		],
		"meta": {
			"include": [],
			"custom": [],
			"pagination": {
				"total": 7,
				"count": 7,
				"per_page": 10,
				"current_page": 1,
				"total_pages": 1,
				"links": {}
			}
		}
	}`

	res := &CouponsResponse{}
	if err := json.Unmarshal([]byte(payload), res); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(res.Data) != 1 {
		t.Fatalf("expected 1 coupon, got %d", len(res.Data))
	}

	coupon := res.Data[0]
	if coupon.DiscountID != "dc_75d302b4" {
		t.Errorf("expected discount_id dc_75d302b4, got %s", coupon.DiscountID)
	}
	if coupon.AmountOff == nil || *coupon.AmountOff != 10 {
		t.Errorf("expected amount_off 10, got %v", coupon.AmountOff)
	}
	if coupon.PercentOff != nil {
		t.Errorf("expected percent_off nil, got %v", *coupon.PercentOff)
	}
}

func TestSubscriptionResponseUnmarshal(t *testing.T) {
	payload := `{
		"data": {
			"object": "Subscription",
			"id": "Vg0rxj00XXljPAoX",
			"customer_id": "jDPAnVxA4xPDVpKM",
			"application_fee_percent": null,
			"billing_type": 1,
			"payment_due_days": null,
			"cancel_at_period_end": 0,
			"canceled_at": null,
			"current_period_end": 1726490792,
			"current_period_start": 1726145192,
			"days_until_due": null,
			"plan_ref": "n2pv6j6p2pvj3Ygq",
			"start_at": 1726490792,
			"end_at": null,
			"status": "trial",
			"tax_percent": 0,
			"trial_end": 1726490792,
			"trial_days": 4,
			"trial_start": 1726145192,
			"created_at": 1726145192,
			"updated_at": 1726145192,
			"description": null,
			"payment_type": "card",
			"paused_at": null,
			"pause_end_at": null,
			"resumed_at": null,
			"pause_units_allowed": ["days"],
			"plan": {
				"data": {
					"object": "Plan",
					"plan_id": "plan_d2b50d00",
					"amount": 99,
					"interval": "day",
					"interval_count": 1,
					"name": "test",
					"plan_description": null,
					"statement_descriptor": "dfghdfhg",
					"trial_period_days": "4",
					"currency": "usd",
					"created_at": 1724503321,
					"updated_at": 1724505447
				}
			}
		},
		"meta": {
			"include": ["plan", "discount", "customer"],
			"custom": []
		}
	}`

	res := &SubscriptionResponse{}
	if err := json.Unmarshal([]byte(payload), res); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	sub := res.Data
	if sub.ID != "Vg0rxj00XXljPAoX" {
		t.Errorf("expected id Vg0rxj00XXljPAoX, got %s", sub.ID)
	}
	if sub.Status != "trial" {
		t.Errorf("expected status trial, got %s", sub.Status)
	}
	if sub.CurrentPeriodEnd == nil || *sub.CurrentPeriodEnd != 1726490792 {
		t.Errorf("expected current_period_end 1726490792, got %v", sub.CurrentPeriodEnd)
	}
	if sub.Plan.Data.PlanID != "plan_d2b50d00" {
		t.Errorf("expected plan_id plan_d2b50d00, got %s", sub.Plan.Data.PlanID)
	}
	if sub.Plan.Data.TrialPeriodDays.Int64() != 4 {
		t.Errorf("expected plan trial_period_days 4, got %d", sub.Plan.Data.TrialPeriodDays.Int64())
	}
	if len(sub.PauseUnitsAllowed) != 1 || sub.PauseUnitsAllowed[0] != "days" {
		t.Errorf("expected pause_units_allowed [days], got %v", sub.PauseUnitsAllowed)
	}
}

func TestCancelSubscriptionResponseUnmarshal(t *testing.T) {
	payload := `{
		"data": {
			"object": "Subscription",
			"id": "xlRP0jgPllljAXog",
			"customer_id": "jDPAnVxA4xPDVpKM",
			"application_fee_percent": null,
			"billing_type": 1,
			"payment_due_days": null,
			"cancel_at_period_end": 0,
			"canceled_at": 1726163436,
			"current_period_end": 1726407240,
			"current_period_start": 1726061640,
			"days_until_due": null,
			"plan_ref": "n2pv6j6p2pvj3Ygq",
			"start_at": 1726407240,
			"end_at": null,
			"status": "canceled",
			"message": "canceled - to be canceled on 2024-09-15",
			"tax_percent": 5.8,
			"trial_end": 1726407240,
			"trial_days": 4,
			"trial_start": 1726061640,
			"created_at": 1726061640,
			"updated_at": 1726163436,
			"description": null,
			"plan": {
				"data": {
					"object": "Plan",
					"plan_id": "plan_d2b50d00",
					"amount": 99,
					"interval": "day",
					"interval_count": 1,
					"name": "test",
					"plan_description": null,
					"statement_descriptor": "dfghdfhg",
					"trial_period_days": "4",
					"currency": "usd",
					"created_at": 1724503321,
					"updated_at": 1724505447
				}
			}
		},
		"meta": {
			"include": ["plan", "discount", "customer"],
			"custom": []
		}
	}`

	res := &SubscriptionResponse{}
	if err := json.Unmarshal([]byte(payload), res); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.Data.Status != "canceled" {
		t.Errorf("expected status canceled, got %s", res.Data.Status)
	}
	if res.Data.Message == "" {
		t.Error("expected message to be set")
	}
	if res.Data.TaxPercent != 5.8 {
		t.Errorf("expected tax_percent 5.8, got %f", res.Data.TaxPercent)
	}
	if res.Data.CanceledAt == nil || *res.Data.CanceledAt != 1726163436 {
		t.Errorf("expected canceled_at 1726163436, got %v", res.Data.CanceledAt)
	}
}
