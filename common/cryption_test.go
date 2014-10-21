package common

import (
	"bytes"
	"testing"
)

var PublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MDwwDQYJKoZIhvcNAQEBBQADKwAwKAIhALEsynDW+CmeFceZ8OHMK2wmtc8Cyvuv
cHgEjwCBjvd5AgMBAAE=
-----END PUBLIC KEY-----
`)

func TestRSA(t *testing.T) {
	message := "leverly@126.com"
	data, err := RsaEncrypt(PublicKey, []byte(message))
	if err != nil {
		t.Error("encrypt data failed", err)
	}

	old, err := RsaDecrypt(PrivateKey, data)
	if err != nil {
		t.Error("decrypt data failed", err)
	}

	if !bytes.Equal([]byte(message), old) {
		t.Error("check decrypt result failed")
	}
}

func TestAES(t *testing.T) {
	message := "leverly@126.com"
	sessionKey := []byte("abcdefghijklmnop")
	data, err := AesEncrypt(sessionKey, []byte(message))
	if err != nil {
		t.Error("encrypt data failed", err)
	}
	old, err := AesDecrypt(sessionKey, data)
	if err != nil {
		t.Error("decrypt data failed", err)
	}

	if !bytes.Equal([]byte(message), old) {
		t.Error("check decrypt result failed")
	}
}

func TestDES(t *testing.T) {
	message := "leverly@126.com"
	sessionKey := []byte("abcdefgh")
	data, err := DesEncrypt(sessionKey, []byte(message))
	if err != nil {
		t.Error("encrypt data failed", err)
	}
	old, err := DesDecrypt(sessionKey, data)
	if err != nil {
		t.Error("decrypt data failed", err)
	}

	if !bytes.Equal([]byte(message), old) {
		t.Error("check decrypt result failed")
	}
}

func TestDecryptUsingPem(t *testing.T) {
	old := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	data, err := RsaEncrypt2([]byte("80138512665003396643737838315916663972728479914654754587175091902061894104953"), old)
	if err != nil {
		t.Error("encrypt failed", err)
	}
	old2, err := RsaDecrypt(PrivateKey, data)
	if err != nil {
		t.Error("decrypted failed", err)
	}

	if !bytes.Equal(old, old2) {
		t.Error("check decrypted data failed")
	}
}
