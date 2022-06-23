package paypal

import (
	"context"
	"net/http"
)

// VerifyWebhookSignatureRequest verify webhook signature request
type VerifyWebhookSignatureRequest struct {
	AuthAlgo         string `json:"auth_algo,omitempty"`
	CertURL          string `json:"cert_url,omitempty"`
	TransmissionID   string `json:"transmission_id,omitempty"`
	TransmissionSig  string `json:"transmission_sig,omitempty"`
	TransmissionTime string `json:"transmission_time,omitempty"`
	WebhookID        string `json:"webhook_id,omitempty"`
	Event            []byte `json:"webhook_event,omitempty"`
}

// VerifyWebhookSignature POST /v1/notifications/verify-webhook-signature
func (c *Client) VerifyWebhookSignature(ctx context.Context, in *VerifyWebhookSignatureRequest) (*VerifyWebhookResponse, error) {
	var resp VerifyWebhookResponse
	if err := c.Do(ctx, http.MethodPost, "/v1/notifications/verify-webhook-signature", in, &resp); err != nil {
		return &resp, err
	}
	return &resp, nil
}
