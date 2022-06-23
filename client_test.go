package paypal

var testClient *Client

func newTestClient() error {
	client, err := NewClient(testAppID, testAppSecret,
		WithSandbox(true))
	if err != nil {
		return err
	}
	client.Debug(true)
	testClient = client
	return err
}
