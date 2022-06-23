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
	resp, err := client.VerifyWebhookSignature(context.Background(), &VerifyWebhookSignatureRequest{
		AuthAlgo:         "Sha256WithRsa",
		CertURL:          "https://api.sandbox.paypal.com/v1/notifications/certs/CERT-360caa42-fca2a594-5a29e601",
		TransmissionID:   "4e5294a0-f1cc-11ec-9abc-d537a3be4540",
		TransmissionSig:  "Z7gYv3T27fO6IpkRTPlebXApj7S1gdcublLeg/9+J7rRpdWb97xaXJw2tzwWGlkK0RMvoxXHk3CDumTjZIJTia8OsFYFvi0FE3cpL7a1L6bLtjcUGXmVEdf17ID/Nt3cDAI3cEty0PMFzdU+KlbJYsK5n2QK47VDSr3rcI7GiwI7zxtpGN3ancub9+vfCLEIvNyVUChv3fLSSAViOrO3GZYvodW/TQU/FTJQX6wpIJ62syXzbXz8cpfAwKTI9z2Iyoj+eOcHos4aOmNiB2ZQqdlW7koDJrbqXbv+kucVmn36nYhO2GD7jZLqt4jDtKpqiuajU3Ff6V8OKfDqbQswWw==",
		TransmissionTime: "2022-06-22T01:40:33Z",
		WebhookID:        "abc123",
		Event:            []byte("{}"),
	})
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	t.Log(resp.VerificationStatus)
}
