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

// checks if the WorkflowStatusDetailsRest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStatusDetailsRest{}

// WorkflowStatusDetailsRest struct for WorkflowStatusDetailsRest
type WorkflowStatusDetailsRest struct {
	Completed *string `json:"completed,omitempty"`
	Result *string `json:"result,omitempty"`
	Status *string `json:"status,omitempty"`
	Startdate *int64 `json:"startdate,omitempty"`
	Enddate *int64 `json:"enddate,omitempty"`
	Jobtag *string `json:"jobtag,omitempty"`
	Current *string `json:"current,omitempty"`
	Pending *string `json:"pending,omitempty"`
	// Unique ID for this object
	Id *string `json:"id,omitempty"`
	// URL to access this object
	Href *string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate *int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale *bool `json:"stale,omitempty"`
}

// NewWorkflowStatusDetailsRest instantiates a new WorkflowStatusDetailsRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStatusDetailsRest() *WorkflowStatusDetailsRest {
	this := WorkflowStatusDetailsRest{}
	return &this
}

// NewWorkflowStatusDetailsRestWithDefaults instantiates a new WorkflowStatusDetailsRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStatusDetailsRestWithDefaults() *WorkflowStatusDetailsRest {
	this := WorkflowStatusDetailsRest{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetCompleted() string {
	if o == nil || IsNil(o.Completed) {
		var ret string
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetCompletedOk() (*string, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given string and assigns it to the Completed field.
func (o *WorkflowStatusDetailsRest) SetCompleted(v string) {
	o.Completed = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *WorkflowStatusDetailsRest) SetResult(v string) {
	o.Result = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowStatusDetailsRest) SetStatus(v string) {
	o.Status = &v
}

// GetStartdate returns the Startdate field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetStartdate() int64 {
	if o == nil || IsNil(o.Startdate) {
		var ret int64
		return ret
	}
	return *o.Startdate
}

// GetStartdateOk returns a tuple with the Startdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetStartdateOk() (*int64, bool) {
	if o == nil || IsNil(o.Startdate) {
		return nil, false
	}
	return o.Startdate, true
}

// HasStartdate returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasStartdate() bool {
	if o != nil && !IsNil(o.Startdate) {
		return true
	}

	return false
}

// SetStartdate gets a reference to the given int64 and assigns it to the Startdate field.
func (o *WorkflowStatusDetailsRest) SetStartdate(v int64) {
	o.Startdate = &v
}

// GetEnddate returns the Enddate field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetEnddate() int64 {
	if o == nil || IsNil(o.Enddate) {
		var ret int64
		return ret
	}
	return *o.Enddate
}

// GetEnddateOk returns a tuple with the Enddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetEnddateOk() (*int64, bool) {
	if o == nil || IsNil(o.Enddate) {
		return nil, false
	}
	return o.Enddate, true
}

// HasEnddate returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasEnddate() bool {
	if o != nil && !IsNil(o.Enddate) {
		return true
	}

	return false
}

// SetEnddate gets a reference to the given int64 and assigns it to the Enddate field.
func (o *WorkflowStatusDetailsRest) SetEnddate(v int64) {
	o.Enddate = &v
}

// GetJobtag returns the Jobtag field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetJobtag() string {
	if o == nil || IsNil(o.Jobtag) {
		var ret string
		return ret
	}
	return *o.Jobtag
}

// GetJobtagOk returns a tuple with the Jobtag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetJobtagOk() (*string, bool) {
	if o == nil || IsNil(o.Jobtag) {
		return nil, false
	}
	return o.Jobtag, true
}

// HasJobtag returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasJobtag() bool {
	if o != nil && !IsNil(o.Jobtag) {
		return true
	}

	return false
}

// SetJobtag gets a reference to the given string and assigns it to the Jobtag field.
func (o *WorkflowStatusDetailsRest) SetJobtag(v string) {
	o.Jobtag = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetCurrent() string {
	if o == nil || IsNil(o.Current) {
		var ret string
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetCurrentOk() (*string, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given string and assigns it to the Current field.
func (o *WorkflowStatusDetailsRest) SetCurrent(v string) {
	o.Current = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetPending() string {
	if o == nil || IsNil(o.Pending) {
		var ret string
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetPendingOk() (*string, bool) {
	if o == nil || IsNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasPending() bool {
	if o != nil && !IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given string and assigns it to the Pending field.
func (o *WorkflowStatusDetailsRest) SetPending(v string) {
	o.Pending = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowStatusDetailsRest) SetId(v string) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *WorkflowStatusDetailsRest) SetHref(v string) {
	o.Href = &v
}

// GetSyncdate returns the Syncdate field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetSyncdate() int64 {
	if o == nil || IsNil(o.Syncdate) {
		var ret int64
		return ret
	}
	return *o.Syncdate
}

// GetSyncdateOk returns a tuple with the Syncdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetSyncdateOk() (*int64, bool) {
	if o == nil || IsNil(o.Syncdate) {
		return nil, false
	}
	return o.Syncdate, true
}

// HasSyncdate returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasSyncdate() bool {
	if o != nil && !IsNil(o.Syncdate) {
		return true
	}

	return false
}

// SetSyncdate gets a reference to the given int64 and assigns it to the Syncdate field.
func (o *WorkflowStatusDetailsRest) SetSyncdate(v int64) {
	o.Syncdate = &v
}

// GetStale returns the Stale field value if set, zero value otherwise.
func (o *WorkflowStatusDetailsRest) GetStale() bool {
	if o == nil || IsNil(o.Stale) {
		var ret bool
		return ret
	}
	return *o.Stale
}

// GetStaleOk returns a tuple with the Stale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStatusDetailsRest) GetStaleOk() (*bool, bool) {
	if o == nil || IsNil(o.Stale) {
		return nil, false
	}
	return o.Stale, true
}

// HasStale returns a boolean if a field has been set.
func (o *WorkflowStatusDetailsRest) HasStale() bool {
	if o != nil && !IsNil(o.Stale) {
		return true
	}

	return false
}

// SetStale gets a reference to the given bool and assigns it to the Stale field.
func (o *WorkflowStatusDetailsRest) SetStale(v bool) {
	o.Stale = &v
}

func (o WorkflowStatusDetailsRest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStatusDetailsRest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Startdate) {
		toSerialize["startdate"] = o.Startdate
	}
	if !IsNil(o.Enddate) {
		toSerialize["enddate"] = o.Enddate
	}
	if !IsNil(o.Jobtag) {
		toSerialize["jobtag"] = o.Jobtag
	}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	if !IsNil(o.Pending) {
		toSerialize["pending"] = o.Pending
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

type NullableWorkflowStatusDetailsRest struct {
	value *WorkflowStatusDetailsRest
	isSet bool
}

func (v NullableWorkflowStatusDetailsRest) Get() *WorkflowStatusDetailsRest {
	return v.value
}

func (v *NullableWorkflowStatusDetailsRest) Set(val *WorkflowStatusDetailsRest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStatusDetailsRest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStatusDetailsRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStatusDetailsRest(val *WorkflowStatusDetailsRest) *NullableWorkflowStatusDetailsRest {
	return &NullableWorkflowStatusDetailsRest{value: val, isSet: true}
}

func (v NullableWorkflowStatusDetailsRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStatusDetailsRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


