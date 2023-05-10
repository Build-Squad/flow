/*
 * Access API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Account struct {
	Address string `json:"address"`
	// Flow balance of the account.
	Balance string `json:"balance"`
	Keys []AccountPublicKey `json:"keys,omitempty"`
	Contracts map[string]string `json:"contracts,omitempty"`
	Expandable *AccountExpandable `json:"_expandable"`
	Links *Links `json:"_links,omitempty"`
}
