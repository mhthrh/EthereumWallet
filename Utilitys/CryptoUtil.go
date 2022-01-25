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
	key      string
	FilePath string
	Text     string
	Result   string
}

func NewKey() (*Crypto, *LogInstance) {
	k := new(Crypto)
	k.key = "AnKoloft@~delNazok!12345" // key parameter must be 16, 24 or 32,
	return k, nil
}

func (k *Crypto) Sha256() *LogInstance {
	h := sha256.New()
	h.Write([]byte(k.Text))
	k.Result = fmt.Sprintf("%x", h.Sum(nil))
	return nil
}

func (k *Crypto) Md5Sum() *LogInstance {
	file, err := os.Open(k.FilePath)
	if err != nil {
		Logger("Md5Sum", "Md5 Error", k, err)
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		Logger("Md5Sum", "Md5 Error", k, err)
	}
	k.Result = hex.EncodeToString(hash.Sum(nil))
	return nil
}

func (k *Crypto) Encrypt() *LogInstance {
	key := []byte(k.key)
	plaintext := []byte(k.Text)
	c, err := aes.NewCipher(key)
	if err != nil {
		return Logger("Encrypt", "Md5 Error", k, err)

	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return Logger("Encrypt", "Md5 Error", k, err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return Logger("Encrypt", "Md5 Error", k, err)

	}

	k.Result = Byte64(gcm.Seal(nonce, nonce, plaintext, nil))
	return nil
}

func (k *Crypto) Decrypt() *LogInstance {
	key := []byte(k.key)
	bb, _ := base64.StdEncoding.DecodeString(k.Text)
	ciphertext := []byte(bb)
	c, err := aes.NewCipher(key)
	if err != nil {
		return Logger("Decrypt", "Decrypt Error", k, err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return Logger("Decrypt", "Decrypt Error", k, err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return Logger("Decrypt", "Decrypt Error", k, err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	t, e := gcm.Open(nil, nonce, ciphertext, nil)
	if e != nil {
		return Logger("Decrypt", "Decrypt Error", k, err)
	}
	k.Result = BytesToString(t)
	return nil

}
