/*
Backup and DR Service: Management Console API Spec

Testing SLATemplateAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/umeshkumhar/backupdr-client"
	"testing"
)

func Test_openapi_SLATemplateAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SLATemplateAPIService CloneTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId int64

		resp, httpRes, err := apiClient.SLATemplateAPI.CloneTemplates(context.Background(), sltId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService CountSlts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SLATemplateAPI.CountSlts(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService CreateOptionForPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string
		var policyId string

		resp, httpRes, err := apiClient.SLATemplateAPI.CreateOptionForPolicy(context.Background(), sltId, policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService CreatePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string

		resp, httpRes, err := apiClient.SLATemplateAPI.CreatePolicy(context.Background(), sltId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService CreateSlt", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SLATemplateAPI.CreateSlt(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService DeleteOptionForPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string
		var policyId string
		var optionId string

		httpRes, err := apiClient.SLATemplateAPI.DeleteOptionForPolicy(context.Background(), sltId, policyId, optionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService DeletePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string
		var policyId string

		httpRes, err := apiClient.SLATemplateAPI.DeletePolicy(context.Background(), sltId, policyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService DeleteSlt", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string

		httpRes, err := apiClient.SLATemplateAPI.DeleteSlt(context.Background(), sltId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService GetOptionForPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string
		var policyId string
		var optionId string

		resp, httpRes, err := apiClient.SLATemplateAPI.GetOptionForPolicy(context.Background(), sltId, policyId, optionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService GetPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string
		var policyId string

		resp, httpRes, err := apiClient.SLATemplateAPI.GetPolicy(context.Background(), sltId, policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService GetSlt", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string

		resp, httpRes, err := apiClient.SLATemplateAPI.GetSlt(context.Background(), sltId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService ListOptionForPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string
		var policyId string

		resp, httpRes, err := apiClient.SLATemplateAPI.ListOptionForPolicy(context.Background(), sltId, policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService ListPolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId int64

		resp, httpRes, err := apiClient.SLATemplateAPI.ListPolicies(context.Background(), sltId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService ListSlts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SLATemplateAPI.ListSlts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService OptionsForList15", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SLATemplateAPI.OptionsForList15(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService SettableOptionMetadataForPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string
		var policyId string

		httpRes, err := apiClient.SLATemplateAPI.SettableOptionMetadataForPolicy(context.Background(), sltId, policyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService SettableOptionMetadataForPolicyType1", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policytype string

		httpRes, err := apiClient.SLATemplateAPI.SettableOptionMetadataForPolicyType1(context.Background(), policytype).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService UpdateOptionForPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string
		var policyId string
		var optionId string

		resp, httpRes, err := apiClient.SLATemplateAPI.UpdateOptionForPolicy(context.Background(), sltId, policyId, optionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService UpdatePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string
		var policyId string

		resp, httpRes, err := apiClient.SLATemplateAPI.UpdatePolicy(context.Background(), sltId, policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SLATemplateAPIService UpdateSlt", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sltId string

		resp, httpRes, err := apiClient.SLATemplateAPI.UpdateSlt(context.Background(), sltId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
