/*
 * Backup and DR Service: Management Console API Spec
 *
 * This document defines the API for the Global Manager. All communication is done over HTTPS with UTF-8 encoding. JSON is the only supported format for both request and response payloads. <p></p>Please read <a href=\"https://cloud.google.com/backup-disaster-recovery/docs/api/RestAPIGeneralConcepts.pdf\">Management Console API General concept</a><p></p>To login, use the /session POST API below.<br></br>Then copy the resulting session_id from the output and click on the Authorize button on the top right. Paste the string \"Actifio \" followed by the session id into the form and click Authorize.<p></p>Login is not necessary for reading the rest of this API document. However, login will allow you to try the APIs out within this page.
 *
 * API version: V11.0.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Appliance where the application resides.
type ClusterRest struct {
	SharedSecret string `json:"shared_secret,omitempty"`
	Serviceaccount string `json:"serviceaccount,omitempty"`
	Zone string `json:"zone,omitempty"`
	Region string `json:"region,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Password string `json:"password,omitempty"`
	Version string `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	Type_ string `json:"type,omitempty"`
	Protocol int32 `json:"protocol,omitempty"`
	Datacenter string `json:"datacenter,omitempty"`
	Masterid string `json:"masterid,omitempty"`
	Clusterid string `json:"clusterid,omitempty"`
	Importstatus string `json:"importstatus,omitempty"`
	Username string `json:"username,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	Clusterstatus *ClusterStatusRest `json:"clusterstatus,omitempty"`
	Lastsync int64 `json:"lastsync,omitempty"`
	Publicip string `json:"publicip,omitempty"`
	Secureconnect bool `json:"secureconnect,omitempty"`
	PkiBootstrapped bool `json:"pkiBootstrapped,omitempty"`
	Clusterlist []ClusterRest `json:"clusterlist,omitempty"`
	CallhomeInfo *CallhomeInfoRest `json:"callhomeInfo,omitempty"`
	Rmipaddress []string `json:"rmipaddress,omitempty"`
	// Shows the support status of cluster
	Supportstatus string `json:"supportstatus,omitempty"`
	// Unique ID for this object
	Id string `json:"id,omitempty"`
	// URL to access this object
	Href string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale bool `json:"stale,omitempty"`
}