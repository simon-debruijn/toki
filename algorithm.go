package main

type Algorithm string

var Algorithms = struct {
	INVALID Algorithm
	HS256   Algorithm
}{
	INVALID: "INVALID",
	HS256:   "HS256",
}

func NewAlgorithmFrom(init string) Algorithm {
	if init == "HS256" {
		return Algorithms.HS256
	}

	return Algorithms.INVALID
}
