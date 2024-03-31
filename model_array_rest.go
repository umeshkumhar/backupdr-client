/*
 * Backup and DR Service: Management Console API Spec
 *
 * This document defines the API for the Global Manager. All communication is done over HTTPS with UTF-8 encoding. JSON is the only supported format for both request and response payloads. <p></p>Please read <a href=\"https://cloud.google.com/backup-disaster-recovery/docs/api/RestAPIGeneralConcepts.pdf\">Management Console API General concept</a><p></p>To login, use the /session POST API below.<br></br>Then copy the resulting session_id from the output and click on the Authorize button on the top right. Paste the string \"Actifio \" followed by the session id into the form and click Authorize.<p></p>Login is not necessary for reading the rest of this API document. However, login will allow you to try the APIs out within this page.
 *
 * API version: V11.0.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ArrayRest struct {
	Sources []ArrayRest `json:"sources,omitempty"`
	Name string `json:"name,omitempty"`
	Properties []KeyValueRest `json:"properties,omitempty"`
	Srcid int64 `json:"srcid,omitempty"`
	Clusterid int64 `json:"clusterid,omitempty"`
	Modifydate int64 `json:"modifydate,omitempty"`
	Username string `json:"username,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Status string `json:"status,omitempty"`
	Overallstatus string `json:"overallstatus,omitempty"`
	Model string `json:"model,omitempty"`
	Arraytype string `json:"arraytype,omitempty"`
	Orglist []OrganizationRest `json:"orglist,omitempty"`
	Storage []MdiskGroupRest `json:"storage,omitempty"`
	Arraytypelabel string `json:"arraytypelabel,omitempty"`
	Hostcount int32 `json:"hostcount,omitempty"`
	Diskpools []DiskPoolRest `json:"diskpools,omitempty"`
	Reset bool `json:"reset,omitempty"`
	Appliance *ClusterRest `json:"appliance,omitempty"`
	// Unique ID for this object
	Id string `json:"id,omitempty"`
	// URL to access this object
	Href string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale bool `json:"stale,omitempty"`
}
