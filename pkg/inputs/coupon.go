package inputs

import "github.com/Ricardxdev/payarc-sdk-go/pkg/extra"

// CreateCouponDTO represents the parameters required to create a coupon.
// A coupon has either a PercentOff or an AmountOff and Currency. If AmountOff
// is set, that amount will be subtracted from any invoice's subtotal.
type CreateCouponDTO struct {
	Name             string               `json:"name" form:"name"`                                       //  - Name: Name of the coupon (5 - 50 characters).
	Duration         extra.CouponDuration `json:"duration" form:"duration"`                               //  - Duration: The duration for which the subscription invoices are discounted (once, repeating, forever).
	DiscountID       string               `json:"discount_id,omitempty" form:"discount_id"`               //  - DiscountID: A custom ID for the coupon. If empty, PAYARC generates one automatically.
	Currency         extra.Currency       `json:"currency,omitempty" form:"currency"`                     //  - Currency: Three-letter ISO currency code, in lowercase.
	AmountOff        *int64               `json:"amount_off,omitempty" form:"amount_off"`                 //  - AmountOff: Amount in cents subtracted from an invoice total (1 - 10000000). Required if PercentOff is not passed.
	PercentOff       *float64             `json:"percent_off,omitempty" form:"percent_off"`               //  - PercentOff: Discount percentage applied when a coupon is used (0.1 - 99.99). Required if AmountOff is not passed.
	DurationInMonths *int                 `json:"duration_in_months,omitempty" form:"duration_in_months"` //  - DurationInMonths: Duration in months of the discount (1 - 36). Required when Duration is repeating.
	MaxRedemptions   *int                 `json:"max_redemptions,omitempty" form:"max_redemptions"`       //  - MaxRedemptions: Maximum number of times a coupon can be used on subscriptions (1 - 99).
	RedeemBy         string               `json:"redeem_by,omitempty" form:"redeem_by"`                   //  - RedeemBy: Last date at which the coupon can be redeemed. Format: YYYY-MM-DD.
}
