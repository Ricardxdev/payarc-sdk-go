package outputs

import "github.com/Ricardxdev/payarc-sdk-go/extra"

type CardSource string

var (
	CardSourceManual   CardSource = "MANUAL"
	CardSourcePhone    CardSource = "PHONE"
	CardSourceMail     CardSource = "MAIL"
	CardSourceInternet CardSource = "INTERNET"
)

func (cs CardSource) Valid() bool {
	switch cs {
	case CardSourceManual, CardSourcePhone, CardSourceMail, CardSourceInternet:
		return true
	default:
		return false
	}
}

type CardBrand string

var (
	CardBrandVisa            CardBrand = "V"
	CardBrandMastercard      CardBrand = "M"
	CardBrandDiscover        CardBrand = "R"
	CardBrandAmericanExpress CardBrand = "X"
)

type CardsResponse struct {
	Cards    []Card   `json:"data"`
	Metadata Metadata `json:"meta"`
}

type CardResponse struct {
	Data Card `json:"data"`
	Meta Meta `json:"meta"`
}

type Card struct {
	Object      string     `json:"object"`
	ID          string     `json:"id"`
	CustomerID  string     `json:"customer_id"`
	Brand       CardBrand  `json:"brand"`
	First6Digit int        `json:"first6digit"`
	Last4Digit  string     `json:"last4digit"`
	ExpMonth    string     `json:"exp_month"`
	ExpYear     string     `json:"exp_year"`
	Fingerprint string     `json:"fingerprint"`
	CardSource  CardSource `json:"card_source"`

	IsVerified extra.Boolean `json:"is_verified"`
	IsDefault  extra.Boolean `json:"is_default"`

	HolderName string  `json:"card_holder_name"`
	Address1   string  `json:"address1"`
	Address2   string  `json:"address2"`
	State      string  `json:"state"`
	City       *string `json:"city"`
	Zip        *string `json:"zip"`
	Country    *string `json:"country"`

	AvsStatus          string        `json:"avs_status"`
	CvcStatus          string        `json:"cvc_status"`
	AddressCheckPassed extra.Boolean `json:"address_check_passed"`
	ZipCheckPassed     extra.Boolean `json:"zip_check_passed"`

	CardType    string `json:"card_type"`
	BinCountry  string `json:"bin_country"`
	BankName    *any   `json:"bank_name"`
	BankWebsite *any   `json:"bank_website"`
	BankPhone   *any   `json:"bank_phone"`

	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}
