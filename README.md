# NUID - An entropy friendly and highly performant unique identifier generator.

[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)
[![ReportCard](http://goreportcard.com/badge/nats-io/nuid)](http://goreportcard.com/report/nats-io/nuid) [![Build Status](https://travis-ci.org/nats-io/nuid.svg?branch=master)](http://travis-ci.org/nats-io/nuid) [![GoDoc](http://godoc.org/github.com/nats-io/nuid?status.png)](http://godoc.org/github.com/nats-io/nuid) [![Coverage Status](https://coveralls.io/repos/nats-io/nuid/badge.svg?branch=master)](https://coveralls.io/r/nats-io/nuid?branch=master)

## Installation

Use the `go` command:

	$ go get github.com/nats-io/nuid

## Basic Usage

```go

// Utilize the global locked instance
nuid := nuid.Next()

// Create an instance, these are not locked.
n := nuid.New()
nuid = n.Next()

// Generate a new crypto/rand seeded prefix.
// Generally not needed, happens automatically.
n.RandomizePrefix()
```

## License

(The MIT License)

Copyright (c) 2016 Apcera Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
