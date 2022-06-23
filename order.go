package paypal

import (
	"context"
	"net/http"
)

// CreateOrderRequest create order request body
type CreateOrderRequest struct {
	Intent             Intent                 `json:"intent"`
	Payer              *CreateOrderPayer      `json:"payer,omitempty"`
	PurchaseUnits      []*PurchaseUnitRequest `json:"purchase_units"`
	ApplicationContext *ApplicationContext    `json:"application_context,omitempty"`
}

// CreateOrder  POST /v2/checkout/orders
func (c *Client) CreateOrder(ctx context.Context, requestId string, in *CreateOrderRequest) (resp *Order, err error) {
	reqBody, req, err := c.newRequest(ctx, http.MethodPost, "/v2/checkout/orders", in)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PayPal-Request-Id", requestId)
	err = c.DoWithReq(reqBody, req, &resp)
	return
}

// UpdateOrder update a order
func (c *Client) UpdateOrder(ctx context.Context, orderId string, op PatchOp, path string, value map[string]string) (resp *Order, err error) {
	type patchRequest struct {
		Op    PatchOp           `json:"op"`
		Path  string            `json:"path"`
		Value map[string]string `json:"value"`
	}
	err = c.Do(ctx, http.MethodPatch, "/v2/checkout/orders/"+orderId, []*patchRequest{
		{Op: op, Path: path, Value: value},
	}, resp)
	return
}

// ShowOrderDetail /v2/checkout/orders/{orderId}
func (c *Client) ShowOrderDetail(ctx context.Context, orderId string) (resp *Order, err error) {
	err = c.Do(ctx, http.MethodGet, "/v2/checkout/orders/"+orderId, nil, &resp)
	return
}

// AuthorizePaymentForOrder POST /v2/checkout/orders/{orderId}/authorize
func (c *Client) AuthorizePaymentForOrder(ctx context.Context, orderId string, in *AuthorizeOrderRequest) (resp *AuthorizeOrderResponse, err error) {
	err = c.Do(ctx, http.MethodPost, "/v2/checkout/orders/"+orderId+"/authorize", in, resp)
	return
}

// CapturePaymentForOrder POST /v2/checkout/orders/{orderId}/capture
func (c *Client) CapturePaymentForOrder(ctx context.Context, requestId, orderId string, in *CaptureOrderRequest) (resp *CaptureOrderResponse, err error) {
	reqBody, req, err := c.newRequest(ctx, http.MethodPost, "/v2/checkout/orders/"+orderId+"/capture", in)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PayPal-Request-Id", requestId)
	err = c.DoWithReq(reqBody, req, &resp)
	return
}

// ConfirmOrder POST /v2/checkout/orders/{orderId}/confirm-payment-source
// func (c *Client) ConfirmOrder(ctx context.Context, orderId string)
