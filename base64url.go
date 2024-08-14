package main

import "encoding/base64"

type base64url struct{}

var Base64URL base64url

func (base64url) EncodeToString(src []byte) string {
	return base64.RawURLEncoding.EncodeToString(src)
}

func (base64url) DecodeString(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}
