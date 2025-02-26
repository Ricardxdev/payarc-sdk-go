package inputs

import "github.com/Ricardxdev/payarc-sdk-go/extra"

type CreateCardDTO struct {
	CustomerID string `json:"customer_id"`
	CreateTokenDTO
}

type CreateTokenDTO struct {
	CardSource     string        `json:"card_source,omitempty"`
	CardNumber     string        `json:"card_number,omitempty"`
	ExpMonth       int           `json:"exp_month,omitempty"`
	ExpYear        int           `json:"exp_year,omitempty"`
	CVV            int           `json:"cvv,omitempty"`
	CardHolderName string        `json:"card_holder_name,omitempty"`
	AddressLine1   string        `json:"address_line1,omitempty"`
	AddressLine2   string        `json:"address_line2,omitempty"`
	City           string        `json:"city,omitempty"`
	State          string        `json:"state,omitempty"`
	ZIPCode        string        `json:"zip,omitempty"`
	Country        string        `json:"country,omitempty"`
	AuthorizeCard  extra.Boolean `json:"authorize_card,omitempty"`
}

type UpdateCardDTO struct {
	ExpMonth       int    `json:"exp_month,omitempty"`
	ExpYear        int    `json:"exp_year,omitempty"`
	CardHolderName string `json:"name,omitempty"`
	AddressLine1   string `json:"address_line1,omitempty"`
	AddressLine2   string `json:"address_line2,omitempty"`
	City           string `json:"city,omitempty"`
	StateCode      string `json:"state_code,omitempty"`
	ZIPCode        int    `json:"zip,omitempty"`
	CountryCode    string `json:"country_code,omitempty"`
}
