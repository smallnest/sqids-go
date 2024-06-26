# Squids Go

[![Go Report Card](https://goreportcard.com/badge/github.com/smallnest/sqids)](https://goreportcard.com/report/github.com/smallnest/sqids)
[![GoDoc](https://godoc.org/github.com/smallnest/sqids?status.svg)](https://godoc.org/github.com/smallnest/sqids)
[![Github Actions](https://img.shields.io/github/actions/workflow/status/smallnest/sqids/tests.yml)](https://github.com/smallnest/sqids/actions)

[Sqids](https://sqids.org/go) (*pronounced "squids"*) is a small library that lets you **generate unique IDs from numbers**. It's good for link shortening, fast & URL-safe ID generation and decoding back into numbers for quicker database lookups.

Fork from [sqids-go](https://github.com/sqids/sqids-go) and refactor:
- refactor `New` to use function options pattern
- refactor `Encode` to use variadic arguments
- add `DecodeOne` to decode one id to one number

Features:

- **Encode multiple numbers** - generate short IDs from one or several non-negative numbers
- **Quick decoding** - easily decode IDs back into numbers
- **Unique IDs** - generate unique IDs by shuffling the alphabet once
- **ID padding** - provide minimum length to make IDs more uniform
- **URL safe** - auto-generated IDs do not contain common profanity
- **Randomized output** - Sequential input provides nonconsecutive IDs
- **Many implementations** - Support for [40+ programming languages](https://sqids.org/)

## 🧰 Use-cases

Good for:

- Generating IDs for public URLs (eg: link shortening)
- Generating IDs for internal systems (eg: event tracking)
- Decoding for quicker database lookups (eg: by primary keys)

Not good for:

- Sensitive data (this is not an encryption library)
- User IDs (can be decoded revealing user count)

## 🚀 Getting started

Use go get.

```bash
go get github.com/smallnest/sqids
```

Then import the package into your own code.

```golang
import "github.com/smallnest/sqids"
```

## 👩‍💻 Examples

> **Note**
> Please note that the following examples omit proper error handling.

Simple encode & decode:

[embedmd]:# (examples/sqids-encode-decode/sqids-encode-decode.go /.+sqids.New/ /\[1, 2, 3\]/)
```go
	s, _ := sqids.New()
	id, _ := s.Encode(1, 2, 3) // "86Rf07"
	numbers := s.Decode(id)              // [1, 2, 3]
```

> **Note**
> 🚧 Because of the algorithm's design, **multiple IDs can decode back into the same sequence of numbers**. If it's important to your design that IDs are canonical, you have to manually re-encode decoded numbers and check that the generated ID matches.

Enforce a *minimum* length for IDs:

[embedmd]:# (examples/sqids-minimum-length/sqids-minimum-length.go /.+sqids.New/ /\[1, 2, 3\]/)
```go
	s, _ := sqids.New(sqids.WithMinLength(10))

	id, _ := s.Encode([]uint64{1, 2, 3}) // "86Rf07xd4z"
	numbers := s.Decode(id)              // [1, 2, 3]
```
Randomize IDs by providing a custom alphabet:

[embedmd]:# (examples/sqids-custom-alphabet/sqids-custom-alphabet.go /.+sqids.New/ /\[1, 2, 3\]/)
```go
	s, _ := sqids.New(sqids.WithAlphabet("FxnXM1kBN6cuhsAvjW3Co7l2RePyY8DwaU04Tzt9fHQrqSVKdpimLGIJOgb5ZE"))

	id, _ := s.Encode(1, 2, 3) // "B4aajs"
	numbers := s.Decode(id)              // [1, 2, 3]
```

Prevent specific words from appearing anywhere in the auto-generated IDs:

[embedmd]:# (examples/sqids-blocklist/sqids-blocklist.go /.+sqids.New/ /\[1, 2, 3\]/)
```go
	s, _ := sqids.New(sqids.WithBlocklist([]string{"86Rf07"}))

	id, _ := s.Encode(1, 2, 3) // "se8ojk"
	numbers := s.Decode(id)              // [1, 2, 3]
```

## 📝 License

[MIT](LICENSE)
