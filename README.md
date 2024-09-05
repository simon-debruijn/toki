# toki
A tool to decode and verify your jwt tokens.

## Installation
For now you can clone the source code and do a `go build` to generate a binary.

```
git clone https://github.com/simon-debruijn/toki.git
cd ./toki
go build
```

After this you can add the binary ./toki to your path.

## Usage

```
toki <command> [options]

The commands are:

    toki decode <token>                                  decode a token
    toki verify <token> --secret <secret>                verify the signature of a token

Options:
    --secret                The secret to verify the signature of the token
```
