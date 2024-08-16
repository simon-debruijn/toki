# toki
A tool to decode and verify your jwt tokens.

## Installation
For now you can clone the source code and do a `go build` to generate a binary.

## Usage

```
toki <command> [options]

The commands are:

    toki decode <token>                                  decode a token
    toki verify <token> --secret <secret>                verify the signature of a token

Options:
    --secret                The secret to verify the signature of the token
```
