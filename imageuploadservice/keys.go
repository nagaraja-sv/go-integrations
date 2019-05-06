package cloudstorageexample

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "Keys/app.rsa"
	pubKeyPath  = "Keys/app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)
var verifyBytes, signBytes []byte

func initKeys() {
	var err error

	signBytes, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		//log.Fatalf("Error reading private key: %v", err)

	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		//log.Fatalf("Error parsing private key: %v", err)
	}
	verifyBytes, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		//log.Fatalf("Error reading public key: %v", err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		//log.Fatalf("Error parsing public key: %v", err)
	}
}
