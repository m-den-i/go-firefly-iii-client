/*
Firefly III API v2.0.4

Testing WebhooksApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/m-den-i/go-firefly-iii-client"
)

func Test_openapi_WebhooksApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WebhooksApiService DeleteWebhook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.WebhooksApi.DeleteWebhook(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService DeleteWebhookMessage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var messageId int32

		httpRes, err := apiClient.WebhooksApi.DeleteWebhookMessage(context.Background(), id, messageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService DeleteWebhookMessageAttempt", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var messageId int32
		var attemptId int32

		httpRes, err := apiClient.WebhooksApi.DeleteWebhookMessageAttempt(context.Background(), id, messageId, attemptId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetSingleWebhookMessage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var messageId int32

		resp, httpRes, err := apiClient.WebhooksApi.GetSingleWebhookMessage(context.Background(), id, messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetSingleWebhookMessageAttempt", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var messageId int32
		var attemptId int32

		resp, httpRes, err := apiClient.WebhooksApi.GetSingleWebhookMessageAttempt(context.Background(), id, messageId, attemptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetWebhook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.WebhooksApi.GetWebhook(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetWebhookMessageAttempts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var messageId int32

		resp, httpRes, err := apiClient.WebhooksApi.GetWebhookMessageAttempts(context.Background(), id, messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService GetWebhookMessages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.WebhooksApi.GetWebhookMessages(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService ListWebhook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WebhooksApi.ListWebhook(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService StoreWebhook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WebhooksApi.StoreWebhook(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService SubmitWebook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.WebhooksApi.SubmitWebook(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksApiService UpdateWebhook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.WebhooksApi.UpdateWebhook(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
