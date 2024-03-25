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

// checks if the ListClusterRest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListClusterRest{}

// ListClusterRest struct for ListClusterRest
type ListClusterRest struct {
	Count *int32 `json:"count,omitempty"`
	Items []ClusterRest `json:"items,omitempty"`
}

// NewListClusterRest instantiates a new ListClusterRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClusterRest() *ListClusterRest {
	this := ListClusterRest{}
	return &this
}

// NewListClusterRestWithDefaults instantiates a new ListClusterRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClusterRestWithDefaults() *ListClusterRest {
	this := ListClusterRest{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListClusterRest) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterRest) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListClusterRest) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ListClusterRest) SetCount(v int32) {
	o.Count = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListClusterRest) GetItems() []ClusterRest {
	if o == nil || IsNil(o.Items) {
		var ret []ClusterRest
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterRest) GetItemsOk() ([]ClusterRest, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListClusterRest) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ClusterRest and assigns it to the Items field.
func (o *ListClusterRest) SetItems(v []ClusterRest) {
	o.Items = v
}

func (o ListClusterRest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListClusterRest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableListClusterRest struct {
	value *ListClusterRest
	isSet bool
}

func (v NullableListClusterRest) Get() *ListClusterRest {
	return v.value
}

func (v *NullableListClusterRest) Set(val *ListClusterRest) {
	v.value = val
	v.isSet = true
}

func (v NullableListClusterRest) IsSet() bool {
	return v.isSet
}

func (v *NullableListClusterRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListClusterRest(val *ListClusterRest) *NullableListClusterRest {
	return &NullableListClusterRest{value: val, isSet: true}
}

func (v NullableListClusterRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListClusterRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


