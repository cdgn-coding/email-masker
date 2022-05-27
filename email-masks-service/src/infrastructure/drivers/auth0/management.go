package auth0

import "github.com/auth0/go-auth0/management"

func NewManagementClient(domain string, clientID string, clientSecret string) *management.Management {
	m, err := management.New(domain, management.WithClientCredentials(clientID, clientSecret))
	if err != nil {
		panic(err)
	}
	return m
}
