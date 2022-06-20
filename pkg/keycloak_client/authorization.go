package keycloak_client

import (
	"encoding/json"
	"github.com/cristalhq/jwt/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"github.com/harisaginting/ginting/pkg/utils/helper"
)

type middleware struct {
	env int
}

func Start(env int) middleware { 
	return middleware{env: env}
}

type claims struct {
	jwt.StandardClaims
	Authorization authorization `json:"authorization,omitempty"`
	Username      string        `json:"preferred_username,omitempty"`
	Name          string        `json:"name,omitempty"`
	Email         string        `json:"email,omitempty"`
}

type authorization struct {
	Permissions []permission `json:"permissions,omitempty"`
}

type permission struct {
	RsID   string   `json:"rsid,omitempty"`
	RsName string   `json:"rsname,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

type keycloakJWKDetail struct {
	Key     string   `json:"kty"`
	Kid     string   `json:"kid"`
	Use     string   `json:"sig"`
	Alg     string   `json:"alg"`
	N       string   `json:"n"`
	E       string   `json:"e"`
	X5c     []string `json:"x5c"`
	X5t     string   `json:"x5t"`
	X5tS256 string   `json:"x5t#S256"`
}

type keycloakJWKCerts struct {
	Keys []keycloakJWKDetail `json:"keys"`
}

func (middleware *middleware) Validate(scopes []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		var verifier jwt.Verifier
		var certUrl = getEnv("KEYCLOAK_CERTS")
		var iss 	= helper.MustGetEnv("KEYCLOAK_ISSUER")
		// RETURN IF ENV KEYCLOAK NOT 1 or true
		if middleware.env != 1 {
			context.Next()
			return
		}

		// GET TOKEN
		s := strings.SplitN(context.Request.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			msg := "Authorization token is not found"
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}		
		var headerToken = s[1]
		unverifiedToken, err := jwt.Parse([]byte(headerToken))
		if err != nil {
			msg := err.Error()
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		// VALIDATE TOKEN
		kid   				:= unverifiedToken.Header().KeyID
		certs, jwtKey, err  := getCerts(kid, certUrl)
		if err != nil {
			msg := err.Error()
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		switch certs.Alg {
		case "RS256":
			verifier, err = jwt.NewVerifierRS(jwt.RS256, jwtKey)
		case "RS384":
			verifier, err = jwt.NewVerifierRS(jwt.RS384, jwtKey)
		case "RS512":
			verifier, err = jwt.NewVerifierRS(jwt.RS512, jwtKey)
		default:
			verifier, err = jwt.NewVerifierRS(jwt.RS256, jwtKey)
		}
		if err != nil {
			msg := err.Error()
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}
		
		token, err := jwt.ParseAndVerifyString(headerToken, verifier)
		if err != nil {
			msg := err.Error()
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		var claims claims
		errClaims := json.Unmarshal(token.RawClaims(), &claims)
		if errClaims != nil {
			msg := errClaims.Error()
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		if claims.Issuer != iss {
			msg := "Token issuer is not valid"
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			msg := "Token expired"
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		if !scope(claims, scopes) {
			msg := "Access to this endpoint is not allowed"
			middleware.abort(http.StatusForbidden, context, msg)
			return
		}

		context.Set("keycloak_username", claims.Username)
		context.Set("keycloak_name", claims.Name)
		context.Set("keycloak_email", claims.Email)
		context.Next()
	}
}

func (middleware *middleware) abort(status int, context *gin.Context, message interface{}) {
	context.AbortWithStatusJSON(status, gin.H{
		"status":        status,
		"error_message": message,
		"data":          nil,
	})
}