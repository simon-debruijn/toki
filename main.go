package main

import (
	"fmt"
	"os"
)

func main() {
	defer PanicGracefully()

	if len(os.Args) < 2 {
		fmt.Println(HelpCommands)
		return
	}

	command := CommandFrom(os.Args[1])

	switch command {
	case Commands.DECODE:
		HandleDecode()
	case Commands.VERIFY:
		HandleVerify()
	case Commands.INVALID:
		HandleInvalid()
	}
}
