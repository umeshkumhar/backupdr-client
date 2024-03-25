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

// checks if the ComponentDetailRest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentDetailRest{}

// ComponentDetailRest struct for ComponentDetailRest
type ComponentDetailRest struct {
	Component *string `json:"component,omitempty"`
	Summary *string `json:"summary,omitempty"`
	// Install time in UNIX Epoch time (microseconds since Jan 1, 1970)
	Installed *int64 `json:"installed,omitempty"`
}

// NewComponentDetailRest instantiates a new ComponentDetailRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentDetailRest() *ComponentDetailRest {
	this := ComponentDetailRest{}
	return &this
}

// NewComponentDetailRestWithDefaults instantiates a new ComponentDetailRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentDetailRestWithDefaults() *ComponentDetailRest {
	this := ComponentDetailRest{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ComponentDetailRest) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentDetailRest) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ComponentDetailRest) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ComponentDetailRest) SetComponent(v string) {
	o.Component = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ComponentDetailRest) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentDetailRest) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ComponentDetailRest) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *ComponentDetailRest) SetSummary(v string) {
	o.Summary = &v
}

// GetInstalled returns the Installed field value if set, zero value otherwise.
func (o *ComponentDetailRest) GetInstalled() int64 {
	if o == nil || IsNil(o.Installed) {
		var ret int64
		return ret
	}
	return *o.Installed
}

// GetInstalledOk returns a tuple with the Installed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentDetailRest) GetInstalledOk() (*int64, bool) {
	if o == nil || IsNil(o.Installed) {
		return nil, false
	}
	return o.Installed, true
}

// HasInstalled returns a boolean if a field has been set.
func (o *ComponentDetailRest) HasInstalled() bool {
	if o != nil && !IsNil(o.Installed) {
		return true
	}

	return false
}

// SetInstalled gets a reference to the given int64 and assigns it to the Installed field.
func (o *ComponentDetailRest) SetInstalled(v int64) {
	o.Installed = &v
}

func (o ComponentDetailRest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentDetailRest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Installed) {
		toSerialize["installed"] = o.Installed
	}
	return toSerialize, nil
}

type NullableComponentDetailRest struct {
	value *ComponentDetailRest
	isSet bool
}

func (v NullableComponentDetailRest) Get() *ComponentDetailRest {
	return v.value
}

func (v *NullableComponentDetailRest) Set(val *ComponentDetailRest) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentDetailRest) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentDetailRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentDetailRest(val *ComponentDetailRest) *NullableComponentDetailRest {
	return &NullableComponentDetailRest{value: val, isSet: true}
}

func (v NullableComponentDetailRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentDetailRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


