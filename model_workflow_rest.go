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

// checks if the WorkflowRest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRest{}

// WorkflowRest struct for WorkflowRest
type WorkflowRest struct {
	Name *string `json:"name,omitempty"`
	Application *ApplicationRest `json:"application,omitempty"`
	Cluster *ClusterRest `json:"cluster,omitempty"`
	Status *WorkflowStatusRest `json:"status,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	Props []KeyValueRest `json:"props,omitempty"`
	Schedule *WorkScheduleRest `json:"schedule,omitempty"`
	Items []WorkflowItemRest `json:"items,omitempty"`
	// Unique ID for this object
	Id *string `json:"id,omitempty"`
	// URL to access this object
	Href *string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate *int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale *bool `json:"stale,omitempty"`
}

// NewWorkflowRest instantiates a new WorkflowRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRest() *WorkflowRest {
	this := WorkflowRest{}
	return &this
}

// NewWorkflowRestWithDefaults instantiates a new WorkflowRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRestWithDefaults() *WorkflowRest {
	this := WorkflowRest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowRest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowRest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowRest) SetName(v string) {
	o.Name = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *WorkflowRest) GetApplication() ApplicationRest {
	if o == nil || IsNil(o.Application) {
		var ret ApplicationRest
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetApplicationOk() (*ApplicationRest, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *WorkflowRest) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given ApplicationRest and assigns it to the Application field.
func (o *WorkflowRest) SetApplication(v ApplicationRest) {
	o.Application = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *WorkflowRest) GetCluster() ClusterRest {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterRest
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetClusterOk() (*ClusterRest, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *WorkflowRest) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterRest and assigns it to the Cluster field.
func (o *WorkflowRest) SetCluster(v ClusterRest) {
	o.Cluster = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowRest) GetStatus() WorkflowStatusRest {
	if o == nil || IsNil(o.Status) {
		var ret WorkflowStatusRest
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetStatusOk() (*WorkflowStatusRest, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowRest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WorkflowStatusRest and assigns it to the Status field.
func (o *WorkflowRest) SetStatus(v WorkflowStatusRest) {
	o.Status = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *WorkflowRest) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *WorkflowRest) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *WorkflowRest) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetProps returns the Props field value if set, zero value otherwise.
func (o *WorkflowRest) GetProps() []KeyValueRest {
	if o == nil || IsNil(o.Props) {
		var ret []KeyValueRest
		return ret
	}
	return o.Props
}

// GetPropsOk returns a tuple with the Props field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetPropsOk() ([]KeyValueRest, bool) {
	if o == nil || IsNil(o.Props) {
		return nil, false
	}
	return o.Props, true
}

// HasProps returns a boolean if a field has been set.
func (o *WorkflowRest) HasProps() bool {
	if o != nil && !IsNil(o.Props) {
		return true
	}

	return false
}

// SetProps gets a reference to the given []KeyValueRest and assigns it to the Props field.
func (o *WorkflowRest) SetProps(v []KeyValueRest) {
	o.Props = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *WorkflowRest) GetSchedule() WorkScheduleRest {
	if o == nil || IsNil(o.Schedule) {
		var ret WorkScheduleRest
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetScheduleOk() (*WorkScheduleRest, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *WorkflowRest) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given WorkScheduleRest and assigns it to the Schedule field.
func (o *WorkflowRest) SetSchedule(v WorkScheduleRest) {
	o.Schedule = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *WorkflowRest) GetItems() []WorkflowItemRest {
	if o == nil || IsNil(o.Items) {
		var ret []WorkflowItemRest
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetItemsOk() ([]WorkflowItemRest, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *WorkflowRest) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []WorkflowItemRest and assigns it to the Items field.
func (o *WorkflowRest) SetItems(v []WorkflowItemRest) {
	o.Items = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowRest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowRest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowRest) SetId(v string) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *WorkflowRest) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *WorkflowRest) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *WorkflowRest) SetHref(v string) {
	o.Href = &v
}

// GetSyncdate returns the Syncdate field value if set, zero value otherwise.
func (o *WorkflowRest) GetSyncdate() int64 {
	if o == nil || IsNil(o.Syncdate) {
		var ret int64
		return ret
	}
	return *o.Syncdate
}

// GetSyncdateOk returns a tuple with the Syncdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetSyncdateOk() (*int64, bool) {
	if o == nil || IsNil(o.Syncdate) {
		return nil, false
	}
	return o.Syncdate, true
}

// HasSyncdate returns a boolean if a field has been set.
func (o *WorkflowRest) HasSyncdate() bool {
	if o != nil && !IsNil(o.Syncdate) {
		return true
	}

	return false
}

// SetSyncdate gets a reference to the given int64 and assigns it to the Syncdate field.
func (o *WorkflowRest) SetSyncdate(v int64) {
	o.Syncdate = &v
}

// GetStale returns the Stale field value if set, zero value otherwise.
func (o *WorkflowRest) GetStale() bool {
	if o == nil || IsNil(o.Stale) {
		var ret bool
		return ret
	}
	return *o.Stale
}

// GetStaleOk returns a tuple with the Stale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRest) GetStaleOk() (*bool, bool) {
	if o == nil || IsNil(o.Stale) {
		return nil, false
	}
	return o.Stale, true
}

// HasStale returns a boolean if a field has been set.
func (o *WorkflowRest) HasStale() bool {
	if o != nil && !IsNil(o.Stale) {
		return true
	}

	return false
}

// SetStale gets a reference to the given bool and assigns it to the Stale field.
func (o *WorkflowRest) SetStale(v bool) {
	o.Stale = &v
}

func (o WorkflowRest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Props) {
		toSerialize["props"] = o.Props
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
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

type NullableWorkflowRest struct {
	value *WorkflowRest
	isSet bool
}

func (v NullableWorkflowRest) Get() *WorkflowRest {
	return v.value
}

func (v *NullableWorkflowRest) Set(val *WorkflowRest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRest(val *WorkflowRest) *NullableWorkflowRest {
	return &NullableWorkflowRest{value: val, isSet: true}
}

func (v NullableWorkflowRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


