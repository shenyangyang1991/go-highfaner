package utils

import "testing"

func TestMD5CreateStrings(t *testing.T) {
	t.Log(MD5CreateStrings("123456"))
}

func TestAesEncrypt(t *testing.T) {
	ciphertext, err := AesEncrypt([]byte("1"), []byte("highfaner1234567"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ciphertext)
}
