package inputs

import "github.com/Ricardxdev/payarc-sdk-go/extra"

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
