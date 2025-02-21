package payarcsdk

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Ricardxdev/payarc-sdk-go/client"
	"github.com/Ricardxdev/payarc-sdk-go/inputs"
	"github.com/Ricardxdev/payarc-sdk-go/outputs"
)

var (
	PAYARC_TOKEN        = os.Getenv("PAYARC_TOKEN")
	PAYARC_URL          = os.Getenv("PAYARC_URL")
	BASE_DOMAIN         = os.Getenv("BASE_DOMAIN")
	BASE_CHARGES_PATH   = "payarc/charges"
	BASE_TOKENS_PATH    = "payarc/tokens"
	BASE_CUSTOMERS_PATH = "payarc/customers"
)

type PayarcClient interface {
	GetCharges(page int64, pageLimit int64) (*outputs.ResponseCharges, error)
	GetCharge(chargeID string) (*outputs.ResponseCharge, error)
	CreateCharge(input inputs.ChargeInput) (*outputs.CreateChargeResponse, error)
	GetCustomer(customerId string) (*outputs.CustomerResponse, error)
	GetCustomers(page, pageLimit int) (*outputs.CustomersResponse, error)
	CreateCustomer(input inputs.CreateCustomerDTO) (*outputs.CustomerResponse, error)
	GetCards(page, pageLimit int) (*outputs.CardsResponse, error)
	GetCard(cardId string) (*outputs.CardResponse, error)
	CreateCard(input inputs.CreateCardDTO) (*outputs.Card, error)
	CreateToken(input inputs.CreateTokenDTO) (*outputs.TokenResponse, error)
}

type payarcClientImpl struct {
	ctx    context.Context
	client *client.Client
}

func NewPayarcClient(ctx context.Context, baseUrl, version, apiVersion, token string) PayarcClient {
	return &payarcClientImpl{
		ctx:    ctx,
		client: NewClient(baseUrl, version, apiVersion, token),
	}
}

func NewClient(baseUrl, version, apiVersion, token string) *client.Client {
	if version == "" {
		version = "1.0"
	}
	if apiVersion == "" {
		apiVersion = "/v1/"
	}
	baseUrl += apiVersion

	return &client.Client{
		BaseURL:    baseUrl,
		Token:      token,
		Version:    version,
		HTTPClient: &http.Client{Timeout: 2 * time.Minute},
	}
}

func (p *payarcClientImpl) GetCharge(chargeId string) (*outputs.ResponseCharge, error) {
	path := fmt.Sprintf("charges/%s", chargeId)
	response := &outputs.ResponseCharge{}
	err := p.client.Get(path, nil, response, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *payarcClientImpl) GetCharges(page int64, pageLimit int64) (*outputs.ResponseCharges, error) {
	response := &outputs.ResponseCharges{}
	body := map[string]interface{}{
		"page":  page,
		"limit": pageLimit,
		//"search": search, <- Search permited???
	}

	if err := p.client.Get("charges", nil, response, body); err != nil {
		return nil, err
	}

	meta := &response.Metadata.Pagination
	totalPages := meta.TotalPages
	currentPage := meta.CurrentPage

	if currentPage < totalPages {
		meta.Links.Next = fmt.Sprintf("%s/%s/%d", BASE_DOMAIN, BASE_CHARGES_PATH, currentPage+1)
	}

	if currentPage > 1 {
		meta.Links.Previous = fmt.Sprintf("%s/%s/%d", BASE_DOMAIN, BASE_CHARGES_PATH, currentPage-1)
	} else {
		meta.Links.Previous = ""
	}

	return response, nil
}

func (p *payarcClientImpl) CreateCharge(input inputs.ChargeInput) (*outputs.CreateChargeResponse, error) {
	response := &outputs.CreateChargeResponse{}
	err := p.client.PostJSON("charges", input, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *payarcClientImpl) GetCustomer(customerId string) (*outputs.CustomerResponse, error) {
	path := fmt.Sprintf("customers/%s", customerId)
	res := &outputs.CustomerResponse{}
	if err := p.client.Get(path, nil, res, nil); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *payarcClientImpl) GetCustomers(page, pageLimit int) (*outputs.CustomersResponse, error) {
	body := map[string]interface{}{
		"page":  page,
		"limit": pageLimit,
		//"search": search, <- Search permited???
	}

	res := &outputs.CustomersResponse{}
	if err := p.client.Get("customers", nil, res, body); err != nil {
		return nil, err
	}

	meta := res.Metadata.Pagination
	totalPages := meta.TotalPages
	currentPage := meta.CurrentPage

	if currentPage < totalPages {
		meta.Links.Next = fmt.Sprintf("%s/%s/%d", BASE_DOMAIN, BASE_CUSTOMERS_PATH, currentPage+1)
	} else {
		meta.Links.Next = ""
	}

	if currentPage > 1 {
		meta.Links.Previous = fmt.Sprintf("%s/%s/%d", BASE_DOMAIN, BASE_CUSTOMERS_PATH, currentPage-1)
	} else {
		meta.Links.Previous = ""
	}
	return res, nil
}

func (p *payarcClientImpl) CreateCustomer(input inputs.CreateCustomerDTO) (*outputs.CustomerResponse, error) {
	response := &outputs.CustomerResponse{}
	err := p.client.PostJSON("customers", input, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *payarcClientImpl) GetCard(cardId string) (*outputs.CardResponse, error) {
	path := fmt.Sprintf("cards/%s", cardId)
	res := &outputs.CardResponse{}
	if err := p.client.Get(path, nil, res, nil); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *payarcClientImpl) GetCards(page, pageLimit int) (*outputs.CardsResponse, error) {
	body := map[string]interface{}{
		"page":  page,
		"limit": pageLimit,
		//"search": search, <- Search permited???
	}

	res := &outputs.CardsResponse{}
	if err := p.client.Get("cards", nil, res, body); err != nil {
		return nil, err
	}

	meta := res.Metadata.Pagination
	totalPages := meta.TotalPages
	currentPage := meta.CurrentPage

	if currentPage < totalPages {
		meta.Links.Next = fmt.Sprintf("%s/%s/%d", BASE_DOMAIN, BASE_CUSTOMERS_PATH, currentPage+1)
	} else {
		meta.Links.Next = ""
	}

	if currentPage > 1 {
		meta.Links.Previous = fmt.Sprintf("%s/%s/%d", BASE_DOMAIN, BASE_CUSTOMERS_PATH, currentPage-1)
	} else {
		meta.Links.Previous = ""
	}
	return res, nil
}

func (p *payarcClientImpl) CreateCard(input inputs.CreateCardDTO) (*outputs.Card, error) {
	// Create token
	tokenID := ""
	cardNumber := input.CardNumber
	if token, err := p.CreateToken(input.CreateTokenDTO); err != nil {
		return nil, err
	} else {
		tokenID = token.Data.ID
	}

	// Create card from token to customer
	response := &outputs.CustomerResponse{}
	path := fmt.Sprintf("customers/%s", input.CustomerID)
	err := p.client.PatchJSON(path, struct {
		TokenID string `json:"token_id"`
	}{
		TokenID: tokenID,
	}, response)
	if err != nil {
		return nil, err
	}

	// Extract Card from Customer
	card := &outputs.Card{}
	for _, c := range response.Data.Card.Data {
		if c.Last4Digit == cardNumber[len(cardNumber)-4:] {
			*card = c
			break
		}
	}
	return card, nil
}

func (p *payarcClientImpl) CreateToken(input inputs.CreateTokenDTO) (*outputs.TokenResponse, error) {
	response := &outputs.TokenResponse{}
	err := p.client.PostJSON("tokens", input, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
