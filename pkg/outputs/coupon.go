package outputs

// Coupon represents a PayArc coupon (discount object).
type Coupon struct {
	Object           string   `json:"object"`
	DiscountID       string   `json:"discount_id"`
	AmountOff        *int64   `json:"amount_off"`
	PercentOff       *float64 `json:"percent_off"`
	Duration         string   `json:"duration"`
	DurationInMonths *int64   `json:"duration_in_months"`
	MaxRedemptions   *int64   `json:"max_redemptions"`
	RedeemBy         *string  `json:"redeem_by"`
	TimesRedeemed    *int64   `json:"times_redeemed"`
	Status           string   `json:"status"`
	Name             string   `json:"name"`
	CreatedAt        int64    `json:"created_at"`
	UpdatedAt        int64    `json:"updated_at"`
	ID               string   `json:"id"`
}

type CouponResponse struct {
	Data Coupon `json:"data"`
	Meta Meta   `json:"meta"`
}

type CouponsResponse struct {
	Data     []Coupon `json:"data"`
	Metadata Metadata `json:"meta"`
}
