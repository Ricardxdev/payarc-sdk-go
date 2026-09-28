package inputs

import "github.com/Ricardxdev/payarc-sdk-go/pkg/extra"

type ChargeInput struct {
	Amount                   int64          `json:"amount" form:"amount"`
	Capture                  extra.Boolean  `json:"capture" form:"capture"`
	CustomerID               string         `json:"customer_id" form:"customer_id"`
	CardID                   string         `json:"card_id,omitempty" form:"card_id"`
	ExternalOrderID          *int           `json:"external_order_id,omitempty" form:"external_order_id"`
	ChargeDescription        string         `json:"charge_description,omitempty" form:"charge_description"`
	Currency                 extra.Currency `json:"currency" form:"currency"`
	StatementDescription     *string        `json:"statement_description,omitempty" form:"statement_description"`
	DoNotSendEmailToCustomer extra.YesOrNo  `json:"do_not_send_email_to_customer" form:"do_not_send_email_to_customer,omitempty"`
	DoNotSendSmsToCustomer   extra.YesOrNo  `json:"do_not_send_sms_to_customer" form:"do_not_send_sms_to_customer,omitempty"`
}

// CaptureChargeDTO represents the parameters required to capture an existing uncaptured charge.
type CaptureChargeDTO struct {
	Amount    *int `json:"amount,omitempty" form:"amount"`         //  - Amount: A positive integer in cents representing how much to charge. Required if tip_amount is passed.
	TipAmount *int `json:"tip_amount,omitempty" form:"tip_amount"` //  - TipAmount: The tip amount of the transaction in cents.
}

type CaptureChargeInput = CaptureChargeDTO

// VoidChargeDTO represents the parameters required to void a charge.
type VoidChargeDTO struct {
	Reason          extra.VoidReason `json:"reason" form:"reason"`                               //  - Reason: Reason for voiding a transaction (requested_by_customer, fraudulent, duplicate, other).
	VoidDescription string           `json:"void_description,omitempty" form:"void_description"` //  - VoidDescription: Description for voiding a transaction (5 - 190 characters).
}

type VoidChargeInput = VoidChargeDTO

// RefundChargeDTO represents the parameters required to refund an existing charge.
type RefundChargeDTO struct {
	Amount                   *int               `json:"amount,omitempty" form:"amount"`                                                         //  - Amount: A positive integer in cents representing how much to refund.
	Reason                   extra.RefundReason `json:"reason,omitempty" form:"reason"`                                                         //  - Reason: Reason for refunding a transaction (requested_by_customer, fraudulent, duplicate, other).
	Description              string             `json:"description,omitempty" form:"description"`                                               //  - Description: Description for refunding a transaction (5 - 190 characters).
	DoNotSendEmailToCustomer extra.YesOrNo      `json:"do_not_send_email_to_customer,omitempty" form:"do_not_send_email_to_customer,omitempty"` //  - DoNotSendEmailToCustomer: Prevents sending confirmation email to customer (yes/no).
	DoNotSendSmsToCustomer   extra.YesOrNo      `json:"do_not_send_sms_to_customer,omitempty" form:"do_not_send_sms_to_customer,omitempty"`     //  - DoNotSendSmsToCustomer: Prevents sending confirmation SMS to customer (yes/no).
}

type RefundChargeInput = RefundChargeDTO
