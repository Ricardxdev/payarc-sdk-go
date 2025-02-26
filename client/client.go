package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Ricardxdev/payarc-sdk-go/utils"
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
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		content, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading POST request body from %s: %v", req.URL, err)
		}

		return fmt.Errorf("unexpected response with status: %d, message: %s", resp.StatusCode, content)
	}

	if v != nil {
		if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
	}
	return nil
}
