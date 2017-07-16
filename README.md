# s3-fuzzer

> 🔐 Command-line AWS S3 Fuzzer

## Why

Recently, a [security researcher][security-researcher] uncovered massive data leaks on AWS S3, like [this][one], [this][two], and [this][three].

## Install

### With Go

```sh
go get -u -v github.com/petermbenjamin/s3-fuzzer
```

### Alternative Method

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

[security-researcher]: https://twitter.com/VickerySec
[one]: http://gizmodo.com/top-defense-contractor-left-sensitive-pentagon-files-on-1795669632
[two]: http://gizmodo.com/gop-data-firm-accidentally-leaks-personal-details-of-ne-1796211612
[three]: https://www.theverge.com/2017/7/12/15962520/verizon-nice-systems-data-breach-exposes-millions-customer-records