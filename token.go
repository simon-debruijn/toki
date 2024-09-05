package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
)

type myToken struct {
	header    Header
	payload   Payload
	signature Signature
}

func (t *myToken) Header() Header {
	return t.header
}

func (t *myToken) Payload() Payload {
	return t.payload
}

func (t *myToken) Signature() Signature {
	return t.signature
}

func (t *myToken) hash(secret string) string {
	data := fmt.Sprintf("%s.%s", t.header.EncodedContent(), t.payload.EncodedContent())
	h := hmac.New(sha256.New, []byte(secret))

	h.Write([]byte(data))

	hash := h.Sum(nil)

	return Base64URL.EncodeToString(hash[:])
}

func (t *myToken) Verify(secret string) error {
	t.header.Algorithm() // HS256 for now

	hash := t.hash(secret)

	if t.signature.Content() != hash {
		return fmt.Errorf("the signature is not valid")
	}

	return nil
}

type Token interface {
	Header() Header
	Payload() Payload
	Signature() Signature
	Verify(secret string) error
	MarshalJSON() ([]byte, error)
}

func (t *myToken) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(map[string]any{
		"header":    t.header.ToMarshalable(),
		"payload":   t.payload.ToMarshalable(),
		"signature": t.signature.ToMarshalable(),
	}, "", "  ")
}

func NewToken(token string) Token {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		panic(fmt.Errorf("the provided jwt token isn't valid"))
	}
	return &myToken{
		header:    NewHeader((parts[0])),
		payload:   NewPayload(parts[1]),
		signature: NewSignature(parts[2]),
	}
}
