package main

import (
	"flag"
	"fmt"
	"os"
)

func HandleDecode() {
	if len(os.Args) < 3 {
		panic("the token was not passed")
	}

	token := NewToken(os.Args[2])

	fmt.Println(Prettify(token))
}

func HandleVerify() {
	secret := flag.String("secret", "", "a secret")

	if len(os.Args) < 3 {
		panic("the token was not passed")
	}

	parseErr := flag.CommandLine.Parse(os.Args[3:])

	if parseErr != nil {
		panic(parseErr)
	}

	token := NewToken(os.Args[2])

	if *secret == "" {
		panic("should provide a --secret flag")
	}

	err := token.Verify(*secret)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Signature verified")
}

func HandleHelp() {
	fmt.Println(GetHelpCommands())
}

func HandleInvalid() {
	help := flag.Bool("help", false, "")

	flag.CommandLine.Parse(os.Args[1:])

	if !*help {
		fmt.Println("Unknown command")
	}

	HandleHelp()
}

func GetHelpCommands() string {
	return fmt.Sprintf(`
toki is a tool to decode and verify your jwt tokens.
	
Usage:
	
toki <command> [options]
	
The commands are:
	
	%s
	%s
	%s

Options:
	--secret                The secret to verify the signature of the token
	--help                  Get information about the possible commands
`, Commands.DECODE.Help(), Commands.VERIFY.Help(), Commands.HELP.Help())
}
