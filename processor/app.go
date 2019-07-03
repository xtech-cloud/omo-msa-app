package processor

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"time"
)

func CreateApp(_appname string) (string, string, string, string, error) {
	if _appname == "" {
		return "", "", "", "", errors.New("appname is nil or empty")
	}

	now := time.Now().Unix()
	keyCode := fmt.Sprintf("%v-%v-key", _appname, now)
	secretCode := fmt.Sprintf("%v-%v-secret", _appname, now)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(keyCode))
	appKey := hex.EncodeToString(md5Ctx.Sum(nil))
	md5Ctx.Write([]byte(secretCode))
	appSecret := hex.EncodeToString(md5Ctx.Sum(nil))
	publicKey, privateKey, err := rsaGenerateKey()

	return appKey, appSecret, string(publicKey), string(privateKey), err
}

func ResetSecret(_appname string) string {
	now := time.Now().Unix()
	secretCode := fmt.Sprintf("%v-%v-secret", _appname, now)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(secretCode))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func ResetKey(_appname string) (string, string, error) {
	publicKey, privateKey, err := rsaGenerateKey()
	return string(publicKey), string(privateKey), err
}

//return publickey, privatekey, error
func rsaGenerateKey() ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if nil != err {
		return nil, nil, err
	}
	derStram := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStram,
	}
	privateKeyBytes := pem.EncodeToMemory(block)

	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if nil != err {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}
	publicKeyBytes := pem.EncodeToMemory(block)

	return publicKeyBytes, privateKeyBytes, err
}
