package outputs

import "github.com/Ricardxdev/payarc-sdk-go/pkg/extra"

type CreateChargeResponse struct {
	Charge ChargeResult `json:"data"`
}

type ChargeResult struct {
	Object                   string                `json:"object"`
	ID                       string                `json:"id"`
	Type                     string                `json:"type"`
	ChargeDescription        string                `json:"charge_description"`
	StatementDescription     string                `json:"statement_description"`
	ExternalOrderID          int                   `json:"external_order_id"`
	Amount                   int                   `json:"amount"`
	AmountApproved           string                `json:"amount_approved"`
	AmountCaptured           int                   `json:"amount_captured"`
	AmountRefunded           int                   `json:"amount_refunded"`
	AmountVoided             int                   `json:"amount_voided"`
	ApplicationFeeAmount     int                   `json:"application_fee_amount"`
	TipAmount                int                   `json:"tip_amount"`
	PayArcFees               int                   `json:"payarc_fees"`
	NetAmount                int                   `json:"net_amount"`
	Surcharge                int                   `json:"surcharge"`
	Captured                 int                   `json:"captured"`
	IsRefunded               int                   `json:"is_refunded"`
	Status                   extra.ChargeStatus    `json:"status"`
	UnderReview              bool                  `json:"under_review"`
	CardLevel                extra.ChargeCardLevel `json:"card_level"`
	AuthCode                 string                `json:"auth_code"`
	FailureCode              string                `json:"failure_code"`
	FailureMessage           string                `json:"failure_message"`
	DoNotSendEmailToCustomer int                   `json:"do_not_send_email_to_customer"`
	DoNotSendSmsToCustomer   int                   `json:"do_not_send_sms_to_customer"`
	KountDetails             string                `json:"kount_details"`
	KountStatus              string                `json:"kount_status"`
	TsysResponseCode         string                `json:"tsys_response_code"`
	HostResponseCode         string                `json:"host_response_code"`
	HostResponseMessage      string                `json:"host_response_message"`
	HostReferenceNumber      string                `json:"host_reference_number"`
	CreatedBy                string                `json:"created_by"`
	CreatedAt                int                   `json:"created_at"`
	UpdatedAt                int                   `json:"updated_at"`
	Card                     CardResponse          `json:"card"`
}

type ResponseCharge struct {
	Data Charge `json:"data"`
}

type ResponseCharges struct {
	Data     []Charge `json:"data"`
	Metadata Metadata `json:"meta"`
}

type Charge struct {
	Object                      string            `json:"object"`
	ID                          string            `json:"id"`
	Amount                      int               `json:"amount"`
	AmountApproved              int               `json:"amount_approved"`
	AmountRefunded              int               `json:"amount_refunded"`
	AmountCaptured              int               `json:"amount_captured"`
	AmountVoided                int               `json:"amount_voided"`
	ApplicationFeeAmount        int               `json:"application_fee_amount"`
	TipAmount                   int               `json:"tip_amount"`
	PayarcFees                  int               `json:"payarc_fees"`
	Type                        string            `json:"type"`
	CustomerEmail               *string           `json:"customer_email"`
	NetAmount                   int               `json:"net_amount"`
	Captured                    int               `json:"captured"`
	IsRefunded                  int               `json:"is_refunded"`
	Status                      string            `json:"status"`
	AuthCode                    *string           `json:"auth_code"`
	FailureCode                 *string           `json:"failure_code"`
	FailureMessage              *string           `json:"failure_message"`
	ChargeDescription           *string           `json:"charge_description"`
	KountDetails                string            `json:"kount_details"`
	KountStatus                 string            `json:"kount_status"`
	StatementDescription        string            `json:"statement_description"`
	Invoice                     *any              `json:"invoice"`
	UnderReview                 int               `json:"under_review"`
	CreatedAt                   int64             `json:"created_at"`
	UpdatedAt                   int64             `json:"updated_at"`
	Email                       *string           `json:"email"`
	PhoneNumber                 *string           `json:"phone_number"`
	CardLevel                   string            `json:"card_level"`
	SalesTax                    *any              `json:"sales_tax"`
	PurchaseOrder               *any              `json:"purchase_order"`
	SupplierReferenceNumber     *string           `json:"supplier_reference_number"`
	CustomerRefID               *string           `json:"customer_ref_id"`
	ShipToZip                   *string           `json:"ship_to_zip"`
	AmexDescriptor              *string           `json:"amex_descriptor"`
	CustomerVatNumber           *any              `json:"customer_vat_number"`
	SummaryCommodityCode        *any              `json:"summary_commodity_code"`
	ShippingCharges             *any              `json:"shipping_charges"`
	DutyCharges                 *int              `json:"duty_charges"`
	ShipFromZip                 *string           `json:"ship_from_zip"`
	DestinationCountryCode      *string           `json:"destination_country_code"`
	VatInvoice                  *any              `json:"vat_invoice"`
	OrderDate                   *string           `json:"order_date"`
	TaxCategory                 *any              `json:"tax_category"`
	TaxType                     *any              `json:"tax_type"`
	TaxRate                     *any              `json:"tax_rate"`
	TaxAmount                   *any              `json:"tax_amount"` // float64
	CreatedBy                   *string           `json:"created_by"`
	TerminalRegister            *TerminalRegister `json:"terminal_register"`
	AmexLevel3                  *any              `json:"amex_level3"`
	TipAmountRefunded           *any              `json:"tip_amount_refunded"`
	SalesTaxRefunded            *any              `json:"sales_tax_refunded"`
	ShippingChargesRefunded     *any              `json:"shipping_charges_refunded"`
	DutyChargesRefunded         *any              `json:"duty_charges_refunded"`
	PaxReferenceNumber          *any              `json:"pax_reference_number"`
	RefundReason                *string           `json:"refund_reason"`
	RefundDescription           *string           `json:"refund_description"`
	Surcharge                   int               `json:"surcharge"`
	TollAmount                  *any              `json:"toll_amount"`
	AirportFee                  *any              `json:"airport_fee"`
	HealthCare                  *float64          `json:"health_care"`
	HealthCareType              *string           `json:"health_care_type"`
	PrescriptionAmount          *any              `json:"prescription_amount"`
	VisionAmount                *any              `json:"vision_amount"`
	ClinicAmount                *any              `json:"clinic_amount"`
	DentalAmount                *any              `json:"dental_amount"`
	IndustryType                *any              `json:"industry_type"`
	VoidReason                  *string           `json:"void_reason"`
	VoidDescription             *string           `json:"void_description"`
	ServerID                    *string           `json:"server_id"`
	ExternalInvoiceID           *any              `json:"external_invoice_id"`
	ExternalOrderID             *any              `json:"external_order_id"`
	TsysResponseCode            string            `json:"tsys_response_code"`
	HostResponseCode            string            `json:"host_response_code"`
	HostResponseMessage         string            `json:"host_response_message"`
	EmvIssuerScripts            *any              `json:"emv_issuer_scripts"`
	EmvIssuerAuthenticationData *any              `json:"emv_issuer_authentication_data"`
	HostReferenceNumber         *string           `json:"host_reference_number"`
	SaleTerminalID              *string           `json:"sale_terminal_id"`
	SaleMID                     *any              `json:"sale_mid"`
	EdcType                     *any              `json:"edc_type"`
	EcrReferenceNumber          *any              `json:"ecr_reference_number"`
	HostTransactionIdentifier   *string           `json:"host_transaction_identifier"`
	Refund                      struct {
		Data []Refund `json:"data"`
	} `json:"refund"`
	Card struct {
		Data Card `json:"data"`
	} `json:"card"`
}

type TerminalRegister struct {
	Code             string  `json:"code"`
	CreatedAt        int64   `json:"created_at"` // string
	DatawireClientID *string `json:"datawire_client_id"`
	DeletedAt        *int64  `json:"deleted_at"`
	DeviceID         *string `json:"device_id"`
	IsEnabled        bool    `json:"is_enabled"`
	PosIdentifier    *string `json:"pos_identifier"`
	Terminal         string  `json:"terminal"`
	Type             string  `json:"type"`
	UpdatedAt        int64   `json:"updated_at"` // string
}
