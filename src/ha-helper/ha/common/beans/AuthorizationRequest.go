package beans

import ()

type AuthorizationRequest struct {
	AuthBaseURL  string
	ClientId     string
	ClientSecret string
	TenantId     string
	
	ClientEmailId string 
	PrivateKeyId string
	PrivateKey string
	Scopes []string 
}
