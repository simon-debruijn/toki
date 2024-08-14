package main

import "fmt"

type Header interface {
	EncodedContent() string
	DecodedContent() map[string]any
	Algorithm() Algorithm
	ToMarshalable() map[string]any
}

type header struct {
	encodedContent string
	decodedContent map[string]any
}

func (h *header) EncodedContent() string {
	return h.encodedContent
}

func (h *header) DecodedContent() map[string]any {
	return h.decodedContent
}

func (h *header) Algorithm() Algorithm {
	input := h.DecodedContent()["alg"].(string)
	algo := NewAlgorithmFrom(input)

	if algo == Algorithms.INVALID {
		panic(fmt.Sprintf("Algorithm %s is not yet supported", input))
	}

	return algo
}

func (h *header) ToMarshalable() map[string]any {
	return h.DecodedContent()
}

func NewHeader(encodedContent string) Header {
	decodedContent, err := DecodeContent(encodedContent)

	if err != nil {
		panic("should handle decodeContent error")
	}

	return &header{
		encodedContent,
		decodedContent,
	}
}
