package payarcsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/Ricardxdev/payarc-sdk-go/pkg/extra"
	"github.com/Ricardxdev/payarc-sdk-go/pkg/inputs"
)

type mockHTTPClient struct {
	doFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.doFunc(req)
}

func TestCaptureCharge(t *testing.T) {
	mockClient := &mockHTTPClient{
		doFunc: func(req *http.Request) (*http.Response, error) {
			if req.Method != http.MethodPost {
				t.Errorf("expected POST method, got %s", req.Method)
			}
			if req.URL.Path != "/v1/charges/ch_test_capture/capture" {
				t.Errorf("expected path /v1/charges/ch_test_capture/capture, got %s", req.URL.Path)
			}
			if req.Header.Get("Authorization") != "Bearer test_token" {
				t.Errorf("expected Authorization header Bearer test_token, got %s", req.Header.Get("Authorization"))
			}

			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %v", err)
			}

			var reqBody map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &reqBody); err != nil {
				t.Fatalf("failed to unmarshal request body: %v", err)
			}

			if floatVal, ok := reqBody["amount"].(float64); !ok || int(floatVal) != 500 {
				t.Errorf("expected amount 500, got %v", reqBody["amount"])
			}
			if floatVal, ok := reqBody["tip_amount"].(float64); !ok || int(floatVal) != 50 {
				t.Errorf("expected tip_amount 50, got %v", reqBody["tip_amount"])
			}

			respJSON := `{
				"data": {
					"object": "Charge",
					"id": "ch_test_capture",
					"amount": 500,
					"amount_captured": 500,
					"tip_amount": 50,
					"captured": 1,
					"status": "submitted_for_settlement"
				}
			}`

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(respJSON)),
				Header:     make(http.Header),
			}, nil
		},
	}

	client := NewPayarcClient(context.Background(), PayarcClientOptions{
		BaseUrl:    "https://testapi.payarc.net",
		ApiVersion: "v1",
		Token:      "test_token",
		HTTPClient: mockClient,
	})

	amount := 500
	tip := 50
	res, err := client.CaptureCharge("ch_test_capture", inputs.CaptureChargeDTO{
		Amount:    &amount,
		TipAmount: &tip,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.Data.ID != "ch_test_capture" {
		t.Errorf("expected ID ch_test_capture, got %s", res.Data.ID)
	}
	if res.Data.Status != "submitted_for_settlement" {
		t.Errorf("expected status submitted_for_settlement, got %s", res.Data.Status)
	}
	if res.Data.AmountCaptured != 500 {
		t.Errorf("expected amount_captured 500, got %d", res.Data.AmountCaptured)
	}
	if res.Data.TipAmount != 50 {
		t.Errorf("expected tip_amount 50, got %d", res.Data.TipAmount)
	}
}

func TestVoidCharge(t *testing.T) {
	mockClient := &mockHTTPClient{
		doFunc: func(req *http.Request) (*http.Response, error) {
			if req.Method != http.MethodPost {
				t.Errorf("expected POST method, got %s", req.Method)
			}
			if req.URL.Path != "/v1/charges/ch_test_void/void" {
				t.Errorf("expected path /v1/charges/ch_test_void/void, got %s", req.URL.Path)
			}

			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %v", err)
			}

			var reqBody map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &reqBody); err != nil {
				t.Fatalf("failed to unmarshal request body: %v", err)
			}

			if reqBody["reason"] != "requested_by_customer" {
				t.Errorf("expected reason requested_by_customer, got %v", reqBody["reason"])
			}
			if reqBody["void_description"] != "Order canceled" {
				t.Errorf("expected void_description 'Order canceled', got %v", reqBody["void_description"])
			}

			respJSON := `{
				"data": {
					"object": "Charge",
					"id": "ch_test_void",
					"amount": 500,
					"amount_voided": 500,
					"status": "void",
					"void_reason": "requested_by_customer",
					"void_description": "Order canceled"
				}
			}`

			return &http.Response{
				StatusCode: http.StatusCreated,
				Body:       io.NopCloser(bytes.NewBufferString(respJSON)),
				Header:     make(http.Header),
			}, nil
		},
	}

	client := NewPayarcClient(context.Background(), PayarcClientOptions{
		BaseUrl:    "https://testapi.payarc.net",
		ApiVersion: "v1",
		Token:      "test_token",
		HTTPClient: mockClient,
	})

	res, err := client.VoidCharge("ch_test_void", inputs.VoidChargeDTO{
		Reason:          extra.VoidReasonRequestedByCustomer,
		VoidDescription: "Order canceled",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.Data.ID != "ch_test_void" {
		t.Errorf("expected ID ch_test_void, got %s", res.Data.ID)
	}
	if res.Data.Status != "void" {
		t.Errorf("expected status void, got %s", res.Data.Status)
	}
	if res.Data.AmountVoided != 500 {
		t.Errorf("expected amount_voided 500, got %d", res.Data.AmountVoided)
	}
	if res.Data.VoidReason == nil || *res.Data.VoidReason != "requested_by_customer" {
		t.Errorf("expected void_reason requested_by_customer, got %v", res.Data.VoidReason)
	}
}

func TestRefundCharge(t *testing.T) {
	mockClient := &mockHTTPClient{
		doFunc: func(req *http.Request) (*http.Response, error) {
			if req.Method != http.MethodPost {
				t.Errorf("expected POST method, got %s", req.Method)
			}
			if req.URL.Path != "/v1/charges/ch_test_refund/refunds" {
				t.Errorf("expected path /v1/charges/ch_test_refund/refunds, got %s", req.URL.Path)
			}

			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %v", err)
			}

			var reqBody map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &reqBody); err != nil {
				t.Fatalf("failed to unmarshal request body: %v", err)
			}

			if floatVal, ok := reqBody["amount"].(float64); !ok || int(floatVal) != 250 {
				t.Errorf("expected amount 250, got %v", reqBody["amount"])
			}
			if reqBody["reason"] != "duplicate" {
				t.Errorf("expected reason duplicate, got %v", reqBody["reason"])
			}

			respJSON := `{
				"data": {
					"object": "Charge",
					"id": "ch_test_refund",
					"amount": 500,
					"amount_refunded": 250,
					"status": "refunded",
					"is_refunded": 1,
					"refund_reason": "duplicate"
				}
			}`

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(respJSON)),
				Header:     make(http.Header),
			}, nil
		},
	}

	client := NewPayarcClient(context.Background(), PayarcClientOptions{
		BaseUrl:    "https://testapi.payarc.net",
		ApiVersion: "v1",
		Token:      "test_token",
		HTTPClient: mockClient,
	})

	amount := 250
	res, err := client.RefundCharge("ch_test_refund", inputs.RefundChargeDTO{
		Amount: &amount,
		Reason: extra.RefundReasonDuplicate,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.Data.ID != "ch_test_refund" {
		t.Errorf("expected ID ch_test_refund, got %s", res.Data.ID)
	}
	if res.Data.AmountRefunded != 250 {
		t.Errorf("expected amount_refunded 250, got %d", res.Data.AmountRefunded)
	}
	if !res.Data.IsRefunded.AsBool() {
		t.Errorf("expected is_refunded to be true")
	}
}
