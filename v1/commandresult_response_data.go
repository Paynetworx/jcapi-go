/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type CommandresultResponseData struct {

	// The stderr output from the command that ran.
	ExitCode int32 `json:"exitCode,omitempty"`

	// The output of the command that was executed.
	Output string `json:"output,omitempty"`
}
