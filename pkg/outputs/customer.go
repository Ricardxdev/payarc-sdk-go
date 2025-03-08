package outputs

type Customer struct {
	Object            string  `json:"object"`              //
	CustomerID        string  `json:"customer_id"`         //
	Name              string  `json:"name"`                //
	Email             string  `json:"email"`               //
	Description       string  `json:"description"`         //
	PaymentOverdue    *int    `json:"payment_overdue"`     //
	SendEmailAddress  *string `json:"send_email_address"`  //
	CCEmailAddress    *string `json:"cc_email_address"`    //
	SourceID          *string `json:"source_id"`           //
	Address1          *string `json:"address_1"`           //
	Address2          *string `json:"address_2"`           //
	City              *string `json:"city"`                //
	State             *string `json:"state"`               //
	ZIP               *string `json:"zip"`                 //
	Phone             *string `json:"phone"`               //
	Country           *string `json:"country"`             //
	CreatedAt         int64   `json:"created_at"`          //
	UpdatedAt         int64   `json:"updated_at"`          //
	ReadableCreatedAt string  `json:"readable_created_at"` //
	ReadableUpdatedAt string  `json:"readable_updated_at"` //
	InvoicePrefix     string  `json:"invoice_prefix"`      //
	Card              struct {
		Data []Card `json:"data"`
	} `json:"card"`
	BankAccount struct {
		Data []interface{} `json:"data"`
	} `json:"bank_account"`
	Charge struct {
		Data []Charge `json:"data"`
	} `json:"charge"`
}

type Meta struct {
	Include []string `json:"include"`
	Custom  []string `json:"custom"`
}

type CustomerResponse struct {
	Data Customer `json:"data"`
	Meta Meta     `json:"meta"`
}

type CustomersResponse struct {
	Data     []Customer `json:"data"`
	Metadata Metadata   `json:"meta"`
}
