package main

type signature struct {
	content string
}

type Signature interface {
	Content() string
	ToMarshalable() string
}

func (s *signature) Content() string {
	return s.content
}

func (s *signature) ToMarshalable() string {
	return s.Content()
}

func NewSignature(content string) Signature {
	return &signature{
		content,
	}
}
