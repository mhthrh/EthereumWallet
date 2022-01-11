package Utilitys

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func Sha256(a string) string {
	h := sha256.New()
	h.Write([]byte(a))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Md5Sum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

type CryptoInterface interface {
	Decrypt(string) (string, error)
	Encrypt(string) (string, error)
}

//KeyStr lll
type KeyStr struct {
	strkey string
}

//NewKey ddd
func NewKey() CryptoInterface {
	key := new(KeyStr)
	key.strkey = "AnKoloft@~delNazok!12345" // key parameter must be 16, 24 or 32,
	return key
}

func (k *KeyStr) Encrypt(text string) (string, error) {
	key := []byte(k.strkey)
	plaintext := []byte(text)
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	return Byteto64(gcm.Seal(nonce, nonce, plaintext, nil)), nil
}

func (k *KeyStr) Decrypt(text string) (string, error) {
	key := []byte(k.strkey)
	bb, _ := base64.StdEncoding.DecodeString(text)
	ciphertext := []byte(bb)
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	t, e := gcm.Open(nil, nonce, ciphertext, nil)
	return BytesToString(t), e
}
