/*
Backup and DR Service: Management Console API Spec

Testing ApplianceDelegationAPIService

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

func Test_openapi_ApplianceDelegationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplianceDelegationAPIService DelegateGetCallDownloadLog", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int64

		resp, httpRes, err := apiClient.ApplianceDelegationAPI.DelegateGetCallDownloadLog(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplianceDelegationAPIService DownloadConnector", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorname string
		var clusterId int64

		resp, httpRes, err := apiClient.ApplianceDelegationAPI.DownloadConnector(context.Background(), connectorname, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplianceDelegationAPIService DownloadOssNotice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int64

		resp, httpRes, err := apiClient.ApplianceDelegationAPI.DownloadOssNotice(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplianceDelegationAPIService UploadSoftwareUpgradeToAppliance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int64

		resp, httpRes, err := apiClient.ApplianceDelegationAPI.UploadSoftwareUpgradeToAppliance(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
