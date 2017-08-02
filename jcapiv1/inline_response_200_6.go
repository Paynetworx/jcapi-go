/* 
 * JumpCloud APIs
 *
 * V1 and V2 versions of JumpCloud's API. The previous version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * OpenAPI spec version: 1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package jcapiv1

type InlineResponse2006 struct {

	Organization string `json:"organization,omitempty"`

	Created string `json:"created,omitempty"`

	LastContact string `json:"lastContact,omitempty"`

	Os string `json:"os,omitempty"`

	Version string `json:"version,omitempty"`

	Arch string `json:"arch,omitempty"`

	NetworkInterfaces []InlineResponse2005NetworkInterfaces `json:"networkInterfaces,omitempty"`

	Hostname string `json:"hostname,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	SystemTimezone int32 `json:"systemTimezone,omitempty"`

	TemplateName string `json:"templateName,omitempty"`

	RemoteIP string `json:"remoteIP,omitempty"`

	Active bool `json:"active,omitempty"`

	SshdParams []interface{} `json:"sshdParams,omitempty"`

	AllowSshPasswordAuthentication bool `json:"allowSshPasswordAuthentication,omitempty"`

	AllowSshRootLogin bool `json:"allowSshRootLogin,omitempty"`

	AllowMultiFactorAuthentication bool `json:"allowMultiFactorAuthentication,omitempty"`

	AllowPublicKeyAuthentication bool `json:"allowPublicKeyAuthentication,omitempty"`

	ModifySSHDConfig bool `json:"modifySSHDConfig,omitempty"`

	AgentVersion string `json:"agentVersion,omitempty"`

	ConnectionHistory []interface{} `json:"connectionHistory,omitempty"`

	SshRootEnabled bool `json:"sshRootEnabled,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Id string `json:"id,omitempty"`

	Id string `json:"_id,omitempty"`
}
