package main

type Command string

var Commands = struct {
	INVALID Command
	DECODE  Command
	VERIFY  Command
}{
	INVALID: "invalid",
	DECODE:  "decode",
	VERIFY:  "verify",
}

func (c Command) Help() string {
	switch c {
	case Commands.DECODE:
		return "toki decode <token>                                  decode a token"
	case Commands.VERIFY:
		return "toki verify <token> --secret <secret>                verify the signature of a token"
	default:
		return ""
	}
}

func CommandFrom(input string) Command {
	if input == "decode" {
		return Commands.DECODE
	}

	if input == "verify" {
		return Commands.VERIFY
	}

	return Commands.INVALID
}
