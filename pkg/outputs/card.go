package outputs

import (
	"fmt"

	"encoding/json"

	"github.com/Ricardxdev/payarc-sdk-go/pkg/extra"
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
	Object             string        `json:"object"`
	ID                 string        `json:"id"`
	CustomerID         *string       `json:"customer_id"`
	Brand              CardBrand     `json:"brand"`
	First6Digit        string        `json:"first6digit"`
	Last4Digit         string        `json:"last4digit"`
	ExpMonth           string        `json:"exp_month"`
	ExpYear            string        `json:"exp_year"`
	Fingerprint        string        `json:"fingerprint"`
	CardSource         CardSource    `json:"card_source"`
	IsVerified         extra.Boolean `json:"is_verified"`
	IsDefault          extra.Boolean `json:"is_default"`
	HolderName         *string       `json:"card_holder_name"`
	Address1           *string       `json:"address1,omitempty"`
	Address2           *string       `json:"address2,omitempty"`
	State              *string       `json:"state,omitempty"`
	City               *string       `json:"city,omitempty"`
	Zip                *string       `json:"zip,omitempty"`
	Country            *string       `json:"country,omitempty"`
	AvsStatus          *string       `json:"avs_status,omitempty"`
	CvcStatus          *string       `json:"cvc_status,omitempty"`
	AddressCheckPassed extra.Boolean `json:"address_check_passed"`
	ZipCheckPassed     extra.Boolean `json:"zip_check_passed"`
	CardType           string        `json:"card_type"`
	BinCountry         string        `json:"bin_country"`
	BankName           *string       `json:"bank_name,omitempty"`
	BankWebsite        *string       `json:"bank_website,omitempty"`
	BankPhone          *string       `json:"bank_phone,omitempty"`
	CreatedAt          int64         `json:"created_at"`
	UpdatedAt          int64         `json:"updated_at"`
}

func (c *Card) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Object             string        `json:"object"`
		ID                 string        `json:"id"`
		CustomerID         string        `json:"customer_id"`
		Brand              CardBrand     `json:"brand"`
		First6Digit        interface{}   `json:"first6digit"` // float64
		Last4Digit         interface{}   `json:"last4digit"`
		ExpMonth           interface{}   `json:"exp_month"`
		ExpYear            interface{}   `json:"exp_year"`
		Fingerprint        string        `json:"fingerprint"`
		CardSource         CardSource    `json:"card_source"`
		IsVerified         extra.Boolean `json:"is_verified"`
		IsDefault          extra.Boolean `json:"is_default"`
		HolderName         string        `json:"card_holder_name"`
		Address1           *string       `json:"address1"`
		Address2           *string       `json:"address2"`
		State              *string       `json:"state"`
		City               *string       `json:"city"`
		Zip                *string       `json:"zip"`
		Country            *string       `json:"country"`
		AvsStatus          string        `json:"avs_status"`
		CvcStatus          string        `json:"cvc_status"`
		AddressCheckPassed extra.Boolean `json:"address_check_passed"`
		ZipCheckPassed     extra.Boolean `json:"zip_check_passed"`
		CardType           string        `json:"card_type"`
		BinCountry         string        `json:"bin_country"`
		BankName           *string       `json:"bank_name"`
		BankWebsite        *string       `json:"bank_website"`
		BankPhone          *string       `json:"bank_phone"`
		CreatedAt          int64         `json:"created_at"`
		UpdatedAt          int64         `json:"updated_at"`
	}{}

	if err := json.Unmarshal(data, aux); err != nil {
		fmt.Println("from Payarc SDK package at Card UnmarshalJSON Custom Implementation: ", err)
		return err
	}

	c.Object = aux.Object
	c.ID = aux.ID
	c.CustomerID = &aux.CustomerID
	c.Brand = aux.Brand
	c.CardSource = aux.CardSource
	c.Fingerprint = aux.Fingerprint
	c.HolderName = &aux.HolderName
	c.Address1 = aux.Address1
	c.Address2 = aux.Address2
	c.State = aux.State
	c.City = aux.City
	c.Zip = aux.Zip
	c.Country = aux.Country
	c.AvsStatus = &aux.AvsStatus
	c.CvcStatus = &aux.CvcStatus
	c.CardType = aux.CardType
	c.BinCountry = aux.BinCountry
	c.BankName = aux.BankName
	c.BankWebsite = aux.BankWebsite
	c.BankPhone = aux.BankPhone
	c.CreatedAt = aux.CreatedAt
	c.UpdatedAt = aux.UpdatedAt

	switch aux.First6Digit.(type) {
	case string:
		c.First6Digit = aux.First6Digit.(string)
	case float64:
		c.First6Digit = fmt.Sprintf("%d", int(aux.First6Digit.(float64)))
	}

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

	return nil
}
