package extra

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type ChargeStatus string

var (
	ChargeStatusSubmittedForSettlement ChargeStatus = "submitted_for_settlement"
	ChargeStatusSettled                ChargeStatus = "settled"
	ChargeStatusVoid                   ChargeStatus = "void"
)

// VoidReason represents the reason for voiding a transaction.
type VoidReason string

var (
	VoidReasonRequestedByCustomer VoidReason = "requested_by_customer"
	VoidReasonFraudulent          VoidReason = "fraudulent"
	VoidReasonDuplicate           VoidReason = "duplicate"
	VoidReasonOther               VoidReason = "other"
)

// RefundReason represents the reason for refunding a transaction.
type RefundReason string

var (
	RefundReasonRequestedByCustomer RefundReason = "requested_by_customer"
	RefundReasonFraudulent          RefundReason = "fraudulent"
	RefundReasonDuplicate           RefundReason = "duplicate"
	RefundReasonOther               RefundReason = "other"
)

type Boolean uint8

var (
	False Boolean = 0
	True  Boolean = 1
)

func (b Boolean) FromBool(input bool) Boolean {
	if input {
		return True
	}
	return False
}

func (b Boolean) AsBool() bool {
	return b == True
}

type YesOrNo string

var (
	Yes YesOrNo = "yes"
	No  YesOrNo = "no"
)

func (y YesOrNo) AsBool() bool {
	return y == Yes
}

func (y YesOrNo) String() string {
	return string(y)
}

type ChargeCardLevel string

var (
	ChargeCardLevel1 ChargeCardLevel = "LEVEL1"
	ChargeCardLevel2 ChargeCardLevel = "LEVEL2"
	ChargeCardLevel3 ChargeCardLevel = "LEVEL3"
)

type Currency string

var CurrencyUSD Currency = "usd"

// FlexInt is an integer that can be unmarshaled from a JSON number or a
// numeric JSON string. The PayArc API is inconsistent with some numeric
// fields (e.g. trial_period_days, amount on plans), returning them as
// strings in some endpoints and as numbers in others.
type FlexInt int64

func (f *FlexInt) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	if str == "" || str == "null" {
		return nil
	}

	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	*f = FlexInt(n)
	return nil
}

func (f FlexInt) Int64() int64 {
	return int64(f)
}

func (f FlexInt) String() string {
	return strconv.FormatInt(int64(f), 10)
}

// PlanType represents the type of a plan.
type PlanType string

var (
	PlanTypePhysical PlanType = "physical"
	PlanTypeDigital  PlanType = "digital"
)

// PlanInterval represents the unit of interval for billing a plan.
type PlanInterval string

var (
	PlanIntervalDay   PlanInterval = "day"
	PlanIntervalMonth PlanInterval = "month"
	PlanIntervalWeek  PlanInterval = "week"
	PlanIntervalYear  PlanInterval = "year"
)

// CouponDuration represents the duration for which subscription invoices
// are discounted from the time a coupon is redeemed or applied.
type CouponDuration string

var (
	CouponDurationOnce      CouponDuration = "once"
	CouponDurationRepeating CouponDuration = "repeating"
	CouponDurationForever   CouponDuration = "forever"
)

// SubscriptionStatus represents the status of a subscription.
type SubscriptionStatus string

var (
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusTrial     SubscriptionStatus = "trial"
	SubscriptionStatusPending   SubscriptionStatus = "pending"
	SubscriptionStatusPassedDue SubscriptionStatus = "passed_due"
	SubscriptionStatusPaused    SubscriptionStatus = "paused"
	SubscriptionStatusUnpaid    SubscriptionStatus = "unpaid"
	SubscriptionStatusCanceled  SubscriptionStatus = "canceled"
)

// PauseUnit represents the unit used to pause a subscription.
type PauseUnit string

var (
	PauseUnitDays   PauseUnit = "days"
	PauseUnitWeeks  PauseUnit = "weeks"
	PauseUnitMonths PauseUnit = "months"
	PauseUnitYears  PauseUnit = "years"
)

// SortOrder represents the direction used to sort list results.
type SortOrder string

var (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

type DateTime struct {
	time.Time
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format(time.RFC3339))
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	format := time.DateOnly
	if strings.Contains(str, "T") {
		format = time.RFC3339
	} else if strings.Contains(str, " ") {
		format = time.DateTime
	}

	parsed, err := time.Parse(format, str)
	if err != nil {
		return err
	}

	d.Time = parsed

	return nil
}

func (d DateTime) String() string {
	return d.Format(time.DateTime)
}

func (d *DateTime) UnmarshalText(data []byte) error {
	str := string(data)

	format := time.DateOnly
	if strings.Contains(str, "T") {
		format = time.RFC3339
	} else if strings.Contains(str, " ") {
		format = time.DateTime
	}

	parsed, err := time.Parse(format, str)
	if err != nil {
		return err
	}

	d.Time = parsed

	return nil
}
