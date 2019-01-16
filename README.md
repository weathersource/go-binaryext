go-binaryext
============

[![GoDoc](https://godoc.org/github.com/weathersource/go-binaryext?status.svg)](https://godoc.org/github.com/weathersource/go-binaryext)
[![Go Report Card](https://goreportcard.com/badge/github.com/weathersource/go-binaryext)](https://goreportcard.com/report/github.com/weathersource/go-binaryext)
[![Build Status](https://travis-ci.org/weathersource/go-binaryext.svg)](https://travis-ci.org/weathersource/go-binaryext)
[![Codevov](https://codecov.io/gh/weathersource/go-binaryext/branch/master/graphs/badge.svg)](https://codecov.io/gh/weathersource/go-binaryext)

Package binaryext provides functions that supplement the golang encoding/binary library. In particular, while
encoding/binary provides varint encoding, binaryext provides binary encoding of straight integers, providing
twice the efficiency since all 8 bits in a byte are available for use (rather than 7 with varint).
