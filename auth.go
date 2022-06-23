package paypal

import (
	"bytes"
	"context"
	"net/http"
)

// GetAccessToken POST /v1/oauth2/token
func (c *Client) GetAccessToken(ctx context.Context) (*TokenResponse, error) {
	var resp TokenResponse
	reqBody := []byte("grant_type=client_credentials")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.requestPath("/v1/oauth2/token"), bytes.NewBuffer(reqBody))
	if err != nil {
		return &resp, err
	}
	c.setDefaultHeader(req)
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.conf.appId, c.conf.secret)
	if err := c.do(reqBody, req, &resp); err != nil {
		return &resp, err
	}
	return &resp, nil
}
