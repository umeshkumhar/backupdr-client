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

// checks if the RestorePreflightRest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestorePreflightRest{}

// RestorePreflightRest struct for RestorePreflightRest
type RestorePreflightRest struct {
	Status *bool `json:"status,omitempty"`
	Testlist []TestRest `json:"testlist,omitempty"`
	// Unique ID for this object
	Id *string `json:"id,omitempty"`
	// URL to access this object
	Href *string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate *int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale *bool `json:"stale,omitempty"`
}

// NewRestorePreflightRest instantiates a new RestorePreflightRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestorePreflightRest() *RestorePreflightRest {
	this := RestorePreflightRest{}
	return &this
}

// NewRestorePreflightRestWithDefaults instantiates a new RestorePreflightRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestorePreflightRestWithDefaults() *RestorePreflightRest {
	this := RestorePreflightRest{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RestorePreflightRest) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestorePreflightRest) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RestorePreflightRest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *RestorePreflightRest) SetStatus(v bool) {
	o.Status = &v
}

// GetTestlist returns the Testlist field value if set, zero value otherwise.
func (o *RestorePreflightRest) GetTestlist() []TestRest {
	if o == nil || IsNil(o.Testlist) {
		var ret []TestRest
		return ret
	}
	return o.Testlist
}

// GetTestlistOk returns a tuple with the Testlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestorePreflightRest) GetTestlistOk() ([]TestRest, bool) {
	if o == nil || IsNil(o.Testlist) {
		return nil, false
	}
	return o.Testlist, true
}

// HasTestlist returns a boolean if a field has been set.
func (o *RestorePreflightRest) HasTestlist() bool {
	if o != nil && !IsNil(o.Testlist) {
		return true
	}

	return false
}

// SetTestlist gets a reference to the given []TestRest and assigns it to the Testlist field.
func (o *RestorePreflightRest) SetTestlist(v []TestRest) {
	o.Testlist = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RestorePreflightRest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestorePreflightRest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RestorePreflightRest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RestorePreflightRest) SetId(v string) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RestorePreflightRest) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestorePreflightRest) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RestorePreflightRest) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RestorePreflightRest) SetHref(v string) {
	o.Href = &v
}

// GetSyncdate returns the Syncdate field value if set, zero value otherwise.
func (o *RestorePreflightRest) GetSyncdate() int64 {
	if o == nil || IsNil(o.Syncdate) {
		var ret int64
		return ret
	}
	return *o.Syncdate
}

// GetSyncdateOk returns a tuple with the Syncdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestorePreflightRest) GetSyncdateOk() (*int64, bool) {
	if o == nil || IsNil(o.Syncdate) {
		return nil, false
	}
	return o.Syncdate, true
}

// HasSyncdate returns a boolean if a field has been set.
func (o *RestorePreflightRest) HasSyncdate() bool {
	if o != nil && !IsNil(o.Syncdate) {
		return true
	}

	return false
}

// SetSyncdate gets a reference to the given int64 and assigns it to the Syncdate field.
func (o *RestorePreflightRest) SetSyncdate(v int64) {
	o.Syncdate = &v
}

// GetStale returns the Stale field value if set, zero value otherwise.
func (o *RestorePreflightRest) GetStale() bool {
	if o == nil || IsNil(o.Stale) {
		var ret bool
		return ret
	}
	return *o.Stale
}

// GetStaleOk returns a tuple with the Stale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestorePreflightRest) GetStaleOk() (*bool, bool) {
	if o == nil || IsNil(o.Stale) {
		return nil, false
	}
	return o.Stale, true
}

// HasStale returns a boolean if a field has been set.
func (o *RestorePreflightRest) HasStale() bool {
	if o != nil && !IsNil(o.Stale) {
		return true
	}

	return false
}

// SetStale gets a reference to the given bool and assigns it to the Stale field.
func (o *RestorePreflightRest) SetStale(v bool) {
	o.Stale = &v
}

func (o RestorePreflightRest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestorePreflightRest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Testlist) {
		toSerialize["testlist"] = o.Testlist
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Syncdate) {
		toSerialize["syncdate"] = o.Syncdate
	}
	if !IsNil(o.Stale) {
		toSerialize["stale"] = o.Stale
	}
	return toSerialize, nil
}

type NullableRestorePreflightRest struct {
	value *RestorePreflightRest
	isSet bool
}

func (v NullableRestorePreflightRest) Get() *RestorePreflightRest {
	return v.value
}

func (v *NullableRestorePreflightRest) Set(val *RestorePreflightRest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestorePreflightRest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestorePreflightRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestorePreflightRest(val *RestorePreflightRest) *NullableRestorePreflightRest {
	return &NullableRestorePreflightRest{value: val, isSet: true}
}

func (v NullableRestorePreflightRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestorePreflightRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


