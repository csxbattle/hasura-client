package hasura_client

import (
	"fmt"
	"net/http"

	"github.com/hasura/go-graphql-client"
)

const (
	AuthModeSecret = iota
	AuthModeJWT
)

// NewClient return gpl client for Hasura. Needs for every request
func NewClient(mode int, addr, token string) *graphql.Client {
	client := graphql.NewClient(addr, nil)
	switch mode {
	case AuthModeSecret:
		client = client.WithRequestModifier(func(r *http.Request) {
			r.Header.Add("X-Hasura-Admin-Secret", token)
		})
	case AuthModeJWT:
		client = client.WithRequestModifier(func(r *http.Request) {
			r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		})
	}
	return client
}
