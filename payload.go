package main

type payload struct {
	encodedContent string
	decodedContent map[string]any
}

type Payload interface {
	EncodedContent() string
	DecodedContent() map[string]any
	ToMarshalable() map[string]any
}

func (p *payload) EncodedContent() string {
	return p.encodedContent
}

func (p *payload) DecodedContent() map[string]any {
	return p.decodedContent
}

func (p *payload) ToMarshalable() map[string]any {
	return p.decodedContent
}

func NewPayload(encodedContent string) Payload {
	decodedContent, err := DecodeContent(encodedContent)

	if err != nil {
		panic("should handle decodeContent error")
	}

	return &payload{
		encodedContent,
		decodedContent,
	}
}
