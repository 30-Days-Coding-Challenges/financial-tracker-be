package utils

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"log"
	"math/rand"
	"time"

	"os"
)

func selectKey(key string) string {
	if key != "alpha" {
		return os.Getenv("SECRETKEY_BETA")
	} else {
		return os.Getenv("SECRETKEY_ALPHA")
	}
}

func DigestStringUsingMD5(str string) string {

	digestAuth := []byte(str)
	hash := md5.Sum(digestAuth)

	return hex.EncodeToString(hash[:])
}

func EncryptStringToBase64(str string, key string) (string, error) {

	cipher, err := aes.NewCipher([]byte(selectKey(key)))
	if err != nil {
		log.Fatal("error encrypt:", err)
	}

	out := make([]byte, len(str))

	cipher.Encrypt(out, []byte(str))

	hex := hex.EncodeToString(out)

	base64Str := base64.StdEncoding.EncodeToString([]byte(hex))

	return base64Str, err
}

func DecryptBase64ToString(str string, key string) (string, error) {
	hexStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal("error decrypt:", err)
	}

	ciphertext, _ := hex.DecodeString(string(hexStr))

	c, err := aes.NewCipher([]byte(selectKey(key)))

	if err != nil {
		log.Fatal(err)
	}

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	decryptedString := string(pt[:])

	return decryptedString, err
}

func GenerateRandomStrings(length int) string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
