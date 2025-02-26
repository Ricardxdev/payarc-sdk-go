package outputs

import (
	"fmt"

	"encoding/json"

	"github.com/Ricardxdev/payarc-sdk-go/extra"
)

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

func (c *Card) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Object      string      `json:"object"`
		ID          string      `json:"id"`
		CustomerID  string      `json:"customer_id"`
		Brand       CardBrand   `json:"brand"`
		First6Digit int         `json:"first6digit"`
		Last4Digit  interface{} `json:"last4digit"` // string - int
		ExpMonth    interface{} `json:"exp_month"`  // string - int
		ExpYear     interface{} `json:"exp_year"`   // string - int
		Fingerprint string      `json:"fingerprint"`
		CardSource  CardSource  `json:"card_source"`

		IsVerified interface{} `json:"is_verified"` // extra.Boolean - bool
		IsDefault  interface{} `json:"is_default"`  // extra.Boolean - bool

		HolderName string  `json:"card_holder_name"`
		Address1   string  `json:"address1"`
		Address2   string  `json:"address2"`
		State      string  `json:"state"`
		City       *string `json:"city"`
		Zip        *string `json:"zip"`
		Country    *string `json:"country"`

		AvsStatus          string      `json:"avs_status"`
		CvcStatus          string      `json:"cvc_status"`
		AddressCheckPassed interface{} `json:"address_check_passed"` // extra.Boolean - bool
		ZipCheckPassed     interface{} `json:"zip_check_passed"`     // extra.Boolean - bool

		CardType    string `json:"card_type"`
		BinCountry  string `json:"bin_country"`
		BankName    *any   `json:"bank_name"`
		BankWebsite *any   `json:"bank_website"`
		BankPhone   *any   `json:"bank_phone"`

		CreatedAt int `json:"created_at"`
		UpdatedAt int `json:"updated_at"`
	}{}

	if err := json.Unmarshal(data, aux); err != nil {
		fmt.Println("from Payarc SDK package at Card UnmarshalJSON Custom Implementation: ", err)
		return nil
	}

	c.Object = aux.Object
	c.ID = aux.ID
	c.CustomerID = aux.CustomerID
	c.Brand = aux.Brand
	c.First6Digit = aux.First6Digit
	c.Fingerprint = aux.Fingerprint
	c.CardSource = aux.CardSource

	c.HolderName = aux.HolderName
	c.Address1 = aux.Address1
	c.Address2 = aux.Address2
	c.State = aux.State
	c.City = aux.City
	c.Zip = aux.Zip
	c.Country = aux.Country
	c.AvsStatus = aux.AvsStatus
	c.CvcStatus = aux.CvcStatus
	c.CardType = aux.CardType
	c.BinCountry = aux.BinCountry
	c.BankName = aux.BankName
	c.BankWebsite = aux.BankWebsite
	c.BankPhone = aux.BankPhone
	c.CreatedAt = aux.CreatedAt
	c.UpdatedAt = aux.UpdatedAt

	// int - string -> string

	if last4Digit, ok := aux.Last4Digit.(string); ok {
		c.Last4Digit = last4Digit
	} else if last4Digit, ok := aux.Last4Digit.(int); ok {
		c.Last4Digit = fmt.Sprintf("%d", int(last4Digit))
	}
	if expMonth, ok := aux.ExpMonth.(string); ok {
		c.ExpMonth = expMonth
	} else if expMonth, ok := aux.ExpMonth.(int); ok {
		c.ExpMonth = fmt.Sprintf("%d", int(expMonth))
	}
	if expYear, ok := aux.ExpYear.(string); ok {
		c.ExpYear = expYear
	} else if expYear, ok := aux.ExpYear.(int); ok {
		c.ExpYear = fmt.Sprintf("%d", int(expYear))
	}

	// bool - extra.Boolean -> extra.Boolean

	dummyBool := extra.Boolean(0)
	if isVerified, ok := aux.IsVerified.(bool); ok {
		c.IsVerified = dummyBool.FromBool(isVerified)
	} else if isVerified, ok := aux.IsVerified.(extra.Boolean); ok {
		c.IsVerified = isVerified
	}
	if isDefault, ok := aux.IsDefault.(bool); ok {
		c.IsDefault = dummyBool.FromBool(isDefault)
	} else if isDefault, ok := aux.IsDefault.(extra.Boolean); ok {
		c.IsDefault = isDefault
	}

	if addressCheckPassed, ok := aux.AddressCheckPassed.(bool); ok {
		c.AddressCheckPassed = dummyBool.FromBool(addressCheckPassed)
	} else if addressCheckPassed, ok := aux.AddressCheckPassed.(extra.Boolean); ok {
		c.AddressCheckPassed = addressCheckPassed
	}

	if zipCheckPassed, ok := aux.ZipCheckPassed.(bool); ok {
		c.ZipCheckPassed = dummyBool.FromBool(zipCheckPassed)
	} else if zipCheckPassed, ok := aux.ZipCheckPassed.(extra.Boolean); ok {
		c.ZipCheckPassed = zipCheckPassed
	}

	return nil
}
