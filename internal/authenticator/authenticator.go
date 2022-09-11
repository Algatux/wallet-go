package authenticator

import (
	"crypto/rsa"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type RsaKeyPair struct {
	Private string
	Public  string
}

type Auth struct {
	signMethod string
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func (a *Auth) SetRsaKeyPair(keys RsaKeyPair) {
	a.signMethod = "RS512"
	var err error
	a.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(keys.Private))
	if err != nil {
		panic(err)
	}
	a.publicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(keys.Public))
	if err != nil {
		panic(err)
	}
}

func (a *Auth) AuthenticateCredentials(username string, password string) (string, error) {
	return a.generateToken(username)
}

func (a *Auth) ValidateToken(token string) error {
	if true == a.verifyToken(token) {
		return nil
	}

	return errors.New("invalid credentials")
}

func (a *Auth) AuthenticateToken(token string) (jwt.RegisteredClaims, error) {
	tokenData := jwt.RegisteredClaims{}
	if true == a.verifyToken(token) {
		return tokenData, nil
	}

	return tokenData, errors.New("invalid credentials")
}

func (a *Auth) generateToken(username string) (string, error) {
	exp := time.Now().Unix() + 3600

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		Issuer:    "wallet",
		Subject:   username,
	}

	signMethod := jwt.GetSigningMethod(a.signMethod)
	token := jwt.NewWithClaims(signMethod, claims)
	ss, err := token.SignedString(a.privateKey)

	return ss, err
}

func (a *Auth) verifyToken(token string) bool {
	signMethod := jwt.GetSigningMethod(a.signMethod)
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return false
	}

	err := signMethod.Verify(strings.Join(parts[0:2], "."), parts[2], a.publicKey)

	return err == nil
}
