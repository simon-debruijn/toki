package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenCanBeMarshaled(t *testing.T) {
	assert := assert.New(t)

	token := NewToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
	json, err := token.MarshalJSON()

	if err != nil {
		t.Fatalf(err.Error())
	}

	// TODO can we check the properties without decoded and encoded having a specific order
	expected := `{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "iat": 1516239022,
    "name": "John Doe",
    "sub": "1234567890"
  },
  "signature": "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}`

	assert.Equal(expected, string(json))
}
