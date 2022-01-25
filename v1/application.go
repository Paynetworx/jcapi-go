/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type Application struct {

	Id string `json:"_id,omitempty"`

	Beta bool `json:"beta,omitempty"`

	Active bool `json:"active,omitempty"`

	Config *ApplicationConfig `json:"config,omitempty"`

	DisplayLabel string `json:"displayLabel,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	LearnMore string `json:"learnMore,omitempty"`

	Name string `json:"name,omitempty"`

	Organization string `json:"organization,omitempty"`

	SsoUrl string `json:"ssoUrl,omitempty"`
}
