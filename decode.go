package main

import (
	"encoding/json"
)

func DecodeContent(encodedContent string) (map[string]any, error) {
	decoded, decodeErr := Base64URL.DecodeString(encodedContent)

	if decodeErr != nil {
		return nil, decodeErr
	}

	obj := map[string]any{}
	unmarshalErr := json.Unmarshal(decoded, &obj)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return obj, nil
}
