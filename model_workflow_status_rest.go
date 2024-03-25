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

// checks if the WorkflowStatusRest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStatusRest{}

// WorkflowStatusRest struct for WorkflowStatusRest
type WorkflowStatusRest struct {
	Current *WorkflowStatusDetailsRest `json:"current,omitempty"`
	Next *WorkflowStatusDetailsRest `json:"next,omitempty"`
	Prev *WorkflowStatusDetailsRest `json:"prev,omitempty"`
	// Unique ID for this object
	Id *string `json:"id,omitempty"`
	// URL to access this object
	Href *string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate *int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale *bool `json:"stale,omitempty"`
}

// NewWorkflowStatusRest instantiates a new WorkflowStatusRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStatusRest() *WorkflowStatusRest {
	this := WorkflowStatusRest{}
	return &this
}

// NewWorkflowStatusRestWithDefaults instantiates a new WorkflowStatusRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStatusRestWithDefaults() *WorkflowStatusRest {
	this := WorkflowStatusRest{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *WorkflowStatusRest) GetCurrent() WorkflowStatusDetailsRest {
	if o == nil || IsNil(o.Current) {
		var ret WorkflowStatusDetailsRest
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusRest) GetCurrentOk() (*WorkflowStatusDetailsRest, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *WorkflowStatusRest) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given WorkflowStatusDetailsRest and assigns it to the Current field.
func (o *WorkflowStatusRest) SetCurrent(v WorkflowStatusDetailsRest) {
	o.Current = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *WorkflowStatusRest) GetNext() WorkflowStatusDetailsRest {
	if o == nil || IsNil(o.Next) {
		var ret WorkflowStatusDetailsRest
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusRest) GetNextOk() (*WorkflowStatusDetailsRest, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *WorkflowStatusRest) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given WorkflowStatusDetailsRest and assigns it to the Next field.
func (o *WorkflowStatusRest) SetNext(v WorkflowStatusDetailsRest) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *WorkflowStatusRest) GetPrev() WorkflowStatusDetailsRest {
	if o == nil || IsNil(o.Prev) {
		var ret WorkflowStatusDetailsRest
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusRest) GetPrevOk() (*WorkflowStatusDetailsRest, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *WorkflowStatusRest) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given WorkflowStatusDetailsRest and assigns it to the Prev field.
func (o *WorkflowStatusRest) SetPrev(v WorkflowStatusDetailsRest) {
	o.Prev = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowStatusRest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusRest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowStatusRest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowStatusRest) SetId(v string) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *WorkflowStatusRest) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusRest) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *WorkflowStatusRest) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *WorkflowStatusRest) SetHref(v string) {
	o.Href = &v
}

// GetSyncdate returns the Syncdate field value if set, zero value otherwise.
func (o *WorkflowStatusRest) GetSyncdate() int64 {
	if o == nil || IsNil(o.Syncdate) {
		var ret int64
		return ret
	}
	return *o.Syncdate
}

// GetSyncdateOk returns a tuple with the Syncdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusRest) GetSyncdateOk() (*int64, bool) {
	if o == nil || IsNil(o.Syncdate) {
		return nil, false
	}
	return o.Syncdate, true
}

// HasSyncdate returns a boolean if a field has been set.
func (o *WorkflowStatusRest) HasSyncdate() bool {
	if o != nil && !IsNil(o.Syncdate) {
		return true
	}

	return false
}

// SetSyncdate gets a reference to the given int64 and assigns it to the Syncdate field.
func (o *WorkflowStatusRest) SetSyncdate(v int64) {
	o.Syncdate = &v
}

// GetStale returns the Stale field value if set, zero value otherwise.
func (o *WorkflowStatusRest) GetStale() bool {
	if o == nil || IsNil(o.Stale) {
		var ret bool
		return ret
	}
	return *o.Stale
}

// GetStaleOk returns a tuple with the Stale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusRest) GetStaleOk() (*bool, bool) {
	if o == nil || IsNil(o.Stale) {
		return nil, false
	}
	return o.Stale, true
}

// HasStale returns a boolean if a field has been set.
func (o *WorkflowStatusRest) HasStale() bool {
	if o != nil && !IsNil(o.Stale) {
		return true
	}

	return false
}

// SetStale gets a reference to the given bool and assigns it to the Stale field.
func (o *WorkflowStatusRest) SetStale(v bool) {
	o.Stale = &v
}

func (o WorkflowStatusRest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStatusRest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Prev) {
		toSerialize["prev"] = o.Prev
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

type NullableWorkflowStatusRest struct {
	value *WorkflowStatusRest
	isSet bool
}

func (v NullableWorkflowStatusRest) Get() *WorkflowStatusRest {
	return v.value
}

func (v *NullableWorkflowStatusRest) Set(val *WorkflowStatusRest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStatusRest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStatusRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStatusRest(val *WorkflowStatusRest) *NullableWorkflowStatusRest {
	return &NullableWorkflowStatusRest{value: val, isSet: true}
}

func (v NullableWorkflowStatusRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStatusRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


