package main

import (
	"os"
)

func main() {
	defer PanicGracefully()

	if len(os.Args) < 2 {
		HandleHelp()
		return
	}

	command := NewCommand(os.Args[1])

	switch command {
	case Commands.DECODE:
		HandleDecode()
	case Commands.VERIFY:
		HandleVerify()
	case Commands.HELP:
		HandleHelp()
	case Commands.INVALID:
		HandleInvalid()
	}
}
