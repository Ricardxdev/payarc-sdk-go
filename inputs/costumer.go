package inputs

type CreateCustomerDTO struct {
	Email       string `json:"email"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	SendEmail   string `json:"send_email_address,omitempty"`
	CCEmail     string `json:"cc_email_address,omitempty"`
	Country     string `json:"country,omitempty"`
	Address1    string `json:"address_1,omitempty"`
	Address2    string `json:"address_2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	ZIP         string `json:"zip,omitempty"`
	Phone       string `json:"phone,omitempty"`
}

// UpdateCustomerDTO represents the parameters required to update a customer's information.
type UpdateCustomerDTO struct {
	Email         string `json:"email,omitempty"`              //  - Email: The customer's email address.
	Name          string `json:"name,omitempty"`               //  - Name: The customer's full name.
	Description   string `json:"description,omitempty"`        //  - Description: Additional descriptive text about the customer.
	SendEmail     string `json:"send_email_address,omitempty"` //  - SendEmail: The primary email address to which notifications or updates are sent.
	CCEmail       string `json:"cc_email_address,omitempty"`   //  - CCEmail: An email address to receive carbon copies of notifications.
	Country       string `json:"country,omitempty"`            //  - Country: The customer's country of residence.
	Address1      string `json:"address_1,omitempty"`          //  - Address1: The first line of the customer's address.
	Address2      string `json:"address_2,omitempty"`          //  - Address2: An optional second line for additional address information.
	City          string `json:"city,omitempty"`               //  - City: The city in which the customer resides.
	State         string `json:"state,omitempty"`              //  - State: The customer's state or province.
	ZIP           int    `json:"zip,omitempty"`                //  - ZIP: The customer's postal or ZIP code.
	Phone         int    `json:"phone,omitempty"`              //  - Phone: The customer's phone number.
	TokenID       string `json:"token_id,omitempty"`           //  - TokenID: A unique identifier token associated with the customer.
	DefaultCardID string `json:"default_card_id,omitempty"`    //  - DefaultCardID: The identifier of the customer's default payment card.
}
