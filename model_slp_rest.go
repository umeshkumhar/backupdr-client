/*
 * Backup and DR Service: Management Console API Spec
 *
 * This document defines the API for the Global Manager. All communication is done over HTTPS with UTF-8 encoding. JSON is the only supported format for both request and response payloads. <p></p>Please read <a href=\"https://cloud.google.com/backup-disaster-recovery/docs/api/RestAPIGeneralConcepts.pdf\">Management Console API General concept</a><p></p>To login, use the /session POST API below.<br></br>Then copy the resulting session_id from the output and click on the Authorize button on the top right. Paste the string \"Actifio \" followed by the session id into the form and click Authorize.<p></p>Login is not necessary for reading the rest of this API document. However, login will allow you to try the APIs out within this page.
 *
 * API version: V11.0.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Available SLP (profiles) that this application can potentially use for protection.
type SlpRest struct {
	CloudCredential *CloudCredentialRest `json:"cloudCredential,omitempty"`
	Description string `json:"description,omitempty"`
	Modifydate int64 `json:"modifydate,omitempty"`
	Srcid string `json:"srcid,omitempty"`
	Clusterid string `json:"clusterid,omitempty"`
	Profiletype string `json:"profiletype,omitempty"`
	Performancepool string `json:"performancepool,omitempty"`
	Primarystorage string `json:"primarystorage,omitempty"`
	Remotenode string `json:"remotenode,omitempty"`
	Dedupasyncnode string `json:"dedupasyncnode,omitempty"`
	Vaultpool *DiskPoolRest `json:"vaultpool,omitempty"`
	Vaultpool2 *DiskPoolRest `json:"vaultpool2,omitempty"`
	Vaultpool3 *DiskPoolRest `json:"vaultpool3,omitempty"`
	Vaultpool4 *DiskPoolRest `json:"vaultpool4,omitempty"`
	Cid string `json:"cid,omitempty"`
	Createdate int64 `json:"createdate,omitempty"`
	Localnode string `json:"localnode,omitempty"`
	Orglist []OrganizationRest `json:"orglist,omitempty"`
	Name string `json:"name,omitempty"`
	// URL to access this object
	Href string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale bool `json:"stale,omitempty"`
	// Unique ID for this object
	Id string `json:"id,omitempty"`
}
