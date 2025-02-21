package outputs

import "github.com/Ricardxdev/payarc-sdk-go/extra"

type RefundReason string

var (
	RefundReasonRequestedByCustomer RefundReason = "requested_by_customer"
	RefundReasonFraudulent          RefundReason = "fraudulent"
	RefundReasonDuplicate           RefundReason = "duplicate"
	RefundReasonOther               RefundReason = "other"
)

type RefundStatus string

var RefundStatusSubmitted RefundStatus = "submitted_for_refund"

type RefundResponse struct {
	Data Refund   `json:"data"`
	Meta Metadata `json:"meta"`
}

type Refund struct {
	Object                   string         `json:"object"`
	ID                       string         `json:"id"`
	RefundAmount             string         `json:"refund_amount"`
	Currency                 extra.Currency `json:"currency"`
	Status                   RefundStatus   `json:"status"`
	Reason                   RefundReason   `json:"reason"`
	Description              string         `json:"description"`
	Email                    any            `json:"email"`
	ReceiptPhone             any            `json:"receipt_phone"`
	ChargeID                 string         `json:"charge_id"`
	CreatedAt                int            `json:"created_at"`
	UpdatedAt                int            `json:"updated_at"`
	DoNotSendEmailToCustomer extra.Boolean  `json:"do_not_send_email_to_customer"`
	DoNotSendSmsToCustomer   extra.Boolean  `json:"do_not_send_sms_to_customer"`
}
