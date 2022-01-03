package Utilitys

import (
	"bytes"
	"encoding/base64"
)

func Byteto64(msg []byte) string {
	return base64.StdEncoding.EncodeToString(msg)
}

func BytesToString(b []byte) string {

	return bytes.NewBuffer(b).String()
}
