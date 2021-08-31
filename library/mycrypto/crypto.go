package mycrypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

type Password string

func GetSalt() []byte {
	length := 16
	b := make([]byte, length)
	_, _ = rand.Read(b)
	return b
}

func (pwd Password) Encrypt(salt []byte) string {
	if salt == nil {
		salt = GetSalt()
	}
	saltHex := hex.EncodeToString(salt)
	enc := pbkdf2.Key([]byte(pwd), salt, 1, 32, sha256.New)
	encHex := hex.EncodeToString(enc)
	ret := strings.Join([]string{saltHex, encHex}, ".")
	return ret
}

func (pwd Password) Check(encPwd string) bool {
	ret := strings.Split(string(encPwd), ".")
	saltHex := ret[0]
	salt, _ := hex.DecodeString(saltHex)
	req := pwd.Encrypt(salt)
	return req == encPwd
}
