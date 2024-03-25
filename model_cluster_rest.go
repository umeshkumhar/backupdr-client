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

// checks if the ClusterRest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterRest{}

// ClusterRest Appliance where the application resides.
type ClusterRest struct {
	SharedSecret *string `json:"shared_secret,omitempty"`
	Serviceaccount *string `json:"serviceaccount,omitempty"`
	Zone *string `json:"zone,omitempty"`
	Region *string `json:"region,omitempty"`
	Projectid *string `json:"projectid,omitempty"`
	Password *string `json:"password,omitempty"`
	Version *string `json:"version,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Protocol *int32 `json:"protocol,omitempty"`
	Datacenter *string `json:"datacenter,omitempty"`
	Masterid *string `json:"masterid,omitempty"`
	Clusterid *string `json:"clusterid,omitempty"`
	Importstatus *string `json:"importstatus,omitempty"`
	Username *string `json:"username,omitempty"`
	Ipaddress *string `json:"ipaddress,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Clusterstatus *ClusterStatusRest `json:"clusterstatus,omitempty"`
	Lastsync *int64 `json:"lastsync,omitempty"`
	Publicip *string `json:"publicip,omitempty"`
	Secureconnect *bool `json:"secureconnect,omitempty"`
	PkiBootstrapped *bool `json:"pkiBootstrapped,omitempty"`
	Clusterlist []ClusterRest `json:"clusterlist,omitempty"`
	CallhomeInfo *CallhomeInfoRest `json:"callhomeInfo,omitempty"`
	Rmipaddress []string `json:"rmipaddress,omitempty"`
	// Shows the support status of cluster
	Supportstatus *string `json:"supportstatus,omitempty"`
	// Unique ID for this object
	Id *string `json:"id,omitempty"`
	// URL to access this object
	Href *string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate *int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale *bool `json:"stale,omitempty"`
}

// NewClusterRest instantiates a new ClusterRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRest() *ClusterRest {
	this := ClusterRest{}
	return &this
}

// NewClusterRestWithDefaults instantiates a new ClusterRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRestWithDefaults() *ClusterRest {
	this := ClusterRest{}
	return &this
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *ClusterRest) GetSharedSecret() string {
	if o == nil || IsNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetSharedSecretOk() (*string, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *ClusterRest) HasSharedSecret() bool {
	if o != nil && !IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *ClusterRest) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetServiceaccount returns the Serviceaccount field value if set, zero value otherwise.
func (o *ClusterRest) GetServiceaccount() string {
	if o == nil || IsNil(o.Serviceaccount) {
		var ret string
		return ret
	}
	return *o.Serviceaccount
}

// GetServiceaccountOk returns a tuple with the Serviceaccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetServiceaccountOk() (*string, bool) {
	if o == nil || IsNil(o.Serviceaccount) {
		return nil, false
	}
	return o.Serviceaccount, true
}

// HasServiceaccount returns a boolean if a field has been set.
func (o *ClusterRest) HasServiceaccount() bool {
	if o != nil && !IsNil(o.Serviceaccount) {
		return true
	}

	return false
}

// SetServiceaccount gets a reference to the given string and assigns it to the Serviceaccount field.
func (o *ClusterRest) SetServiceaccount(v string) {
	o.Serviceaccount = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *ClusterRest) GetZone() string {
	if o == nil || IsNil(o.Zone) {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetZoneOk() (*string, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *ClusterRest) HasZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *ClusterRest) SetZone(v string) {
	o.Zone = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ClusterRest) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ClusterRest) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ClusterRest) SetRegion(v string) {
	o.Region = &v
}

// GetProjectid returns the Projectid field value if set, zero value otherwise.
func (o *ClusterRest) GetProjectid() string {
	if o == nil || IsNil(o.Projectid) {
		var ret string
		return ret
	}
	return *o.Projectid
}

// GetProjectidOk returns a tuple with the Projectid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetProjectidOk() (*string, bool) {
	if o == nil || IsNil(o.Projectid) {
		return nil, false
	}
	return o.Projectid, true
}

// HasProjectid returns a boolean if a field has been set.
func (o *ClusterRest) HasProjectid() bool {
	if o != nil && !IsNil(o.Projectid) {
		return true
	}

	return false
}

// SetProjectid gets a reference to the given string and assigns it to the Projectid field.
func (o *ClusterRest) SetProjectid(v string) {
	o.Projectid = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ClusterRest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ClusterRest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ClusterRest) SetPassword(v string) {
	o.Password = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ClusterRest) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ClusterRest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ClusterRest) SetVersion(v string) {
	o.Version = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClusterRest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterRest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClusterRest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterRest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterRest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterRest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterRest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterRest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterRest) SetType(v string) {
	o.Type = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ClusterRest) GetProtocol() int32 {
	if o == nil || IsNil(o.Protocol) {
		var ret int32
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetProtocolOk() (*int32, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ClusterRest) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given int32 and assigns it to the Protocol field.
func (o *ClusterRest) SetProtocol(v int32) {
	o.Protocol = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *ClusterRest) GetDatacenter() string {
	if o == nil || IsNil(o.Datacenter) {
		var ret string
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetDatacenterOk() (*string, bool) {
	if o == nil || IsNil(o.Datacenter) {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *ClusterRest) HasDatacenter() bool {
	if o != nil && !IsNil(o.Datacenter) {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given string and assigns it to the Datacenter field.
func (o *ClusterRest) SetDatacenter(v string) {
	o.Datacenter = &v
}

// GetMasterid returns the Masterid field value if set, zero value otherwise.
func (o *ClusterRest) GetMasterid() string {
	if o == nil || IsNil(o.Masterid) {
		var ret string
		return ret
	}
	return *o.Masterid
}

// GetMasteridOk returns a tuple with the Masterid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetMasteridOk() (*string, bool) {
	if o == nil || IsNil(o.Masterid) {
		return nil, false
	}
	return o.Masterid, true
}

// HasMasterid returns a boolean if a field has been set.
func (o *ClusterRest) HasMasterid() bool {
	if o != nil && !IsNil(o.Masterid) {
		return true
	}

	return false
}

// SetMasterid gets a reference to the given string and assigns it to the Masterid field.
func (o *ClusterRest) SetMasterid(v string) {
	o.Masterid = &v
}

// GetClusterid returns the Clusterid field value if set, zero value otherwise.
func (o *ClusterRest) GetClusterid() string {
	if o == nil || IsNil(o.Clusterid) {
		var ret string
		return ret
	}
	return *o.Clusterid
}

// GetClusteridOk returns a tuple with the Clusterid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetClusteridOk() (*string, bool) {
	if o == nil || IsNil(o.Clusterid) {
		return nil, false
	}
	return o.Clusterid, true
}

// HasClusterid returns a boolean if a field has been set.
func (o *ClusterRest) HasClusterid() bool {
	if o != nil && !IsNil(o.Clusterid) {
		return true
	}

	return false
}

// SetClusterid gets a reference to the given string and assigns it to the Clusterid field.
func (o *ClusterRest) SetClusterid(v string) {
	o.Clusterid = &v
}

// GetImportstatus returns the Importstatus field value if set, zero value otherwise.
func (o *ClusterRest) GetImportstatus() string {
	if o == nil || IsNil(o.Importstatus) {
		var ret string
		return ret
	}
	return *o.Importstatus
}

// GetImportstatusOk returns a tuple with the Importstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetImportstatusOk() (*string, bool) {
	if o == nil || IsNil(o.Importstatus) {
		return nil, false
	}
	return o.Importstatus, true
}

// HasImportstatus returns a boolean if a field has been set.
func (o *ClusterRest) HasImportstatus() bool {
	if o != nil && !IsNil(o.Importstatus) {
		return true
	}

	return false
}

// SetImportstatus gets a reference to the given string and assigns it to the Importstatus field.
func (o *ClusterRest) SetImportstatus(v string) {
	o.Importstatus = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ClusterRest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ClusterRest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ClusterRest) SetUsername(v string) {
	o.Username = &v
}

// GetIpaddress returns the Ipaddress field value if set, zero value otherwise.
func (o *ClusterRest) GetIpaddress() string {
	if o == nil || IsNil(o.Ipaddress) {
		var ret string
		return ret
	}
	return *o.Ipaddress
}

// GetIpaddressOk returns a tuple with the Ipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetIpaddressOk() (*string, bool) {
	if o == nil || IsNil(o.Ipaddress) {
		return nil, false
	}
	return o.Ipaddress, true
}

// HasIpaddress returns a boolean if a field has been set.
func (o *ClusterRest) HasIpaddress() bool {
	if o != nil && !IsNil(o.Ipaddress) {
		return true
	}

	return false
}

// SetIpaddress gets a reference to the given string and assigns it to the Ipaddress field.
func (o *ClusterRest) SetIpaddress(v string) {
	o.Ipaddress = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ClusterRest) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ClusterRest) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ClusterRest) SetTimezone(v string) {
	o.Timezone = &v
}

// GetClusterstatus returns the Clusterstatus field value if set, zero value otherwise.
func (o *ClusterRest) GetClusterstatus() ClusterStatusRest {
	if o == nil || IsNil(o.Clusterstatus) {
		var ret ClusterStatusRest
		return ret
	}
	return *o.Clusterstatus
}

// GetClusterstatusOk returns a tuple with the Clusterstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetClusterstatusOk() (*ClusterStatusRest, bool) {
	if o == nil || IsNil(o.Clusterstatus) {
		return nil, false
	}
	return o.Clusterstatus, true
}

// HasClusterstatus returns a boolean if a field has been set.
func (o *ClusterRest) HasClusterstatus() bool {
	if o != nil && !IsNil(o.Clusterstatus) {
		return true
	}

	return false
}

// SetClusterstatus gets a reference to the given ClusterStatusRest and assigns it to the Clusterstatus field.
func (o *ClusterRest) SetClusterstatus(v ClusterStatusRest) {
	o.Clusterstatus = &v
}

// GetLastsync returns the Lastsync field value if set, zero value otherwise.
func (o *ClusterRest) GetLastsync() int64 {
	if o == nil || IsNil(o.Lastsync) {
		var ret int64
		return ret
	}
	return *o.Lastsync
}

// GetLastsyncOk returns a tuple with the Lastsync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetLastsyncOk() (*int64, bool) {
	if o == nil || IsNil(o.Lastsync) {
		return nil, false
	}
	return o.Lastsync, true
}

// HasLastsync returns a boolean if a field has been set.
func (o *ClusterRest) HasLastsync() bool {
	if o != nil && !IsNil(o.Lastsync) {
		return true
	}

	return false
}

// SetLastsync gets a reference to the given int64 and assigns it to the Lastsync field.
func (o *ClusterRest) SetLastsync(v int64) {
	o.Lastsync = &v
}

// GetPublicip returns the Publicip field value if set, zero value otherwise.
func (o *ClusterRest) GetPublicip() string {
	if o == nil || IsNil(o.Publicip) {
		var ret string
		return ret
	}
	return *o.Publicip
}

// GetPublicipOk returns a tuple with the Publicip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetPublicipOk() (*string, bool) {
	if o == nil || IsNil(o.Publicip) {
		return nil, false
	}
	return o.Publicip, true
}

// HasPublicip returns a boolean if a field has been set.
func (o *ClusterRest) HasPublicip() bool {
	if o != nil && !IsNil(o.Publicip) {
		return true
	}

	return false
}

// SetPublicip gets a reference to the given string and assigns it to the Publicip field.
func (o *ClusterRest) SetPublicip(v string) {
	o.Publicip = &v
}

// GetSecureconnect returns the Secureconnect field value if set, zero value otherwise.
func (o *ClusterRest) GetSecureconnect() bool {
	if o == nil || IsNil(o.Secureconnect) {
		var ret bool
		return ret
	}
	return *o.Secureconnect
}

// GetSecureconnectOk returns a tuple with the Secureconnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetSecureconnectOk() (*bool, bool) {
	if o == nil || IsNil(o.Secureconnect) {
		return nil, false
	}
	return o.Secureconnect, true
}

// HasSecureconnect returns a boolean if a field has been set.
func (o *ClusterRest) HasSecureconnect() bool {
	if o != nil && !IsNil(o.Secureconnect) {
		return true
	}

	return false
}

// SetSecureconnect gets a reference to the given bool and assigns it to the Secureconnect field.
func (o *ClusterRest) SetSecureconnect(v bool) {
	o.Secureconnect = &v
}

// GetPkiBootstrapped returns the PkiBootstrapped field value if set, zero value otherwise.
func (o *ClusterRest) GetPkiBootstrapped() bool {
	if o == nil || IsNil(o.PkiBootstrapped) {
		var ret bool
		return ret
	}
	return *o.PkiBootstrapped
}

// GetPkiBootstrappedOk returns a tuple with the PkiBootstrapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetPkiBootstrappedOk() (*bool, bool) {
	if o == nil || IsNil(o.PkiBootstrapped) {
		return nil, false
	}
	return o.PkiBootstrapped, true
}

// HasPkiBootstrapped returns a boolean if a field has been set.
func (o *ClusterRest) HasPkiBootstrapped() bool {
	if o != nil && !IsNil(o.PkiBootstrapped) {
		return true
	}

	return false
}

// SetPkiBootstrapped gets a reference to the given bool and assigns it to the PkiBootstrapped field.
func (o *ClusterRest) SetPkiBootstrapped(v bool) {
	o.PkiBootstrapped = &v
}

// GetClusterlist returns the Clusterlist field value if set, zero value otherwise.
func (o *ClusterRest) GetClusterlist() []ClusterRest {
	if o == nil || IsNil(o.Clusterlist) {
		var ret []ClusterRest
		return ret
	}
	return o.Clusterlist
}

// GetClusterlistOk returns a tuple with the Clusterlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetClusterlistOk() ([]ClusterRest, bool) {
	if o == nil || IsNil(o.Clusterlist) {
		return nil, false
	}
	return o.Clusterlist, true
}

// HasClusterlist returns a boolean if a field has been set.
func (o *ClusterRest) HasClusterlist() bool {
	if o != nil && !IsNil(o.Clusterlist) {
		return true
	}

	return false
}

// SetClusterlist gets a reference to the given []ClusterRest and assigns it to the Clusterlist field.
func (o *ClusterRest) SetClusterlist(v []ClusterRest) {
	o.Clusterlist = v
}

// GetCallhomeInfo returns the CallhomeInfo field value if set, zero value otherwise.
func (o *ClusterRest) GetCallhomeInfo() CallhomeInfoRest {
	if o == nil || IsNil(o.CallhomeInfo) {
		var ret CallhomeInfoRest
		return ret
	}
	return *o.CallhomeInfo
}

// GetCallhomeInfoOk returns a tuple with the CallhomeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetCallhomeInfoOk() (*CallhomeInfoRest, bool) {
	if o == nil || IsNil(o.CallhomeInfo) {
		return nil, false
	}
	return o.CallhomeInfo, true
}

// HasCallhomeInfo returns a boolean if a field has been set.
func (o *ClusterRest) HasCallhomeInfo() bool {
	if o != nil && !IsNil(o.CallhomeInfo) {
		return true
	}

	return false
}

// SetCallhomeInfo gets a reference to the given CallhomeInfoRest and assigns it to the CallhomeInfo field.
func (o *ClusterRest) SetCallhomeInfo(v CallhomeInfoRest) {
	o.CallhomeInfo = &v
}

// GetRmipaddress returns the Rmipaddress field value if set, zero value otherwise.
func (o *ClusterRest) GetRmipaddress() []string {
	if o == nil || IsNil(o.Rmipaddress) {
		var ret []string
		return ret
	}
	return o.Rmipaddress
}

// GetRmipaddressOk returns a tuple with the Rmipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetRmipaddressOk() ([]string, bool) {
	if o == nil || IsNil(o.Rmipaddress) {
		return nil, false
	}
	return o.Rmipaddress, true
}

// HasRmipaddress returns a boolean if a field has been set.
func (o *ClusterRest) HasRmipaddress() bool {
	if o != nil && !IsNil(o.Rmipaddress) {
		return true
	}

	return false
}

// SetRmipaddress gets a reference to the given []string and assigns it to the Rmipaddress field.
func (o *ClusterRest) SetRmipaddress(v []string) {
	o.Rmipaddress = v
}

// GetSupportstatus returns the Supportstatus field value if set, zero value otherwise.
func (o *ClusterRest) GetSupportstatus() string {
	if o == nil || IsNil(o.Supportstatus) {
		var ret string
		return ret
	}
	return *o.Supportstatus
}

// GetSupportstatusOk returns a tuple with the Supportstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetSupportstatusOk() (*string, bool) {
	if o == nil || IsNil(o.Supportstatus) {
		return nil, false
	}
	return o.Supportstatus, true
}

// HasSupportstatus returns a boolean if a field has been set.
func (o *ClusterRest) HasSupportstatus() bool {
	if o != nil && !IsNil(o.Supportstatus) {
		return true
	}

	return false
}

// SetSupportstatus gets a reference to the given string and assigns it to the Supportstatus field.
func (o *ClusterRest) SetSupportstatus(v string) {
	o.Supportstatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterRest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterRest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterRest) SetId(v string) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ClusterRest) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ClusterRest) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ClusterRest) SetHref(v string) {
	o.Href = &v
}

// GetSyncdate returns the Syncdate field value if set, zero value otherwise.
func (o *ClusterRest) GetSyncdate() int64 {
	if o == nil || IsNil(o.Syncdate) {
		var ret int64
		return ret
	}
	return *o.Syncdate
}

// GetSyncdateOk returns a tuple with the Syncdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetSyncdateOk() (*int64, bool) {
	if o == nil || IsNil(o.Syncdate) {
		return nil, false
	}
	return o.Syncdate, true
}

// HasSyncdate returns a boolean if a field has been set.
func (o *ClusterRest) HasSyncdate() bool {
	if o != nil && !IsNil(o.Syncdate) {
		return true
	}

	return false
}

// SetSyncdate gets a reference to the given int64 and assigns it to the Syncdate field.
func (o *ClusterRest) SetSyncdate(v int64) {
	o.Syncdate = &v
}

// GetStale returns the Stale field value if set, zero value otherwise.
func (o *ClusterRest) GetStale() bool {
	if o == nil || IsNil(o.Stale) {
		var ret bool
		return ret
	}
	return *o.Stale
}

// GetStaleOk returns a tuple with the Stale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRest) GetStaleOk() (*bool, bool) {
	if o == nil || IsNil(o.Stale) {
		return nil, false
	}
	return o.Stale, true
}

// HasStale returns a boolean if a field has been set.
func (o *ClusterRest) HasStale() bool {
	if o != nil && !IsNil(o.Stale) {
		return true
	}

	return false
}

// SetStale gets a reference to the given bool and assigns it to the Stale field.
func (o *ClusterRest) SetStale(v bool) {
	o.Stale = &v
}

func (o ClusterRest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterRest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SharedSecret) {
		toSerialize["shared_secret"] = o.SharedSecret
	}
	if !IsNil(o.Serviceaccount) {
		toSerialize["serviceaccount"] = o.Serviceaccount
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Projectid) {
		toSerialize["projectid"] = o.Projectid
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Datacenter) {
		toSerialize["datacenter"] = o.Datacenter
	}
	if !IsNil(o.Masterid) {
		toSerialize["masterid"] = o.Masterid
	}
	if !IsNil(o.Clusterid) {
		toSerialize["clusterid"] = o.Clusterid
	}
	if !IsNil(o.Importstatus) {
		toSerialize["importstatus"] = o.Importstatus
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Ipaddress) {
		toSerialize["ipaddress"] = o.Ipaddress
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Clusterstatus) {
		toSerialize["clusterstatus"] = o.Clusterstatus
	}
	if !IsNil(o.Lastsync) {
		toSerialize["lastsync"] = o.Lastsync
	}
	if !IsNil(o.Publicip) {
		toSerialize["publicip"] = o.Publicip
	}
	if !IsNil(o.Secureconnect) {
		toSerialize["secureconnect"] = o.Secureconnect
	}
	if !IsNil(o.PkiBootstrapped) {
		toSerialize["pkiBootstrapped"] = o.PkiBootstrapped
	}
	if !IsNil(o.Clusterlist) {
		toSerialize["clusterlist"] = o.Clusterlist
	}
	if !IsNil(o.CallhomeInfo) {
		toSerialize["callhomeInfo"] = o.CallhomeInfo
	}
	if !IsNil(o.Rmipaddress) {
		toSerialize["rmipaddress"] = o.Rmipaddress
	}
	if !IsNil(o.Supportstatus) {
		toSerialize["supportstatus"] = o.Supportstatus
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

type NullableClusterRest struct {
	value *ClusterRest
	isSet bool
}

func (v NullableClusterRest) Get() *ClusterRest {
	return v.value
}

func (v *NullableClusterRest) Set(val *ClusterRest) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRest) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRest(val *ClusterRest) *NullableClusterRest {
	return &NullableClusterRest{value: val, isSet: true}
}

func (v NullableClusterRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


