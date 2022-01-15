package Utilitys

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

type CryptoInterface interface {
	Decrypt()
	Encrypt()
	Sha256()
	Md5Sum()
}
type Crypto struct {
	key        string
	FilePath   string
	Text       string
	Result     string
	exceptions *[]Exceptions
	Status     *Exceptions
}

func NewKey() *Crypto {
	k := new(Crypto)
	k.key = "AnKoloft@~delNazok!12345" // key parameter must be 16, 24 or 32,
	return k
}

func (k *Crypto) Sha256() {
	h := sha256.New()
	h.Write([]byte(k.Text))
	k.Result = fmt.Sprintf("%x", h.Sum(nil))
	k.Status = SelectException(0, k.exceptions)
}

func (k *Crypto) Md5Sum() {
	file, err := os.Open(k.FilePath)
	if err != nil {
		k.Status = SelectException(1000, k.exceptions)
		return
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		k.Status = SelectException(1000, k.exceptions)
		return
	}
	k.Result = hex.EncodeToString(hash.Sum(nil))
}

func (k *Crypto) Encrypt() {
	key := []byte(k.key)
	plaintext := []byte(k.Text)
	c, err := aes.NewCipher(key)
	if err != nil {
		k.Status = SelectException(1000, k.exceptions)
		return
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		k.Status = SelectException(1000, k.exceptions)
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		k.Status = SelectException(1000, k.exceptions)
		return
	}

	k.Result = Byteto64(gcm.Seal(nonce, nonce, plaintext, nil))
	k.Status = SelectException(1000, k.exceptions)
}

func (k *Crypto) Decrypt() {
	key := []byte(k.key)
	bb, _ := base64.StdEncoding.DecodeString(k.Text)
	ciphertext := []byte(bb)
	c, err := aes.NewCipher(key)
	if err != nil {
		k.Status = SelectException(1000, k.exceptions)
		return
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		k.Status = SelectException(1000, k.exceptions)
		return
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		k.Status = SelectException(1000, k.exceptions)
		return
		//return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	t, e := gcm.Open(nil, nonce, ciphertext, nil)
	if e != nil {
		k.Status = SelectException(1000, k.exceptions)
		return
	}
	k.Result = BytesToString(t)
	k.Status = SelectException(0, k.exceptions)

}
