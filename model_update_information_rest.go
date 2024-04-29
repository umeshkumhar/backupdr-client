/*
 * Backup and DR Service: Management Console API Spec
 *
 * This document defines the API for the Global Manager. All communication is done over HTTPS with UTF-8 encoding. JSON is the only supported format for both request and response payloads. <p></p>Please read <a href=\"https://cloud.google.com/backup-disaster-recovery/docs/api/RestAPIGeneralConcepts.pdf\">Management Console API General concept</a><p></p>To login, use the /session POST API below.<br></br>Then copy the resulting session_id from the output and click on the Authorize button on the top right. Paste the string \"Actifio \" followed by the session id into the form and click Authorize.<p></p>Login is not necessary for reading the rest of this API document. However, login will allow you to try the APIs out within this page.
 *
 * API version: V11.0.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateInformationRest struct {
	Updatetype string `json:"updatetype,omitempty"`
	Componentname string `json:"componentname,omitempty"`
	Updateid string `json:"updateid,omitempty"`
	Readmelink string `json:"readmelink,omitempty"`
	Expectedduration int64 `json:"expectedduration,omitempty"`
	Releasedate int64 `json:"releasedate,omitempty"`
	Duedate int64 `json:"duedate,omitempty"`
	Updatename string `json:"updatename,omitempty"`
}
