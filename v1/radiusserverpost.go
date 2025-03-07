/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type Radiusserverpost struct {

	Mfa string `json:"mfa,omitempty"`

	Name string `json:"name"`

	NetworkSourceIp string `json:"networkSourceIp"`

	// RADIUS shared secret between the server and client.
	SharedSecret string `json:"sharedSecret"`

	TagNames []string `json:"tagNames,omitempty"`

	UserLockoutAction string `json:"userLockoutAction,omitempty"`

	UserPasswordExpirationAction string `json:"userPasswordExpirationAction,omitempty"`
}
