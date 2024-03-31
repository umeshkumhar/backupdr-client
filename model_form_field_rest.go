/*
 * Backup and DR Service: Management Console API Spec
 *
 * This document defines the API for the Global Manager. All communication is done over HTTPS with UTF-8 encoding. JSON is the only supported format for both request and response payloads. <p></p>Please read <a href=\"https://cloud.google.com/backup-disaster-recovery/docs/api/RestAPIGeneralConcepts.pdf\">Management Console API General concept</a><p></p>To login, use the /session POST API below.<br></br>Then copy the resulting session_id from the output and click on the Authorize button on the top right. Paste the string \"Actifio \" followed by the session id into the form and click Authorize.<p></p>Login is not necessary for reading the rest of this API document. However, login will allow you to try the APIs out within this page.
 *
 * API version: V11.0.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type FormFieldRest struct {
	Required bool `json:"required,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Header []ChoiceValueRest `json:"header,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	Type_ string `json:"type,omitempty"`
	Size int32 `json:"size,omitempty"`
	Maximum int32 `json:"maximum,omitempty"`
	Minimum int32 `json:"minimum,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Choices []ChoiceValueRest `json:"choices,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Warning string `json:"warning,omitempty"`
	Rows []VolumeSelectionRowRest `json:"rows,omitempty"`
	Values []string `json:"values,omitempty"`
	CurrentValue string `json:"currentValue,omitempty"`
	Validation string `json:"validation,omitempty"`
	Title string `json:"title,omitempty"`
	HelpId string `json:"helpId,omitempty"`
	Optiongroup bool `json:"optiongroup,omitempty"`
	Readonly bool `json:"readonly,omitempty"`
	Modified bool `json:"modified,omitempty"`
	Dynamic bool `json:"dynamic,omitempty"`
	Tags []ChoiceValueRest `json:"tags,omitempty"`
	Networktags []ChoiceValueRest `json:"networktags,omitempty"`
	GetGetchoices string `json:"get_getchoices,omitempty"`
	GetGetDefault string `json:"get_getDefault,omitempty"`
	GetDependent []string `json:"get_dependent,omitempty"`
	Hidden bool `json:"hidden,omitempty"`
	GroupBy bool `json:"groupBy,omitempty"`
	GroupType string `json:"groupType,omitempty"`
	Invalid string `json:"invalid,omitempty"`
	InvalidStr string `json:"invalidStr,omitempty"`
	GetDefault string `json:"get_default,omitempty"`
	Checked bool `json:"checked,omitempty"`
	// Unique ID for this object
	Id string `json:"id,omitempty"`
	// URL to access this object
	Href string `json:"href,omitempty"`
	// When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources.
	Syncdate int64 `json:"syncdate,omitempty"`
	// Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources.
	Stale bool `json:"stale,omitempty"`
}
