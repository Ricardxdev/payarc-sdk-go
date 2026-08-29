package payarcsdk

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ricardxdev/payarc-sdk-go/pkg/client"
	"github.com/Ricardxdev/payarc-sdk-go/pkg/inputs"
	"github.com/Ricardxdev/payarc-sdk-go/pkg/outputs"
	"github.com/Ricardxdev/payarc-sdk-go/pkg/utils"
)

type PayarcClient interface {
	GetCharges(page int64, pageLimit int64) (*outputs.ResponseCharges, error)
	GetChargesByDate(startDate, endDate int64) ([]outputs.Charge, error)
	GetCharge(chargeID string) (*outputs.ResponseCharge, error)
	CreateCharge(input inputs.ChargeInput) (*outputs.CreateChargeResponse, error)
	GetCustomer(customerId string) (*outputs.CustomerResponse, error)
	GetCustomers(page, pageLimit int) (*outputs.CustomersResponse, error)
	CreateCustomer(input inputs.CreateCustomerDTO) (*outputs.CustomerResponse, error)
	UpdateCustomer(customerId string, input inputs.UpdateCustomerDTO) (*outputs.CustomerResponse, error)
	DeleteCustomer(id string) error
	GetCards(page, pageLimit int) (*outputs.CardsResponse, error)
	GetCustomerCards(customerId string) (*outputs.CardsResponse, error)
	GetCard(cardId string) (*outputs.CardResponse, error)
	CreateCard(input inputs.CreateCardDTO) (*outputs.Card, error)
	UpdateCard(cardId string, input inputs.UpdateCardDTO) (*outputs.CardResponse, error)
	DeleteCard(customerId, cardId string) error
	SetDefaultCard(customerId, cardId string) (*outputs.CustomerResponse, error)
	CreateToken(input inputs.CreateTokenDTO) (*outputs.TokenResponse, error)
	CreateApplePayToken(token string) (*outputs.TokenResponse, error)
	CreateCardFromToken(input outputs.Token) (*outputs.Card, error)
	GetPlans(page, pageLimit int) (*outputs.PlansResponse, error)
	GetPlan(planId string) (*outputs.PlanResponse, error)
	CreatePlan(input inputs.CreatePlanDTO) (*outputs.CreatePlanResponse, error)
	UpdatePlan(planId string, input inputs.UpdatePlanDTO) (*outputs.PlanResponse, error)
	DeletePlan(planId string) error
	ExportPlans(input inputs.ExportPlansInput) ([]byte, error)
	GetCoupons(page, pageLimit int) (*outputs.CouponsResponse, error)
	GetCoupon(couponId string) (*outputs.CouponResponse, error)
	CreateCoupon(input inputs.CreateCouponDTO) (*outputs.CouponResponse, error)
	DeleteCoupon(couponId string) error
	GetSubscriptions(input inputs.ListSubscriptionsInput) (*outputs.SubscriptionsResponse, error)
	CreateSubscription(input inputs.CreateSubscriptionDTO) (*outputs.SubscriptionResponse, error)
	UpdateSubscription(subscriptionId string, input inputs.UpdateSubscriptionDTO) (*outputs.SubscriptionResponse, error)
	CancelSubscription(subscriptionId string) (*outputs.SubscriptionResponse, error)
	PauseSubscription(subscriptionId string, input inputs.PauseSubscriptionDTO) (*outputs.SubscriptionResponse, error)
	ResumeSubscription(subscriptionId string) (*outputs.SubscriptionResponse, error)
	DeleteSubscription(subscriptionId string) error
	ExportSubscriptions(input inputs.ExportSubscriptionsInput) ([]byte, error)
}

type PayarcClientImpl struct {
	ctx                 context.Context
	client              *client.Client
	apiBaseUrl          string
	prefix              string
	chargesPath         string
	customersPath       string
	cardsPath           string
	tokensPath          string
	plansPath           string
	discountsPath       string
	subscriptionsPath   string
	dashboardExportPath string
}

type PayarcClientOptions struct {
	BaseUrl      string
	Version      string
	ApiVersion   string
	PayarcPrefix string
	Token        string
	HTTPClient   client.HTTPClient
}

func NewPayarcClient(ctx context.Context, options PayarcClientOptions) PayarcClient {
	return &PayarcClientImpl{
		ctx:                 ctx,
		client:              NewClient(options),
		prefix:              options.PayarcPrefix,
		chargesPath:         "charges",
		customersPath:       "customers",
		cardsPath:           "cards",
		tokensPath:          "tokens",
		plansPath:           "plans",
		discountsPath:       "discounts",
		subscriptionsPath:   "subscriptions",
		dashboardExportPath: "dashboard-export",
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
		HTTPClient: options.HTTPClient,
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

func (s *PayarcClientImpl) GetChargesByDate(startDate, endDate int64) ([]outputs.Charge, error) {
	var err error
	var charges *outputs.ResponseCharges
	if charges, err = s.GetCharges(1, 1000000000); err != nil {
		return nil, err
	}

	filteredCharges := make([]outputs.Charge, 0)
	for _, charge := range charges.Data {
		if charge.CreatedAt >= startDate && charge.CreatedAt <= endDate {
			filteredCharges = append(filteredCharges, charge)
		}
	}
	return filteredCharges, nil
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

func (p PayarcClientImpl) UpdateCustomer(customerId string, input inputs.UpdateCustomerDTO) (*outputs.CustomerResponse, error) {
	path := fmt.Sprintf("customers/%s", customerId)
	response := &outputs.CustomerResponse{}
	err := p.client.PatchJSON(path, input, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *PayarcClientImpl) DeleteCustomer(id string) error {
	path := fmt.Sprintf("customers/%s", id)
	err := p.client.DeleteJSON(path, nil, nil)
	if err != nil {
		return err
	}

	return err
}

func (p *PayarcClientImpl) GetCard(cardId string) (*outputs.CardResponse, error) {
	path := fmt.Sprintf("%s/%s", p.cardsPath, cardId)
	res := &outputs.CardResponse{}
	if err := p.client.Get(path, nil, res, nil); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) GetCustomerCards(customerId string) (*outputs.CardsResponse, error) {
	customer, err := p.GetCustomer(customerId)
	if err != nil {
		return nil, err
	}

	res := &outputs.CardsResponse{}
	res.Cards = customer.Data.Card.Data

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

func (p *PayarcClientImpl) UpdateCard(cardId string, input inputs.UpdateCardDTO) (*outputs.CardResponse, error) {
	path := fmt.Sprintf("cards/%s", cardId)
	response := &outputs.CardResponse{}
	err := p.client.PatchJSON(path, input, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *PayarcClientImpl) DeleteCard(customerId, cardId string) error {
	path := fmt.Sprintf("customers/%s/cards/%s", customerId, cardId)
	err := p.client.DeleteJSON(path, nil, nil)
	if err != nil {
		return err
	}
	return err
}

func (p *PayarcClientImpl) SetDefaultCard(customerId, cardId string) (*outputs.CustomerResponse, error) {
	path := fmt.Sprintf("%s/%s", p.customersPath, customerId)
	payload := &inputs.UpdateCustomerDTO{
		DefaultCardID: cardId,
	}
	response := &outputs.CustomerResponse{}
	err := p.client.PatchJSON(path, payload, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *PayarcClientImpl) CreateToken(input inputs.CreateTokenDTO) (*outputs.TokenResponse, error) {
	response := &outputs.TokenResponse{}
	err := p.client.PostJSON(p.tokensPath, input, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *PayarcClientImpl) CreateApplePayToken(token string) (*outputs.TokenResponse, error) {
	response := &outputs.TokenResponse{}
	err := p.client.PostJSON(fmt.Sprintf("%s/%s", p.tokensPath, "apple-pay"), map[string]string{"apple_encrypted_data": token}, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *PayarcClientImpl) CreateCardFromToken(input outputs.Token) (*outputs.Card, error) {
	// Create card from token to customer
	response := &outputs.CustomerResponse{}
	path := fmt.Sprintf("%s/%s", p.customersPath, *input.Card.Data.CustomerID)
	err := p.client.PatchJSON(path, struct {
		TokenID string `json:"token_id"`
	}{
		TokenID: input.ID,
	}, response)
	if err != nil {
		return nil, err
	}

	// Extract Card from Customer
	card := &outputs.Card{}
	for _, c := range response.Data.Card.Data {
		if c.Last4Digit == input.Card.Data.Last4Digit {
			*card = c
			break
		}
	}
	return card, nil
}

func (p *PayarcClientImpl) GetPlans(page, pageLimit int) (*outputs.PlansResponse, error) {
	body := map[string]interface{}{
		"page":  page,
		"limit": pageLimit,
	}

	res := &outputs.PlansResponse{}
	if err := p.client.Get(p.plansPath, nil, res, body); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) GetPlan(planId string) (*outputs.PlanResponse, error) {
	path := fmt.Sprintf("%s/%s", p.plansPath, planId)
	res := &outputs.PlanResponse{}
	if err := p.client.Get(path, nil, res, nil); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) CreatePlan(input inputs.CreatePlanDTO) (*outputs.CreatePlanResponse, error) {
	res := &outputs.CreatePlanResponse{}
	if err := p.client.PostJSON(p.plansPath, input, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) UpdatePlan(planId string, input inputs.UpdatePlanDTO) (*outputs.PlanResponse, error) {
	path := fmt.Sprintf("%s/%s", p.plansPath, planId)
	res := &outputs.PlanResponse{}
	if err := p.client.PatchJSON(path, input, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) DeletePlan(planId string) error {
	path := fmt.Sprintf("%s/%s", p.plansPath, planId)
	return p.client.DeleteJSON(path, nil, nil)
}

// ExportPlans exports plans to an Excel file. The returned bytes contain the
// file contents, which can be saved directly to disk.
func (p *PayarcClientImpl) ExportPlans(input inputs.ExportPlansInput) ([]byte, error) {
	if len(input.SelectedColumns) == 0 {
		return nil, errors.New("payarc: selected_columns is required")
	}

	path := fmt.Sprintf("%s/plan", p.dashboardExportPath)
	return p.client.GetRaw(path, utils.StructToQuery(input))
}

func (p *PayarcClientImpl) GetCoupons(page, pageLimit int) (*outputs.CouponsResponse, error) {
	body := map[string]interface{}{
		"page":  page,
		"limit": pageLimit,
	}

	res := &outputs.CouponsResponse{}
	if err := p.client.Get(p.discountsPath, nil, res, body); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) GetCoupon(couponId string) (*outputs.CouponResponse, error) {
	path := fmt.Sprintf("%s/%s", p.discountsPath, couponId)
	res := &outputs.CouponResponse{}
	if err := p.client.Get(path, nil, res, nil); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) CreateCoupon(input inputs.CreateCouponDTO) (*outputs.CouponResponse, error) {
	res := &outputs.CouponResponse{}
	if err := p.client.PostJSON(p.discountsPath, input, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) DeleteCoupon(couponId string) error {
	path := fmt.Sprintf("%s/%s", p.discountsPath, couponId)
	return p.client.DeleteJSON(path, nil, nil)
}

func (p *PayarcClientImpl) GetSubscriptions(input inputs.ListSubscriptionsInput) (*outputs.SubscriptionsResponse, error) {
	res := &outputs.SubscriptionsResponse{}
	if err := p.client.Get(p.subscriptionsPath, utils.StructToQuery(input), res, nil); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) CreateSubscription(input inputs.CreateSubscriptionDTO) (*outputs.SubscriptionResponse, error) {
	res := &outputs.SubscriptionResponse{}
	if err := p.client.PostJSON(p.subscriptionsPath, input, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) UpdateSubscription(subscriptionId string, input inputs.UpdateSubscriptionDTO) (*outputs.SubscriptionResponse, error) {
	path := fmt.Sprintf("%s/%s", p.subscriptionsPath, subscriptionId)
	res := &outputs.SubscriptionResponse{}
	if err := p.client.PatchJSON(path, input, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) CancelSubscription(subscriptionId string) (*outputs.SubscriptionResponse, error) {
	path := fmt.Sprintf("%s/%s/cancel", p.subscriptionsPath, subscriptionId)
	res := &outputs.SubscriptionResponse{}
	if err := p.client.PatchJSON(path, struct{}{}, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) PauseSubscription(subscriptionId string, input inputs.PauseSubscriptionDTO) (*outputs.SubscriptionResponse, error) {
	path := fmt.Sprintf("%s/%s/pause", p.subscriptionsPath, subscriptionId)
	res := &outputs.SubscriptionResponse{}
	if err := p.client.PostJSON(path, input, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) ResumeSubscription(subscriptionId string) (*outputs.SubscriptionResponse, error) {
	path := fmt.Sprintf("%s/%s/resume", p.subscriptionsPath, subscriptionId)
	res := &outputs.SubscriptionResponse{}
	if err := p.client.PostJSON(path, struct{}{}, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PayarcClientImpl) DeleteSubscription(subscriptionId string) error {
	path := fmt.Sprintf("%s/%s", p.subscriptionsPath, subscriptionId)
	return p.client.DeleteJSON(path, nil, nil)
}

// ExportSubscriptions exports subscriptions to an Excel file. The returned
// bytes contain the file contents, which can be saved directly to disk.
func (p *PayarcClientImpl) ExportSubscriptions(input inputs.ExportSubscriptionsInput) ([]byte, error) {
	if len(input.SelectedColumns) == 0 {
		return nil, errors.New("payarc: selected_columns is required")
	}

	path := fmt.Sprintf("%s/subscription", p.dashboardExportPath)
	return p.client.GetRaw(path, utils.StructToQuery(input))
}
