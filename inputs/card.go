package inputs

import "github.com/Ricardxdev/payarc-sdk-go/extra"

type CreateCardDTO struct {
	CustomerID string `json:"customer_id"`
	CreateTokenDTO
}

type CreateTokenDTO struct {
	CardSource     string        `json:"card_source"`
	CardNumber     string        `json:"card_number"`
	ExpMonth       int           `json:"exp_month"`
	ExpYear        int           `json:"exp_year"`
	CVV            int           `json:"cvv"`
	CardHolderName string        `json:"card_holder_name"`
	AddressLine1   string        `json:"address_line1"`
	AddressLine2   string        `json:"address_line2"`
	City           string        `json:"city"`
	State          string        `json:"state"`
	ZIPCode        string        `json:"zip"`
	Country        string        `json:"country"`
	AuthorizeCard  extra.Boolean `json:"authorize_card"`
}

type UpdateCardDTO struct {
	ExpMonth       int    `json:"exp_month"`
	ExpYear        int    `json:"exp_year"`
	CardHolderName string `json:"name"`
	AddressLine1   string `json:"address_line1"`
	AddressLine2   string `json:"address_line2"`
	City           string `json:"city"`
	StateCode      string `json:"state_code"`
	ZIPCode        int    `json:"zip"`
	CountryCode    string `json:"country_code"`
}
