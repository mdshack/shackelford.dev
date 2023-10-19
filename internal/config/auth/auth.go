package auth

import (
	"crypto/ecdsa"
	"crypto/x509"
	"fmt"
)

type Auth struct {
	AppKey  string `envconfig:"AUTH_APP_KEY"`
	AdminID string `envconfig:"AUTH_ADMIN_ID"`
}

func (a *Auth) GetPrivateKey() *ecdsa.PrivateKey {
	k, err := x509.ParseECPrivateKey([]byte(a.AppKey))

	fmt.Println(err)

	return k
}
