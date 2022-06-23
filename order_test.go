package paypal

import (
	"context"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	client, err := NewClient(testAppID, testAppSecret,
		WithSandbox(true),
		WithSslVerify(false))
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	client.Debug(true)
	resp, err := client.CreateOrder(context.Background(), "", &CreateOrderRequest{
		Intent: IntentCapture,
		PurchaseUnits: []*PurchaseUnitRequest{
			{Amount: &PurchaseUnitAmount{
				Currency: "USD",
				Value:    "100",
			}},
		},
	})
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	t.Logf("%+v", resp)
}
