package paypal

import (
	"context"
	"fmt"
	"net/http"
)

// RefundCapturedPayment Refunds a captured payment, by ID. For a full refund,
// include an empty payload in the JSON request body. For a partial refund, include an amount object in the JSON request body.
// POST /v2/payments/captures/{capture_id}/refund
func (c *Client) RefundCapturedPayment(ctx context.Context, requestId, captureId string, in *RefundCaptureRequest) (resp *RefundResponse, err error) {
	reqBody, req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("/v2/payments/captures/%s/refund", captureId), in)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PayPal-Request-Id", requestId)
	err = c.DoWithReq(reqBody, req, &resp)
	return
}
