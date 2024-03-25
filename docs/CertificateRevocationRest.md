# CertificateRevocationRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**CertRevoked** | Pointer to **bool** |  | [optional] 

## Methods

### NewCertificateRevocationRest

`func NewCertificateRevocationRest() *CertificateRevocationRest`

NewCertificateRevocationRest instantiates a new CertificateRevocationRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateRevocationRestWithDefaults

`func NewCertificateRevocationRestWithDefaults() *CertificateRevocationRest`

NewCertificateRevocationRestWithDefaults instantiates a new CertificateRevocationRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *CertificateRevocationRest) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *CertificateRevocationRest) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *CertificateRevocationRest) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *CertificateRevocationRest) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CertificateRevocationRest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CertificateRevocationRest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CertificateRevocationRest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CertificateRevocationRest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCertRevoked

`func (o *CertificateRevocationRest) GetCertRevoked() bool`

GetCertRevoked returns the CertRevoked field if non-nil, zero value otherwise.

### GetCertRevokedOk

`func (o *CertificateRevocationRest) GetCertRevokedOk() (*bool, bool)`

GetCertRevokedOk returns a tuple with the CertRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertRevoked

`func (o *CertificateRevocationRest) SetCertRevoked(v bool)`

SetCertRevoked sets CertRevoked field to given value.

### HasCertRevoked

`func (o *CertificateRevocationRest) HasCertRevoked() bool`

HasCertRevoked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


