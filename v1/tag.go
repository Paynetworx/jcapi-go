/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type Tag struct {

	Id string `json:"_id,omitempty"`

	Expired bool `json:"expired,omitempty"`

	ExternalDN string `json:"externalDN,omitempty"`

	ExternalSourceType string `json:"externalSourceType,omitempty"`

	ExternallyManaged bool `json:"externallyManaged,omitempty"`

	GroupGid string `json:"groupGid,omitempty"`

	GroupName string `json:"groupName,omitempty"`

	// A unique name for the Tag.
	Name string `json:"name,omitempty"`

	RegularExpressions []string `json:"regularExpressions,omitempty"`

	SendToLDAP bool `json:"sendToLDAP,omitempty"`

	// An array of system ids that are associated to the Tag.
	Systems []string `json:"systems,omitempty"`

	// An array of system user ids that are associated to the Tag.
	Systemusers []string `json:"systemusers,omitempty"`
}
