package outputs

import "github.com/Ricardxdev/payarc-sdk-go/extra"

type CreateChargeResponse struct {
	Charge ChargeResult `json:"data"`
}

type ChargeResult struct {
	Object               string `json:"object"`
	ID                   string `json:"id"`
	Type                 string `json:"type"`
	ChargeDescription    string `json:"charge_description"`
	StatementDescription string `json:"statement_description"`
	ExternalOrderID      int    `json:"external_order_id"`

	Amount         int    `json:"amount"`
	AmountApproved string `json:"amount_approved"`
	AmountCaptured int    `json:"amount_captured"`
	AmountRefunded int    `json:"amount_refunded"`
	AmountVoided   int    `json:"amount_voided"`

	ApplicationFeeAmount int `json:"application_fee_amount"`
	TipAmount            int `json:"tip_amount"`
	PayArcFees           int `json:"payarc_fees"`
	NetAmount            int `json:"net_amount"`
	Surcharge            int `json:"surcharge"`

	Captured    extra.Boolean      `json:"captured"`
	IsRefunded  extra.Boolean      `json:"is_refunded"`
	Status      extra.ChargeStatus `json:"status"`
	UnderReview bool               `json:"under_review"`

	CardLevel      extra.ChargeCardLevel `json:"card_level"`
	AuthCode       string                `json:"auth_code"`
	FailureCode    string                `json:"failure_code"`
	FailureMessage string                `json:"failure_message"`

	DoNotSendEmailToCustomer extra.Boolean `json:"do_not_send_email_to_customer"`
	DoNotSendSmsToCustomer   extra.Boolean `json:"do_not_send_sms_to_customer"`

	KountDetails        string `json:"kount_details"`
	KountStatus         string `json:"kount_status"`
	TsysResponseCode    string `json:"tsys_response_code"`
	HostResponseCode    string `json:"host_response_code"`
	HostResponseMessage string `json:"host_response_message"`
	HostReferenceNumber string `json:"host_reference_number"`

	CreatedBy string `json:"created_by"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`

	Card CardResponse `json:"card"`
}

type ResponseCharge struct {
	Data Charge `json:"data"`
}

type ResponseCharges struct {
	Data     []Charge `json:"data"`
	Metadata Metadata `json:"meta"`
}

type Charge struct {
	Object                      string   `json:"object"`
	ID                          string   `json:"id"`
	Amount                      int      `json:"amount"`
	AmountApproved              int      `json:"amount_approved"`
	AmountRefunded              int      `json:"amount_refunded"`
	AmountCaptured              int      `json:"amount_captured"`
	AmountVoided                int      `json:"amount_voided"`
	ApplicationFeeAmount        int      `json:"application_fee_amount"`
	TipAmount                   int      `json:"tip_amount"`
	PayarcFees                  int      `json:"payarc_fees"`
	Type                        string   `json:"type"`
	CustomerEmail               string   `json:"customer_email"`
	NetAmount                   int      `json:"net_amount"`
	Captured                    int      `json:"captured"`
	IsRefunded                  int      `json:"is_refunded"`
	Status                      string   `json:"status"`
	AuthCode                    string   `json:"auth_code"`
	FailureCode                 *string  `json:"failure_code"`
	FailureMessage              *string  `json:"failure_message"`
	ChargeDescription           *string  `json:"charge_description"`
	KountDetails                string   `json:"kount_details"`
	KountStatus                 string   `json:"kount_status"`
	StatementDescription        string   `json:"statement_description"`
	Invoice                     *string  `json:"invoice"`
	UnderReview                 int      `json:"under_review"`
	CreatedAt                   int64    `json:"created_at"`
	UpdatedAt                   int64    `json:"updated_at"`
	Email                       *string  `json:"email"`
	PhoneNumber                 *string  `json:"phone_number"`
	CardLevel                   string   `json:"card_level"`
	SalesTax                    *string  `json:"sales_tax"`
	PurchaseOrder               *string  `json:"purchase_order"`
	SupplierReferenceNumber     *string  `json:"supplier_reference_number"`
	CustomerRefID               *string  `json:"customer_ref_id"`
	ShipToZip                   *string  `json:"ship_to_zip"`
	AmexDescriptor              *string  `json:"amex_descriptor"`
	CustomerVatNumber           *string  `json:"customer_vat_number"`
	SummaryCommodityCode        *string  `json:"summary_commodity_code"`
	ShippingCharges             *string  `json:"shipping_charges"`
	DutyCharges                 *string  `json:"duty_charges"`
	ShipFromZip                 *string  `json:"ship_from_zip"`
	DestinationCountryCode      *string  `json:"destination_country_code"`
	VatInvoice                  *string  `json:"vat_invoice"`
	OrderDate                   *string  `json:"order_date"`
	TaxCategory                 *string  `json:"tax_category"`
	TaxType                     *string  `json:"tax_type"`
	TaxRate                     *string  `json:"tax_rate"`
	TaxAmount                   *string  `json:"tax_amount"`
	CreatedBy                   string   `json:"created_by"`
	TerminalRegister            *string  `json:"terminal_register"`
	AmexLevel3                  []string `json:"amex_level3"`
	TipAmountRefunded           *string  `json:"tip_amount_refunded"`
	SalesTaxRefunded            *string  `json:"sales_tax_refunded"`
	ShippingChargesRefunded     *string  `json:"shipping_charges_refunded"`
	DutyChargesRefunded         *string  `json:"duty_charges_refunded"`
	PaxReferenceNumber          *string  `json:"pax_reference_number"`
	RefundReason                *string  `json:"refund_reason"`
	RefundDescription           *string  `json:"refund_description"`
	Surcharge                   int      `json:"surcharge"`
	TollAmount                  *string  `json:"toll_amount"`
	AirportFee                  *string  `json:"airport_fee"`
	HealthCare                  *string  `json:"health_care"`
	HealthCareType              *string  `json:"health_care_type"`
	PrescriptionAmount          *string  `json:"prescription_amount"`
	VisionAmount                *string  `json:"vision_amount"`
	ClinicAmount                *string  `json:"clinic_amount"`
	DentalAmount                *string  `json:"dental_amount"`
	IndustryType                *string  `json:"industry_type"`
	VoidReason                  *string  `json:"void_reason"`
	VoidDescription             *string  `json:"void_description"`
	ServerID                    *string  `json:"server_id"`
	ExternalInvoiceID           *string  `json:"external_invoice_id"`
	ExternalOrderID             *int     `json:"external_order_id"`
	TsysResponseCode            string   `json:"tsys_response_code"`
	HostResponseCode            string   `json:"host_response_code"`
	HostResponseMessage         string   `json:"host_response_message"`
	EmvIssuerScripts            *string  `json:"emv_issuer_scripts"`
	EmvIssuerAuthenticationData *string  `json:"emv_issuer_authentication_data"`
	HostReferenceNumber         string   `json:"host_reference_number"`
	SaleTerminalID              *string  `json:"sale_terminal_id"`
	SaleMID                     *string  `json:"sale_mid"`
	EdcType                     *string  `json:"edc_type"`
	EcrReferenceNumber          *string  `json:"ecr_reference_number"`
	HostTransactionIdentifier   *any     `json:"host_transaction_identifier"`
	Refund                      struct {
		Data []Refund `json:"data"`
	} `json:"refund"`
	Card struct {
		Data Card `json:"data"`
	} `json:"card"`
}
