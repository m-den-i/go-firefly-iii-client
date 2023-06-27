/*
Firefly III API v2.0.4

Testing TagsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_TagsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagsApiService DeleteTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tag string

		httpRes, err := apiClient.TagsApi.DeleteTag(context.Background(), tag).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GetTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tag string

		resp, httpRes, err := apiClient.TagsApi.GetTag(context.Background(), tag).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService ListAttachmentByTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tag string

		resp, httpRes, err := apiClient.TagsApi.ListAttachmentByTag(context.Background(), tag).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService ListTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsApi.ListTag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService ListTransactionByTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tag string

		resp, httpRes, err := apiClient.TagsApi.ListTransactionByTag(context.Background(), tag).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService StoreTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsApi.StoreTag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService UpdateTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tag string

		resp, httpRes, err := apiClient.TagsApi.UpdateTag(context.Background(), tag).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
