/*
Backup and DR Service: Management Console API Spec

This document defines the API for the Global Manager. All communication is done over HTTPS with UTF-8 encoding. JSON is the only supported format for both request and response payloads. <p></p>Please read <a href=\"https://cloud.google.com/backup-disaster-recovery/docs/api/RestAPIGeneralConcepts.pdf\">Management Console API General concept</a><p></p>To login, use the /session POST API below.<br></br>Then copy the resulting session_id from the output and click on the Authorize button on the top right. Paste the string \"Actifio \" followed by the session id into the form and click Authorize.<p></p>Login is not necessary for reading the rest of this API document. However, login will allow you to try the APIs out within this page.

API version: V11.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the VersionDetailRest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionDetailRest{}

// VersionDetailRest struct for VersionDetailRest
type VersionDetailRest struct {
	Product *string `json:"product,omitempty"`
	Components []ComponentDetailRest `json:"components,omitempty"`
	Summary *string `json:"summary,omitempty"`
	// Install time in UNIX Epoch time (microseconds since Jan 1, 1970)
	Installed *int64 `json:"installed,omitempty"`
}

// NewVersionDetailRest instantiates a new VersionDetailRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionDetailRest() *VersionDetailRest {
	this := VersionDetailRest{}
	return &this
}

// NewVersionDetailRestWithDefaults instantiates a new VersionDetailRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionDetailRestWithDefaults() *VersionDetailRest {
	this := VersionDetailRest{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *VersionDetailRest) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetailRest) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *VersionDetailRest) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *VersionDetailRest) SetProduct(v string) {
	o.Product = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *VersionDetailRest) GetComponents() []ComponentDetailRest {
	if o == nil || IsNil(o.Components) {
		var ret []ComponentDetailRest
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetailRest) GetComponentsOk() ([]ComponentDetailRest, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *VersionDetailRest) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ComponentDetailRest and assigns it to the Components field.
func (o *VersionDetailRest) SetComponents(v []ComponentDetailRest) {
	o.Components = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *VersionDetailRest) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetailRest) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *VersionDetailRest) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *VersionDetailRest) SetSummary(v string) {
	o.Summary = &v
}

// GetInstalled returns the Installed field value if set, zero value otherwise.
func (o *VersionDetailRest) GetInstalled() int64 {
	if o == nil || IsNil(o.Installed) {
		var ret int64
		return ret
	}
	return *o.Installed
}

// GetInstalledOk returns a tuple with the Installed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionDetailRest) GetInstalledOk() (*int64, bool) {
	if o == nil || IsNil(o.Installed) {
		return nil, false
	}
	return o.Installed, true
}

// HasInstalled returns a boolean if a field has been set.
func (o *VersionDetailRest) HasInstalled() bool {
	if o != nil && !IsNil(o.Installed) {
		return true
	}

	return false
}

// SetInstalled gets a reference to the given int64 and assigns it to the Installed field.
func (o *VersionDetailRest) SetInstalled(v int64) {
	o.Installed = &v
}

func (o VersionDetailRest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionDetailRest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Installed) {
		toSerialize["installed"] = o.Installed
	}
	return toSerialize, nil
}

type NullableVersionDetailRest struct {
	value *VersionDetailRest
	isSet bool
}

func (v NullableVersionDetailRest) Get() *VersionDetailRest {
	return v.value
}

func (v *NullableVersionDetailRest) Set(val *VersionDetailRest) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionDetailRest) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionDetailRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionDetailRest(val *VersionDetailRest) *NullableVersionDetailRest {
	return &NullableVersionDetailRest{value: val, isSet: true}
}

func (v NullableVersionDetailRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionDetailRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


