package keycloak_client

import "math/big"
import "github.com/harisaginting/ginting/pkg/utils/helper"

func getEnv(p string)string{
	return helper.MustGetEnv(p)
}

func decodeBase64BigInt(s string) *big.Int {
	return helper.DecodeBase64BigInt(s)
}

func scope(claims claims, scopes []string) bool {
	for _, search := range scopes {
		for _, permission := range claims.Authorization.Permissions {
			for _, scope := range permission.Scopes {
				if search == scope {
					return true
				}
			}
		}
	}
	return false
}