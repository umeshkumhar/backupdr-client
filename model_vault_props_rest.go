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

// checks if the VaultPropsRest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VaultPropsRest{}

// VaultPropsRest struct for VaultPropsRest
type VaultPropsRest struct {
	Vaulttype *string `json:"vaulttype,omitempty"`
	Bucket *string `json:"bucket,omitempty"`
	Compression *bool `json:"compression,omitempty"`
	Objectsize *int64 `json:"objectsize,omitempty"`
	Accesskey *string `json:"accesskey,omitempty"`
	Accessid *string `json:"accessid,omitempty"`
	Baseurl *string `json:"baseurl,omitempty"`
	Authversion *string `json:"authversion,omitempty"`
	Region *string `json:"region,omitempty"`
	// Unique ID for this object
	Id *string `json:"id,omitempty"`
	// URL to access this object
	Href *string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate *int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale *bool `json:"stale,omitempty"`
}

// NewVaultPropsRest instantiates a new VaultPropsRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaultPropsRest() *VaultPropsRest {
	this := VaultPropsRest{}
	return &this
}

// NewVaultPropsRestWithDefaults instantiates a new VaultPropsRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultPropsRestWithDefaults() *VaultPropsRest {
	this := VaultPropsRest{}
	return &this
}

// GetVaulttype returns the Vaulttype field value if set, zero value otherwise.
func (o *VaultPropsRest) GetVaulttype() string {
	if o == nil || IsNil(o.Vaulttype) {
		var ret string
		return ret
	}
	return *o.Vaulttype
}

// GetVaulttypeOk returns a tuple with the Vaulttype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetVaulttypeOk() (*string, bool) {
	if o == nil || IsNil(o.Vaulttype) {
		return nil, false
	}
	return o.Vaulttype, true
}

// HasVaulttype returns a boolean if a field has been set.
func (o *VaultPropsRest) HasVaulttype() bool {
	if o != nil && !IsNil(o.Vaulttype) {
		return true
	}

	return false
}

// SetVaulttype gets a reference to the given string and assigns it to the Vaulttype field.
func (o *VaultPropsRest) SetVaulttype(v string) {
	o.Vaulttype = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *VaultPropsRest) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *VaultPropsRest) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *VaultPropsRest) SetBucket(v string) {
	o.Bucket = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *VaultPropsRest) GetCompression() bool {
	if o == nil || IsNil(o.Compression) {
		var ret bool
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetCompressionOk() (*bool, bool) {
	if o == nil || IsNil(o.Compression) {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *VaultPropsRest) HasCompression() bool {
	if o != nil && !IsNil(o.Compression) {
		return true
	}

	return false
}

// SetCompression gets a reference to the given bool and assigns it to the Compression field.
func (o *VaultPropsRest) SetCompression(v bool) {
	o.Compression = &v
}

// GetObjectsize returns the Objectsize field value if set, zero value otherwise.
func (o *VaultPropsRest) GetObjectsize() int64 {
	if o == nil || IsNil(o.Objectsize) {
		var ret int64
		return ret
	}
	return *o.Objectsize
}

// GetObjectsizeOk returns a tuple with the Objectsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetObjectsizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Objectsize) {
		return nil, false
	}
	return o.Objectsize, true
}

// HasObjectsize returns a boolean if a field has been set.
func (o *VaultPropsRest) HasObjectsize() bool {
	if o != nil && !IsNil(o.Objectsize) {
		return true
	}

	return false
}

// SetObjectsize gets a reference to the given int64 and assigns it to the Objectsize field.
func (o *VaultPropsRest) SetObjectsize(v int64) {
	o.Objectsize = &v
}

// GetAccesskey returns the Accesskey field value if set, zero value otherwise.
func (o *VaultPropsRest) GetAccesskey() string {
	if o == nil || IsNil(o.Accesskey) {
		var ret string
		return ret
	}
	return *o.Accesskey
}

// GetAccesskeyOk returns a tuple with the Accesskey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetAccesskeyOk() (*string, bool) {
	if o == nil || IsNil(o.Accesskey) {
		return nil, false
	}
	return o.Accesskey, true
}

// HasAccesskey returns a boolean if a field has been set.
func (o *VaultPropsRest) HasAccesskey() bool {
	if o != nil && !IsNil(o.Accesskey) {
		return true
	}

	return false
}

// SetAccesskey gets a reference to the given string and assigns it to the Accesskey field.
func (o *VaultPropsRest) SetAccesskey(v string) {
	o.Accesskey = &v
}

// GetAccessid returns the Accessid field value if set, zero value otherwise.
func (o *VaultPropsRest) GetAccessid() string {
	if o == nil || IsNil(o.Accessid) {
		var ret string
		return ret
	}
	return *o.Accessid
}

// GetAccessidOk returns a tuple with the Accessid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetAccessidOk() (*string, bool) {
	if o == nil || IsNil(o.Accessid) {
		return nil, false
	}
	return o.Accessid, true
}

// HasAccessid returns a boolean if a field has been set.
func (o *VaultPropsRest) HasAccessid() bool {
	if o != nil && !IsNil(o.Accessid) {
		return true
	}

	return false
}

// SetAccessid gets a reference to the given string and assigns it to the Accessid field.
func (o *VaultPropsRest) SetAccessid(v string) {
	o.Accessid = &v
}

// GetBaseurl returns the Baseurl field value if set, zero value otherwise.
func (o *VaultPropsRest) GetBaseurl() string {
	if o == nil || IsNil(o.Baseurl) {
		var ret string
		return ret
	}
	return *o.Baseurl
}

// GetBaseurlOk returns a tuple with the Baseurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetBaseurlOk() (*string, bool) {
	if o == nil || IsNil(o.Baseurl) {
		return nil, false
	}
	return o.Baseurl, true
}

// HasBaseurl returns a boolean if a field has been set.
func (o *VaultPropsRest) HasBaseurl() bool {
	if o != nil && !IsNil(o.Baseurl) {
		return true
	}

	return false
}

// SetBaseurl gets a reference to the given string and assigns it to the Baseurl field.
func (o *VaultPropsRest) SetBaseurl(v string) {
	o.Baseurl = &v
}

// GetAuthversion returns the Authversion field value if set, zero value otherwise.
func (o *VaultPropsRest) GetAuthversion() string {
	if o == nil || IsNil(o.Authversion) {
		var ret string
		return ret
	}
	return *o.Authversion
}

// GetAuthversionOk returns a tuple with the Authversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetAuthversionOk() (*string, bool) {
	if o == nil || IsNil(o.Authversion) {
		return nil, false
	}
	return o.Authversion, true
}

// HasAuthversion returns a boolean if a field has been set.
func (o *VaultPropsRest) HasAuthversion() bool {
	if o != nil && !IsNil(o.Authversion) {
		return true
	}

	return false
}

// SetAuthversion gets a reference to the given string and assigns it to the Authversion field.
func (o *VaultPropsRest) SetAuthversion(v string) {
	o.Authversion = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *VaultPropsRest) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *VaultPropsRest) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *VaultPropsRest) SetRegion(v string) {
	o.Region = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VaultPropsRest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VaultPropsRest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VaultPropsRest) SetId(v string) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *VaultPropsRest) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *VaultPropsRest) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *VaultPropsRest) SetHref(v string) {
	o.Href = &v
}

// GetSyncdate returns the Syncdate field value if set, zero value otherwise.
func (o *VaultPropsRest) GetSyncdate() int64 {
	if o == nil || IsNil(o.Syncdate) {
		var ret int64
		return ret
	}
	return *o.Syncdate
}

// GetSyncdateOk returns a tuple with the Syncdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetSyncdateOk() (*int64, bool) {
	if o == nil || IsNil(o.Syncdate) {
		return nil, false
	}
	return o.Syncdate, true
}

// HasSyncdate returns a boolean if a field has been set.
func (o *VaultPropsRest) HasSyncdate() bool {
	if o != nil && !IsNil(o.Syncdate) {
		return true
	}

	return false
}

// SetSyncdate gets a reference to the given int64 and assigns it to the Syncdate field.
func (o *VaultPropsRest) SetSyncdate(v int64) {
	o.Syncdate = &v
}

// GetStale returns the Stale field value if set, zero value otherwise.
func (o *VaultPropsRest) GetStale() bool {
	if o == nil || IsNil(o.Stale) {
		var ret bool
		return ret
	}
	return *o.Stale
}

// GetStaleOk returns a tuple with the Stale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPropsRest) GetStaleOk() (*bool, bool) {
	if o == nil || IsNil(o.Stale) {
		return nil, false
	}
	return o.Stale, true
}

// HasStale returns a boolean if a field has been set.
func (o *VaultPropsRest) HasStale() bool {
	if o != nil && !IsNil(o.Stale) {
		return true
	}

	return false
}

// SetStale gets a reference to the given bool and assigns it to the Stale field.
func (o *VaultPropsRest) SetStale(v bool) {
	o.Stale = &v
}

func (o VaultPropsRest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VaultPropsRest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Vaulttype) {
		toSerialize["vaulttype"] = o.Vaulttype
	}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.Compression) {
		toSerialize["compression"] = o.Compression
	}
	if !IsNil(o.Objectsize) {
		toSerialize["objectsize"] = o.Objectsize
	}
	if !IsNil(o.Accesskey) {
		toSerialize["accesskey"] = o.Accesskey
	}
	if !IsNil(o.Accessid) {
		toSerialize["accessid"] = o.Accessid
	}
	if !IsNil(o.Baseurl) {
		toSerialize["baseurl"] = o.Baseurl
	}
	if !IsNil(o.Authversion) {
		toSerialize["authversion"] = o.Authversion
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
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

type NullableVaultPropsRest struct {
	value *VaultPropsRest
	isSet bool
}

func (v NullableVaultPropsRest) Get() *VaultPropsRest {
	return v.value
}

func (v *NullableVaultPropsRest) Set(val *VaultPropsRest) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultPropsRest) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultPropsRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultPropsRest(val *VaultPropsRest) *NullableVaultPropsRest {
	return &NullableVaultPropsRest{value: val, isSet: true}
}

func (v NullableVaultPropsRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultPropsRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


