package keycloak_client

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
)

func getCerts(kid string, keysUrl string) (keycloakJWKDetail, *rsa.PublicKey, error) {
	var certs 		keycloakJWKDetail
	var jwk 		keycloakJWKCerts
	keysRequest, err := http.NewRequest("GET", keysUrl, nil)
	if err != nil {
		return certs, nil, err
	}
	keysResponse, err := http.DefaultClient.Do(keysRequest)
	if err != nil {
		return certs, nil, err
	}
	keysResponseBody, err := ioutil.ReadAll(keysResponse.Body)
	if err != nil {
		return certs, nil, err
	}

	err = json.Unmarshal([]byte(keysResponseBody), &jwk)
	if err != nil {
		return certs, nil, err
	}

	var n *big.Int
	var e int
	for _, key := range jwk.Keys {
		if key.Kid == kid {
			certs = key
			n = decodeBase64BigInt(key.N)
			e = int(decodeBase64BigInt(key.E).Int64())
			break
		}
	}

	if n == nil || e == 0 {
		return certs, nil, err
	}

	jwtKey := &rsa.PublicKey{
		N: n,
		E: e,
	}
	return certs, jwtKey, nil
}