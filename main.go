package payarcsdk

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Ricardxdev/payarc-sdk-go/client"
	"github.com/Ricardxdev/payarc-sdk-go/inputs"
	"github.com/Ricardxdev/payarc-sdk-go/outputs"
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

type PayarcClientImpl struct {
	ctx           context.Context
	client        *client.Client
	apiBaseUrl    string
	prefix        string
	chargesPath   string
	customersPath string
	cardsPath     string
	tokensPath    string
}

type PayarcClientOptions struct {
	BaseUrl      string
	Version      string
	ApiVersion   string
	PayarcPrefix string
	Token        string
}

func NewPayarcClient(ctx context.Context, options PayarcClientOptions) PayarcClient {
	return &PayarcClientImpl{
		ctx:           ctx,
		client:        NewClient(options),
		prefix:        options.PayarcPrefix,
		chargesPath:   "charges",
		customersPath: "customers",
		cardsPath:     "cards",
		tokensPath:    "tokens",
	}
}

func NewClient(options PayarcClientOptions) *client.Client {
	if options.Version == "" {
		options.Version = "1.2.0"
	}
	if options.ApiVersion == "" {
		options.ApiVersion = "v1"
	}

	options.BaseUrl += "/" + options.ApiVersion + "/"

	return &client.Client{
		BaseURL:    options.BaseUrl,
		Token:      options.Token,
		Version:    options.Version,
		HTTPClient: &http.Client{Timeout: 2 * time.Minute},
	}
}

func (p *PayarcClientImpl) GetCharge(chargeId string) (*outputs.ResponseCharge, error) {
	path := fmt.Sprintf("%s/%s", p.chargesPath, chargeId)
	response := &outputs.ResponseCharge{}
	err := p.client.Get(path, nil, response, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *PayarcClientImpl) GetCharges(page int64, pageLimit int64) (*outputs.ResponseCharges, error) {
	response := &outputs.ResponseCharges{}
	body := map[string]interface{}{
		"page":  page,
		"limit": pageLimit,
		//"search": search, <- Search permited???
	}

	if err := p.client.Get(p.chargesPath, nil, response, body); err != nil {
		return nil, err
	}

	meta := &response.Metadata.Pagination
	totalPages := meta.TotalPages
	currentPage := meta.CurrentPage

	if currentPage < totalPages {
		meta.Links.Next = fmt.Sprintf("%s/%s/%s/%d", p.apiBaseUrl, p.prefix, p.chargesPath, currentPage+1)
	}

	if currentPage > 1 {
		meta.Links.Previous = fmt.Sprintf("%s/%s/%s/%d", p.apiBaseUrl, p.prefix, p.chargesPath, currentPage-1)
	} else {
		meta.Links.Previous = ""
	}

	return response, nil
}

func (p *PayarcClientImpl) CreateCharge(input inputs.ChargeInput) (*outputs.CreateChargeResponse, error) {
	response := &outputs.CreateChargeResponse{}
	err := p.client.PostJSON(p.chargesPath, input, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *PayarcClientImpl) GetCustomer(customerId string) (*outputs.CustomerResponse, error) {
	path := fmt.Sprintf("%s/%s", p.customersPath, customerId)
	res := &outputs.CustomerResponse{}
	if err := p.client.Get(path, nil, res, nil); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) GetCustomers(page, pageLimit int) (*outputs.CustomersResponse, error) {
	body := map[string]interface{}{
		"page":  page,
		"limit": pageLimit,
		//"search": search, <- Search permited???
	}

	res := &outputs.CustomersResponse{}
	if err := p.client.Get(p.customersPath, nil, res, body); err != nil {
		return nil, err
	}

	meta := res.Metadata.Pagination
	totalPages := meta.TotalPages
	currentPage := meta.CurrentPage

	if currentPage < totalPages {
		meta.Links.Next = fmt.Sprintf("%s/%s/%s/%d", p.apiBaseUrl, p.prefix, p.customersPath, currentPage+1)
	} else {
		meta.Links.Next = ""
	}

	if currentPage > 1 {
		meta.Links.Previous = fmt.Sprintf("%s/%s/%s/%d", p.apiBaseUrl, p.prefix, p.customersPath, currentPage-1)
	} else {
		meta.Links.Previous = ""
	}
	return res, nil
}

func (p *PayarcClientImpl) CreateCustomer(input inputs.CreateCustomerDTO) (*outputs.CustomerResponse, error) {
	response := &outputs.CustomerResponse{}
	err := p.client.PostJSON(p.customersPath, input, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *PayarcClientImpl) GetCard(cardId string) (*outputs.CardResponse, error) {
	path := fmt.Sprintf("%s/%s", p.cardsPath, cardId)
	res := &outputs.CardResponse{}
	if err := p.client.Get(path, nil, res, nil); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) GetCards(page, pageLimit int) (*outputs.CardsResponse, error) {
	body := map[string]interface{}{
		"page":  page,
		"limit": pageLimit,
		//"search": search, <- Search permited???
	}

	res := &outputs.CardsResponse{}
	if err := p.client.Get(p.cardsPath, nil, res, body); err != nil {
		return nil, err
	}

	meta := res.Metadata.Pagination
	totalPages := meta.TotalPages
	currentPage := meta.CurrentPage

	if currentPage < totalPages {
		meta.Links.Next = fmt.Sprintf("%s/%s/%s/%d", p.apiBaseUrl, p.prefix, p.cardsPath, currentPage+1)
	} else {
		meta.Links.Next = ""
	}

	if currentPage > 1 {
		meta.Links.Previous = fmt.Sprintf("%s/%s/%s/%d", p.apiBaseUrl, p.prefix, p.cardsPath, currentPage-1)
	} else {
		meta.Links.Previous = ""
	}
	return res, nil
}

func (p *PayarcClientImpl) CreateCard(input inputs.CreateCardDTO) (*outputs.Card, error) {
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
	path := fmt.Sprintf("%s/%s", p.customersPath, input.CustomerID)
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

func (p *PayarcClientImpl) CreateToken(input inputs.CreateTokenDTO) (*outputs.TokenResponse, error) {
	response := &outputs.TokenResponse{}
	err := p.client.PostJSON(p.tokensPath, input, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
