package paypal

import (
	"context"
	"testing"
)

func TestVerifyWebhookSignature(t *testing.T) {
	client, err := NewClient(testAppID, testAppSecret,
		WithSandbox(true))
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	client.Debug(true)
	resp, err := client.VerifyWebhookSignature(context.Background(), &VerifyWebhookSignatureRequest{})
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	t.Log(resp.VerificationStatus)
}
