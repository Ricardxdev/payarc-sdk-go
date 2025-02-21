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
