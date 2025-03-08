package outputs

import (
	"encoding/json"

	"github.com/Ricardxdev/payarc-sdk-go/pkg/extra"
)

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
	RefundAmount             int            `json:"refund_amount"`
	Currency                 extra.Currency `json:"currency"`
	Status                   RefundStatus   `json:"status"`
	Reason                   RefundReason   `json:"reason"`
	Description              string         `json:"description"`
	Email                    *string        `json:"email"`
	ReceiptPhone             *string        `json:"receipt_phone"`
	ChargeID                 string         `json:"charge_id"`
	CreatedAt                int            `json:"created_at"`
	UpdatedAt                int            `json:"updated_at"`
	DoNotSendEmailToCustomer extra.Boolean  `json:"do_not_send_email_to_customer"`
	DoNotSendSmsToCustomer   extra.Boolean  `json:"do_not_send_sms_to_customer"`
}

func (r *Refund) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == "" {
		return nil
	}
	aux := &struct {
		Object                   string         `json:"object"`
		ID                       string         `json:"id"`
		RefundAmount             int            `json:"refund_amount"`
		Currency                 extra.Currency `json:"currency"`
		Status                   RefundStatus   `json:"status"`
		Reason                   RefundReason   `json:"reason"`
		Description              string         `json:"description"`
		Email                    *string        `json:"email"`
		ReceiptPhone             *string        `json:"receipt_phone"`
		ChargeID                 string         `json:"charge_id"`
		CreatedAt                int            `json:"created_at"`
		UpdatedAt                int            `json:"updated_at"`
		DoNotSendEmailToCustomer interface{}    `json:"do_not_send_email_to_customer"`
		DoNotSendSmsToCustomer   interface{}    `json:"do_not_send_sms_to_customer"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	r.Object = aux.Object
	r.ID = aux.ID
	r.RefundAmount = aux.RefundAmount
	r.Currency = aux.Currency
	r.Status = aux.Status
	r.Reason = aux.Reason
	r.Description = aux.Description
	r.Email = aux.Email
	r.ReceiptPhone = aux.ReceiptPhone
	r.ChargeID = aux.ChargeID
	r.CreatedAt = aux.CreatedAt
	r.UpdatedAt = aux.UpdatedAt
	r.DoNotSendEmailToCustomer = ParseBoolean(aux.DoNotSendEmailToCustomer)
	r.DoNotSendSmsToCustomer = ParseBoolean(aux.DoNotSendSmsToCustomer)

	return nil
}
