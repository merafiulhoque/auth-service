package clients

import "github.com/resend/resend-go/v3"

func CreateResendClient(resendApiKey string) *resend.Client {
	client := resend.NewClient(resendApiKey)
	return client
}
