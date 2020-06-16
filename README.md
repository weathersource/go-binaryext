# go-binaryext

[![CircleCI](https://circleci.com/gh/weathersource/go-binaryext.svg?style=shield)](https://circleci.com/gh/weathersource/go-binaryext)
[![GoDoc](https://img.shields.io/badge/godoc-ref-blue.svg)](https://godoc.org/github.com/weathersource/go-binaryext)

Package binaryext provides functions that supplement the golang encoding/binary library. In particular, while
encoding/binary provides varint encoding, binaryext provides binary encoding of straight integers, providing
twice the efficiency since all 8 bits in a byte are available for use (rather than 7 with varint).
