package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Ricardxdev/payarc-sdk-go/pkg/utils"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	BaseURL    string
	Token      string
	Version    string
	HTTPClient HTTPClient
}

type MIMEType string

const (
	MIMEJSON     MIMEType = "application/json"
	MIMEPOSTForm MIMEType = "application/x-www-form-urlencoded"
)

func (s MIMEType) String() string {
	return string(s)
}

func (c *Client) Get(path string, queryParams map[string]string, response interface{}, body interface{}) error {
	url, err := url.Parse(fmt.Sprintf("%s%s", c.BaseURL, path))
	if err != nil {
		return err
	}

	if len(queryParams) > 0 {
		q := url.Query()
		for key, value := range queryParams {
			q.Add(key, value)
		}
		url.RawQuery = q.Encode()
	}

	var jsonBody []byte
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return c.sendRequest(req, response)
}

// GetRaw performs a GET request and returns the raw response body. It is
// useful for endpoints that do not return JSON, such as the dashboard
// export endpoints which return Excel files.
func (c *Client) GetRaw(path string, queryParams map[string]string) ([]byte, error) {
	url, err := url.Parse(fmt.Sprintf("%s%s", c.BaseURL, path))
	if err != nil {
		return nil, err
	}

	if len(queryParams) > 0 {
		q := url.Query()
		for key, value := range queryParams {
			q.Add(key, value)
		}
		url.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Accept", "application/json")

	return c.doRequest(req)
}

func (c *Client) Post(path string, body *strings.Reader, response interface{}, headers map[string]string) error {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.BaseURL, path), body)
	if err != nil {
		return err
	}

	for header := range headers {
		req.Header.Set(header, headers[header])
	}

	return c.sendRequest(req, response)
}

func (c *Client) PostForm(path string, body interface{}, response interface{}) error {
	form := utils.StructToForm(body)
	encodedBody := strings.NewReader(form.Encode())

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", c.Token),
		"Accept":        MIMEJSON.String(),
		"Content-Type":  MIMEPOSTForm.String(),
	}

	return c.Post(path, encodedBody, response, headers)
}

func (c *Client) PostJSON(path string, body interface{}, response interface{}) error {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}
	encodedBody := strings.NewReader(string(bodyJson))

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", c.Token),
		"Accept":        MIMEJSON.String(),
		"Content-Type":  MIMEJSON.String(),
	}

	return c.Post(path, encodedBody, response, headers)
}

func (c *Client) Patch(path string, body *strings.Reader, response interface{}, headers map[string]string) error {
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s%s", c.BaseURL, path), body)
	if err != nil {
		return err
	}

	for header := range headers {
		req.Header.Set(header, headers[header])
	}

	return c.sendRequest(req, response)
}

func (c *Client) PatchForm(path string, body interface{}, response interface{}) error {
	form := utils.StructToForm(body)
	encodedBody := strings.NewReader(form.Encode())

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", c.Token),
		"Accept":        MIMEJSON.String(),
		"Content-Type":  MIMEPOSTForm.String(),
	}

	return c.Patch(path, encodedBody, response, headers)
}

func (c *Client) PatchJSON(path string, body interface{}, response interface{}) error {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}
	encodedBody := strings.NewReader(string(bodyJson))

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", c.Token),
		"Accept":        MIMEJSON.String(),
		"Content-Type":  MIMEJSON.String(),
	}

	return c.Patch(path, encodedBody, response, headers)
}

func (c *Client) Delete(path string, body *strings.Reader, response interface{}, headers map[string]string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", c.BaseURL, path), nil)
	if err != nil {
		return err
	}

	for header := range headers {
		req.Header.Set(header, headers[header])
	}

	return c.sendRequest(req, response)
}

func (c *Client) DeleteJSON(path string, body interface{}, response interface{}) error {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}
	encodedBody := strings.NewReader(string(bodyJson))

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", c.Token),
		"Accept":        MIMEJSON.String(),
		"Content-Type":  MIMEJSON.String(),
	}
	return c.Delete(path, encodedBody, response, headers)
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	content, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if v != nil {
		if err = json.Unmarshal(content, v); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body from %s: %v", req.URL, err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var temp map[string]interface{}
		if err = json.Unmarshal(content, &temp); err == nil {
			if message, ok := temp["message"]; ok {
				return nil, fmt.Errorf("%s", message)
			}
		}

		return nil, fmt.Errorf("unexpected response with status: %d, message: %s", resp.StatusCode, content)
	}

	return content, nil
}
