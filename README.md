# s3-fuzzer

> 🔐 Command-line AWS S3 Fuzzer

## Why

Because of [this](https://github.com/petermbenjamin/YAS3BL).

## Install

### With Go

```sh
go get -u -v github.com/petermbenjamin/s3-fuzzer
```

### Alternative Method

Download directly from [Releases](https://github.com/petermbenjamin/s3-fuzzer/releases)

## Usage

```sh
$ s3-fuzzer --help
usage: s3-fuzzer [<flags>]

Flags:
      --help       Show context-sensitive help (also try --help-long and --help-man).
  -f, --file=FILE  Path to file.
  -n, --name=NAME  Bucket name.
  -o, --open       Open URL in default browser.
  -d, --debug      Debug mode.
```

## License

MIT &copy; 2017 [Peter Benjamin](https://github.com/petermbenjamin)
