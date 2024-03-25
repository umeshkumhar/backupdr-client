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

// checks if the SlpRest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlpRest{}

// SlpRest Available SLP (profiles) that this application can potentially use for protection.
type SlpRest struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Srcid *string `json:"srcid,omitempty"`
	Clusterid *string `json:"clusterid,omitempty"`
	Modifydate *int64 `json:"modifydate,omitempty"`
	Cid *string `json:"cid,omitempty"`
	Performancepool *string `json:"performancepool,omitempty"`
	Primarystorage *string `json:"primarystorage,omitempty"`
	Remotenode *string `json:"remotenode,omitempty"`
	Dedupasyncnode *string `json:"dedupasyncnode,omitempty"`
	Vaultpool *DiskPoolRest `json:"vaultpool,omitempty"`
	Vaultpool2 *DiskPoolRest `json:"vaultpool2,omitempty"`
	Vaultpool3 *DiskPoolRest `json:"vaultpool3,omitempty"`
	Vaultpool4 *DiskPoolRest `json:"vaultpool4,omitempty"`
	Createdate *int64 `json:"createdate,omitempty"`
	Localnode *string `json:"localnode,omitempty"`
	Orglist []OrganizationRest `json:"orglist,omitempty"`
	CloudCredential *CloudCredentialRest `json:"cloudCredential,omitempty"`
	// Unique ID for this object
	Id *string `json:"id,omitempty"`
	// URL to access this object
	Href *string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate *int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale *bool `json:"stale,omitempty"`
}

// NewSlpRest instantiates a new SlpRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlpRest() *SlpRest {
	this := SlpRest{}
	return &this
}

// NewSlpRestWithDefaults instantiates a new SlpRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlpRestWithDefaults() *SlpRest {
	this := SlpRest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SlpRest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SlpRest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SlpRest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SlpRest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SlpRest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SlpRest) SetName(v string) {
	o.Name = &v
}

// GetSrcid returns the Srcid field value if set, zero value otherwise.
func (o *SlpRest) GetSrcid() string {
	if o == nil || IsNil(o.Srcid) {
		var ret string
		return ret
	}
	return *o.Srcid
}

// GetSrcidOk returns a tuple with the Srcid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetSrcidOk() (*string, bool) {
	if o == nil || IsNil(o.Srcid) {
		return nil, false
	}
	return o.Srcid, true
}

// HasSrcid returns a boolean if a field has been set.
func (o *SlpRest) HasSrcid() bool {
	if o != nil && !IsNil(o.Srcid) {
		return true
	}

	return false
}

// SetSrcid gets a reference to the given string and assigns it to the Srcid field.
func (o *SlpRest) SetSrcid(v string) {
	o.Srcid = &v
}

// GetClusterid returns the Clusterid field value if set, zero value otherwise.
func (o *SlpRest) GetClusterid() string {
	if o == nil || IsNil(o.Clusterid) {
		var ret string
		return ret
	}
	return *o.Clusterid
}

// GetClusteridOk returns a tuple with the Clusterid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetClusteridOk() (*string, bool) {
	if o == nil || IsNil(o.Clusterid) {
		return nil, false
	}
	return o.Clusterid, true
}

// HasClusterid returns a boolean if a field has been set.
func (o *SlpRest) HasClusterid() bool {
	if o != nil && !IsNil(o.Clusterid) {
		return true
	}

	return false
}

// SetClusterid gets a reference to the given string and assigns it to the Clusterid field.
func (o *SlpRest) SetClusterid(v string) {
	o.Clusterid = &v
}

// GetModifydate returns the Modifydate field value if set, zero value otherwise.
func (o *SlpRest) GetModifydate() int64 {
	if o == nil || IsNil(o.Modifydate) {
		var ret int64
		return ret
	}
	return *o.Modifydate
}

// GetModifydateOk returns a tuple with the Modifydate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetModifydateOk() (*int64, bool) {
	if o == nil || IsNil(o.Modifydate) {
		return nil, false
	}
	return o.Modifydate, true
}

// HasModifydate returns a boolean if a field has been set.
func (o *SlpRest) HasModifydate() bool {
	if o != nil && !IsNil(o.Modifydate) {
		return true
	}

	return false
}

// SetModifydate gets a reference to the given int64 and assigns it to the Modifydate field.
func (o *SlpRest) SetModifydate(v int64) {
	o.Modifydate = &v
}

// GetCid returns the Cid field value if set, zero value otherwise.
func (o *SlpRest) GetCid() string {
	if o == nil || IsNil(o.Cid) {
		var ret string
		return ret
	}
	return *o.Cid
}

// GetCidOk returns a tuple with the Cid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetCidOk() (*string, bool) {
	if o == nil || IsNil(o.Cid) {
		return nil, false
	}
	return o.Cid, true
}

// HasCid returns a boolean if a field has been set.
func (o *SlpRest) HasCid() bool {
	if o != nil && !IsNil(o.Cid) {
		return true
	}

	return false
}

// SetCid gets a reference to the given string and assigns it to the Cid field.
func (o *SlpRest) SetCid(v string) {
	o.Cid = &v
}

// GetPerformancepool returns the Performancepool field value if set, zero value otherwise.
func (o *SlpRest) GetPerformancepool() string {
	if o == nil || IsNil(o.Performancepool) {
		var ret string
		return ret
	}
	return *o.Performancepool
}

// GetPerformancepoolOk returns a tuple with the Performancepool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetPerformancepoolOk() (*string, bool) {
	if o == nil || IsNil(o.Performancepool) {
		return nil, false
	}
	return o.Performancepool, true
}

// HasPerformancepool returns a boolean if a field has been set.
func (o *SlpRest) HasPerformancepool() bool {
	if o != nil && !IsNil(o.Performancepool) {
		return true
	}

	return false
}

// SetPerformancepool gets a reference to the given string and assigns it to the Performancepool field.
func (o *SlpRest) SetPerformancepool(v string) {
	o.Performancepool = &v
}

// GetPrimarystorage returns the Primarystorage field value if set, zero value otherwise.
func (o *SlpRest) GetPrimarystorage() string {
	if o == nil || IsNil(o.Primarystorage) {
		var ret string
		return ret
	}
	return *o.Primarystorage
}

// GetPrimarystorageOk returns a tuple with the Primarystorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetPrimarystorageOk() (*string, bool) {
	if o == nil || IsNil(o.Primarystorage) {
		return nil, false
	}
	return o.Primarystorage, true
}

// HasPrimarystorage returns a boolean if a field has been set.
func (o *SlpRest) HasPrimarystorage() bool {
	if o != nil && !IsNil(o.Primarystorage) {
		return true
	}

	return false
}

// SetPrimarystorage gets a reference to the given string and assigns it to the Primarystorage field.
func (o *SlpRest) SetPrimarystorage(v string) {
	o.Primarystorage = &v
}

// GetRemotenode returns the Remotenode field value if set, zero value otherwise.
func (o *SlpRest) GetRemotenode() string {
	if o == nil || IsNil(o.Remotenode) {
		var ret string
		return ret
	}
	return *o.Remotenode
}

// GetRemotenodeOk returns a tuple with the Remotenode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetRemotenodeOk() (*string, bool) {
	if o == nil || IsNil(o.Remotenode) {
		return nil, false
	}
	return o.Remotenode, true
}

// HasRemotenode returns a boolean if a field has been set.
func (o *SlpRest) HasRemotenode() bool {
	if o != nil && !IsNil(o.Remotenode) {
		return true
	}

	return false
}

// SetRemotenode gets a reference to the given string and assigns it to the Remotenode field.
func (o *SlpRest) SetRemotenode(v string) {
	o.Remotenode = &v
}

// GetDedupasyncnode returns the Dedupasyncnode field value if set, zero value otherwise.
func (o *SlpRest) GetDedupasyncnode() string {
	if o == nil || IsNil(o.Dedupasyncnode) {
		var ret string
		return ret
	}
	return *o.Dedupasyncnode
}

// GetDedupasyncnodeOk returns a tuple with the Dedupasyncnode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetDedupasyncnodeOk() (*string, bool) {
	if o == nil || IsNil(o.Dedupasyncnode) {
		return nil, false
	}
	return o.Dedupasyncnode, true
}

// HasDedupasyncnode returns a boolean if a field has been set.
func (o *SlpRest) HasDedupasyncnode() bool {
	if o != nil && !IsNil(o.Dedupasyncnode) {
		return true
	}

	return false
}

// SetDedupasyncnode gets a reference to the given string and assigns it to the Dedupasyncnode field.
func (o *SlpRest) SetDedupasyncnode(v string) {
	o.Dedupasyncnode = &v
}

// GetVaultpool returns the Vaultpool field value if set, zero value otherwise.
func (o *SlpRest) GetVaultpool() DiskPoolRest {
	if o == nil || IsNil(o.Vaultpool) {
		var ret DiskPoolRest
		return ret
	}
	return *o.Vaultpool
}

// GetVaultpoolOk returns a tuple with the Vaultpool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetVaultpoolOk() (*DiskPoolRest, bool) {
	if o == nil || IsNil(o.Vaultpool) {
		return nil, false
	}
	return o.Vaultpool, true
}

// HasVaultpool returns a boolean if a field has been set.
func (o *SlpRest) HasVaultpool() bool {
	if o != nil && !IsNil(o.Vaultpool) {
		return true
	}

	return false
}

// SetVaultpool gets a reference to the given DiskPoolRest and assigns it to the Vaultpool field.
func (o *SlpRest) SetVaultpool(v DiskPoolRest) {
	o.Vaultpool = &v
}

// GetVaultpool2 returns the Vaultpool2 field value if set, zero value otherwise.
func (o *SlpRest) GetVaultpool2() DiskPoolRest {
	if o == nil || IsNil(o.Vaultpool2) {
		var ret DiskPoolRest
		return ret
	}
	return *o.Vaultpool2
}

// GetVaultpool2Ok returns a tuple with the Vaultpool2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetVaultpool2Ok() (*DiskPoolRest, bool) {
	if o == nil || IsNil(o.Vaultpool2) {
		return nil, false
	}
	return o.Vaultpool2, true
}

// HasVaultpool2 returns a boolean if a field has been set.
func (o *SlpRest) HasVaultpool2() bool {
	if o != nil && !IsNil(o.Vaultpool2) {
		return true
	}

	return false
}

// SetVaultpool2 gets a reference to the given DiskPoolRest and assigns it to the Vaultpool2 field.
func (o *SlpRest) SetVaultpool2(v DiskPoolRest) {
	o.Vaultpool2 = &v
}

// GetVaultpool3 returns the Vaultpool3 field value if set, zero value otherwise.
func (o *SlpRest) GetVaultpool3() DiskPoolRest {
	if o == nil || IsNil(o.Vaultpool3) {
		var ret DiskPoolRest
		return ret
	}
	return *o.Vaultpool3
}

// GetVaultpool3Ok returns a tuple with the Vaultpool3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetVaultpool3Ok() (*DiskPoolRest, bool) {
	if o == nil || IsNil(o.Vaultpool3) {
		return nil, false
	}
	return o.Vaultpool3, true
}

// HasVaultpool3 returns a boolean if a field has been set.
func (o *SlpRest) HasVaultpool3() bool {
	if o != nil && !IsNil(o.Vaultpool3) {
		return true
	}

	return false
}

// SetVaultpool3 gets a reference to the given DiskPoolRest and assigns it to the Vaultpool3 field.
func (o *SlpRest) SetVaultpool3(v DiskPoolRest) {
	o.Vaultpool3 = &v
}

// GetVaultpool4 returns the Vaultpool4 field value if set, zero value otherwise.
func (o *SlpRest) GetVaultpool4() DiskPoolRest {
	if o == nil || IsNil(o.Vaultpool4) {
		var ret DiskPoolRest
		return ret
	}
	return *o.Vaultpool4
}

// GetVaultpool4Ok returns a tuple with the Vaultpool4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetVaultpool4Ok() (*DiskPoolRest, bool) {
	if o == nil || IsNil(o.Vaultpool4) {
		return nil, false
	}
	return o.Vaultpool4, true
}

// HasVaultpool4 returns a boolean if a field has been set.
func (o *SlpRest) HasVaultpool4() bool {
	if o != nil && !IsNil(o.Vaultpool4) {
		return true
	}

	return false
}

// SetVaultpool4 gets a reference to the given DiskPoolRest and assigns it to the Vaultpool4 field.
func (o *SlpRest) SetVaultpool4(v DiskPoolRest) {
	o.Vaultpool4 = &v
}

// GetCreatedate returns the Createdate field value if set, zero value otherwise.
func (o *SlpRest) GetCreatedate() int64 {
	if o == nil || IsNil(o.Createdate) {
		var ret int64
		return ret
	}
	return *o.Createdate
}

// GetCreatedateOk returns a tuple with the Createdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetCreatedateOk() (*int64, bool) {
	if o == nil || IsNil(o.Createdate) {
		return nil, false
	}
	return o.Createdate, true
}

// HasCreatedate returns a boolean if a field has been set.
func (o *SlpRest) HasCreatedate() bool {
	if o != nil && !IsNil(o.Createdate) {
		return true
	}

	return false
}

// SetCreatedate gets a reference to the given int64 and assigns it to the Createdate field.
func (o *SlpRest) SetCreatedate(v int64) {
	o.Createdate = &v
}

// GetLocalnode returns the Localnode field value if set, zero value otherwise.
func (o *SlpRest) GetLocalnode() string {
	if o == nil || IsNil(o.Localnode) {
		var ret string
		return ret
	}
	return *o.Localnode
}

// GetLocalnodeOk returns a tuple with the Localnode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetLocalnodeOk() (*string, bool) {
	if o == nil || IsNil(o.Localnode) {
		return nil, false
	}
	return o.Localnode, true
}

// HasLocalnode returns a boolean if a field has been set.
func (o *SlpRest) HasLocalnode() bool {
	if o != nil && !IsNil(o.Localnode) {
		return true
	}

	return false
}

// SetLocalnode gets a reference to the given string and assigns it to the Localnode field.
func (o *SlpRest) SetLocalnode(v string) {
	o.Localnode = &v
}

// GetOrglist returns the Orglist field value if set, zero value otherwise.
func (o *SlpRest) GetOrglist() []OrganizationRest {
	if o == nil || IsNil(o.Orglist) {
		var ret []OrganizationRest
		return ret
	}
	return o.Orglist
}

// GetOrglistOk returns a tuple with the Orglist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetOrglistOk() ([]OrganizationRest, bool) {
	if o == nil || IsNil(o.Orglist) {
		return nil, false
	}
	return o.Orglist, true
}

// HasOrglist returns a boolean if a field has been set.
func (o *SlpRest) HasOrglist() bool {
	if o != nil && !IsNil(o.Orglist) {
		return true
	}

	return false
}

// SetOrglist gets a reference to the given []OrganizationRest and assigns it to the Orglist field.
func (o *SlpRest) SetOrglist(v []OrganizationRest) {
	o.Orglist = v
}

// GetCloudCredential returns the CloudCredential field value if set, zero value otherwise.
func (o *SlpRest) GetCloudCredential() CloudCredentialRest {
	if o == nil || IsNil(o.CloudCredential) {
		var ret CloudCredentialRest
		return ret
	}
	return *o.CloudCredential
}

// GetCloudCredentialOk returns a tuple with the CloudCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetCloudCredentialOk() (*CloudCredentialRest, bool) {
	if o == nil || IsNil(o.CloudCredential) {
		return nil, false
	}
	return o.CloudCredential, true
}

// HasCloudCredential returns a boolean if a field has been set.
func (o *SlpRest) HasCloudCredential() bool {
	if o != nil && !IsNil(o.CloudCredential) {
		return true
	}

	return false
}

// SetCloudCredential gets a reference to the given CloudCredentialRest and assigns it to the CloudCredential field.
func (o *SlpRest) SetCloudCredential(v CloudCredentialRest) {
	o.CloudCredential = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SlpRest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SlpRest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SlpRest) SetId(v string) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SlpRest) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SlpRest) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SlpRest) SetHref(v string) {
	o.Href = &v
}

// GetSyncdate returns the Syncdate field value if set, zero value otherwise.
func (o *SlpRest) GetSyncdate() int64 {
	if o == nil || IsNil(o.Syncdate) {
		var ret int64
		return ret
	}
	return *o.Syncdate
}

// GetSyncdateOk returns a tuple with the Syncdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetSyncdateOk() (*int64, bool) {
	if o == nil || IsNil(o.Syncdate) {
		return nil, false
	}
	return o.Syncdate, true
}

// HasSyncdate returns a boolean if a field has been set.
func (o *SlpRest) HasSyncdate() bool {
	if o != nil && !IsNil(o.Syncdate) {
		return true
	}

	return false
}

// SetSyncdate gets a reference to the given int64 and assigns it to the Syncdate field.
func (o *SlpRest) SetSyncdate(v int64) {
	o.Syncdate = &v
}

// GetStale returns the Stale field value if set, zero value otherwise.
func (o *SlpRest) GetStale() bool {
	if o == nil || IsNil(o.Stale) {
		var ret bool
		return ret
	}
	return *o.Stale
}

// GetStaleOk returns a tuple with the Stale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlpRest) GetStaleOk() (*bool, bool) {
	if o == nil || IsNil(o.Stale) {
		return nil, false
	}
	return o.Stale, true
}

// HasStale returns a boolean if a field has been set.
func (o *SlpRest) HasStale() bool {
	if o != nil && !IsNil(o.Stale) {
		return true
	}

	return false
}

// SetStale gets a reference to the given bool and assigns it to the Stale field.
func (o *SlpRest) SetStale(v bool) {
	o.Stale = &v
}

func (o SlpRest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlpRest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Srcid) {
		toSerialize["srcid"] = o.Srcid
	}
	if !IsNil(o.Clusterid) {
		toSerialize["clusterid"] = o.Clusterid
	}
	if !IsNil(o.Modifydate) {
		toSerialize["modifydate"] = o.Modifydate
	}
	if !IsNil(o.Cid) {
		toSerialize["cid"] = o.Cid
	}
	if !IsNil(o.Performancepool) {
		toSerialize["performancepool"] = o.Performancepool
	}
	if !IsNil(o.Primarystorage) {
		toSerialize["primarystorage"] = o.Primarystorage
	}
	if !IsNil(o.Remotenode) {
		toSerialize["remotenode"] = o.Remotenode
	}
	if !IsNil(o.Dedupasyncnode) {
		toSerialize["dedupasyncnode"] = o.Dedupasyncnode
	}
	if !IsNil(o.Vaultpool) {
		toSerialize["vaultpool"] = o.Vaultpool
	}
	if !IsNil(o.Vaultpool2) {
		toSerialize["vaultpool2"] = o.Vaultpool2
	}
	if !IsNil(o.Vaultpool3) {
		toSerialize["vaultpool3"] = o.Vaultpool3
	}
	if !IsNil(o.Vaultpool4) {
		toSerialize["vaultpool4"] = o.Vaultpool4
	}
	if !IsNil(o.Createdate) {
		toSerialize["createdate"] = o.Createdate
	}
	if !IsNil(o.Localnode) {
		toSerialize["localnode"] = o.Localnode
	}
	if !IsNil(o.Orglist) {
		toSerialize["orglist"] = o.Orglist
	}
	if !IsNil(o.CloudCredential) {
		toSerialize["cloudCredential"] = o.CloudCredential
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

type NullableSlpRest struct {
	value *SlpRest
	isSet bool
}

func (v NullableSlpRest) Get() *SlpRest {
	return v.value
}

func (v *NullableSlpRest) Set(val *SlpRest) {
	v.value = val
	v.isSet = true
}

func (v NullableSlpRest) IsSet() bool {
	return v.isSet
}

func (v *NullableSlpRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlpRest(val *SlpRest) *NullableSlpRest {
	return &NullableSlpRest{value: val, isSet: true}
}

func (v NullableSlpRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlpRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


