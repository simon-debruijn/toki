package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Prettify(input json.Marshaler) string {
	output, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return ""
	}

	return string(output)
}

func PanicGracefully() {
	if err := recover(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
