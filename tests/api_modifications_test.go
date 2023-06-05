package tests

import (
	"context"
	"github.com/adyen/adyen-go-api-library/v7/src/adyen"
	"github.com/adyen/adyen-go-api-library/v7/src/checkout"
	"github.com/adyen/adyen-go-api-library/v7/src/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_API_Modifications(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
		APIKey          = os.Getenv("ADYEN_API_KEY")
	)

	client := adyen.NewClient(&common.Config{
		ApiKey:      APIKey,
		Environment: "TEST",
	})
	service := client.Checkout()

	t.Run("API Modifications - Captures", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			pspReference := "psp0001"
			req := service.ModificationsApi.CaptureAuthorisedPaymentConfig(context.Background(), pspReference)
			req = req.CreatePaymentCaptureRequest(checkout.CreatePaymentCaptureRequest{
				MerchantAccount: MerchantAccount,
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
			})
			res, httpRes, err := service.ModificationsApi.CaptureAuthorisedPayment(req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Original pspReference required for this operation")
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
	})

	t.Run("API Modifications - Cancels", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			req := service.ModificationsApi.CancelAuthorisedPaymentConfig(context.Background())
			req = req.CreateStandalonePaymentCancelRequest(checkout.CreateStandalonePaymentCancelRequest{
				MerchantAccount:  MerchantAccount,
				PaymentReference: "paymentReference01",
				Reference:        common.PtrString("reference01"),
			})
			res, httpRes, err := service.ModificationsApi.CancelAuthorisedPayment(req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 201, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "received", res.Status)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.ModificationsApi.CancelAuthorisedPaymentConfig(context.Background())
			req = req.CreateStandalonePaymentCancelRequest(checkout.CreateStandalonePaymentCancelRequest{
				MerchantAccount:  MerchantAccount,
				PaymentReference: "",
				Reference:        common.PtrString("reference01"),
			})
			_, httpRes, err := service.ModificationsApi.CancelAuthorisedPayment(req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Required field 'paymentReference' is not provided.")
			assert.Equal(t, 422, httpRes.StatusCode)
		})
	})

	t.Run("API Modifications - Refunds", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			pspReference := "psp0001"
			req := service.ModificationsApi.RefundCapturedPaymentConfig(context.Background(), pspReference)
			req = req.CreatePaymentRefundRequest(checkout.CreatePaymentRefundRequest{
				MerchantAccount: MerchantAccount,
				Reference:       common.PtrString("reference01"),
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
			})
			_, httpRes, err := service.ModificationsApi.RefundCapturedPayment(req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Original pspReference required")
			assert.Equal(t, 422, httpRes.StatusCode)
		})
	})

	t.Run("API Modifications - Reversals", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			pspReference := "psp0001"
			req := service.ModificationsApi.RefundOrCancelPaymentConfig(context.Background(), pspReference)
			req = req.CreatePaymentReversalRequest(checkout.CreatePaymentReversalRequest{
				MerchantAccount: MerchantAccount,
				Reference:       common.PtrString("reference01"),
			})
			_, httpRes, err := service.ModificationsApi.RefundOrCancelPayment(req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Original pspReference required")
			assert.Equal(t, 422, httpRes.StatusCode)
		})
	})
}
