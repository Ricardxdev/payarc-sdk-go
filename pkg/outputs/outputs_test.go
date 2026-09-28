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

func TestVoidChargeResponseUnmarshal(t *testing.T) {
	payload := `{
		"data": {
			"object": "Charge",
			"id": "DLbnOBbBDWMbbOoM",
			"amount": 99,
			"amount_approved": 99,
			"amount_refunded": 0,
			"amount_captured": 500,
			"amount_voided": 500,
			"application_fee_amount": 0,
			"tip_amount": 20,
			"payarc_fees": 0,
			"type": "Sale",
			"customer_email": "eric.yulfo@payarc.com",
			"net_amount": 0,
			"captured": 1,
			"is_refunded": 0,
			"status": "void",
			"auth_code": "TAS697",
			"failure_code": null,
			"failure_message": null,
			"charge_description": null,
			"kount_details": "",
			"kount_status": "",
			"statement_description": "Eric's Energy Drinks",
			"under_review": 0,
			"created_at": 1725459426,
			"updated_at": 1725461016,
			"card_level": "LEVEL1",
			"void_reason": "requested_by_customer",
			"void_description": "Customer requested order cancellation",
			"card": {
				"data": {
					"object": "Card",
					"id": "vmy905NNNm5y0M2L",
					"card_source": "PHONE",
					"is_default": 1,
					"exp_month": "12",
					"exp_year": "2025",
					"is_verified": 1,
					"brand": "V",
					"last4digit": "5439",
					"first6digit": 401200
				}
			}
		},
		"meta": {
			"include": [
				"review",
				"transaction_metadata"
			],
			"custom": []
		}
	}`

	res := &ResponseCharge{}
	if err := json.Unmarshal([]byte(payload), res); err != nil {
		t.Fatalf("unexpected error unmarshaling void charge response: %v", err)
	}

	if res.Data.ID != "DLbnOBbBDWMbbOoM" {
		t.Errorf("expected charge ID DLbnOBbBDWMbbOoM, got %s", res.Data.ID)
	}
	if res.Data.Status != "void" {
		t.Errorf("expected status void, got %s", res.Data.Status)
	}
	if res.Data.AmountVoided != 500 {
		t.Errorf("expected amount_voided 500, got %d", res.Data.AmountVoided)
	}
	if res.Data.VoidReason == nil || *res.Data.VoidReason != "requested_by_customer" {
		t.Errorf("expected void_reason requested_by_customer, got %v", res.Data.VoidReason)
	}
	if res.Data.VoidDescription == nil || *res.Data.VoidDescription != "Customer requested order cancellation" {
		t.Errorf("expected void_description Customer requested order cancellation, got %v", res.Data.VoidDescription)
	}
	if len(res.Metadata.Include) != 2 {
		t.Errorf("expected 2 meta includes, got %d", len(res.Metadata.Include))
	}
}

func TestCaptureChargeResponseUnmarshal(t *testing.T) {
	payload := `{
		"data": {
			"object": "Charge",
			"id": "DMWbOLMyDnLMyOBX",
			"amount": 99,
			"amount_approved": 99,
			"amount_refunded": 0,
			"amount_captured": 500,
			"amount_voided": 0,
			"application_fee_amount": 0,
			"tip_amount": 20,
			"payarc_fees": 0,
			"type": "Sale",
			"customer_email": "eric.yulfo@payarc.com",
			"net_amount": 500,
			"captured": 1,
			"is_refunded": 0,
			"status": "submitted_for_settlement",
			"auth_code": "TAS776",
			"statement_description": "Eric's Energy Drinks",
			"created_at": 1725459517,
			"updated_at": 1725459705,
			"card": {
				"data": {
					"object": "Card",
					"id": "vmy905NNNm5y0M2L",
					"card_source": "PHONE",
					"is_default": 1,
					"brand": "V",
					"last4digit": "5439"
				}
			}
		},
		"meta": {
			"include": [
				"review"
			],
			"custom": []
		}
	}`

	res := &ResponseCharge{}
	if err := json.Unmarshal([]byte(payload), res); err != nil {
		t.Fatalf("unexpected error unmarshaling capture charge response: %v", err)
	}

	if res.Data.ID != "DMWbOLMyDnLMyOBX" {
		t.Errorf("expected charge ID DMWbOLMyDnLMyOBX, got %s", res.Data.ID)
	}
	if res.Data.Status != "submitted_for_settlement" {
		t.Errorf("expected status submitted_for_settlement, got %s", res.Data.Status)
	}
	if res.Data.AmountCaptured != 500 {
		t.Errorf("expected amount_captured 500, got %d", res.Data.AmountCaptured)
	}
	if res.Data.TipAmount != 20 {
		t.Errorf("expected tip_amount 20, got %d", res.Data.TipAmount)
	}
	if !res.Data.Captured.AsBool() {
		t.Errorf("expected captured to be true")
	}
}

func TestRefundChargeResponseUnmarshal(t *testing.T) {
	payload := `{
		"data": {
			"object": "Charge",
			"id": "DLbnOBbBDWMRROoM",
			"amount": 99,
			"amount_approved": 99,
			"amount_refunded": 99,
			"amount_captured": 99,
			"amount_voided": 0,
			"application_fee_amount": 0,
			"tip_amount": 0,
			"payarc_fees": 0,
			"type": "Sale",
			"customer_email": "eric.yulfo@payarc.com",
			"net_amount": 0,
			"captured": 1,
			"is_refunded": 1,
			"status": "refunded",
			"auth_code": "TAS843",
			"statement_description": "Eric's Energy Drinks",
			"refund_reason": "requested_by_customer",
			"refund_description": "Product return",
			"created_at": 1725461676,
			"updated_at": 1725461750
		},
		"meta": {
			"include": [
				"review"
			],
			"custom": []
		}
	}`

	res := &ResponseCharge{}
	if err := json.Unmarshal([]byte(payload), res); err != nil {
		t.Fatalf("unexpected error unmarshaling refund charge response: %v", err)
	}

	if res.Data.ID != "DLbnOBbBDWMRROoM" {
		t.Errorf("expected charge ID DLbnOBbBDWMRROoM, got %s", res.Data.ID)
	}
	if res.Data.AmountRefunded != 99 {
		t.Errorf("expected amount_refunded 99, got %d", res.Data.AmountRefunded)
	}
	if !res.Data.IsRefunded.AsBool() {
		t.Errorf("expected is_refunded to be true")
	}
}
