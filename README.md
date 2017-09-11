# rabin [![GoDoc](https://godoc.org/github.com/aclements/go-rabin/rabin?status.svg)](https://godoc.org/github.com/aclements/go-rabin/rabin) [![Go Report Card](https://goreportcard.com/badge/github.com/aclements/go-rabin)](https://goreportcard.com/report/github.com/aclements/go-rabin)

The rabin package implements Rabin hashing (aka fingerprinting) and
content-defined chunking based on Rabin hashing.

Rabin hashing has the unusual property that it can efficiently compute
a "rolling hash" of a stream of data, where the hash value reflects
only the most recent w bytes of the stream, for some window size w.
This property makes it ideal for "content-defined chunking", which
sub-divides sequential data on boundaries that are robust to
insertions and deletions.

The details of Rabin fingerprinting are described in Rabin, Michael
(1981). "Fingerprinting by Random Polynomials." Center for Research in
Computing Technology, Harvard University. Tech Report TR-CSE-03-01.

Installation
------------

To download go-rabin, run

```sh
go get -d -u github.com/aclements/go-rabin/rabin
```

You can then import this package into your projects with

```go
import "github.com/aclements/go-rabin/rabin"
```

Demos
-----

There is a small program in `cmd/cdsplit` that divides an input file
into content-defined chunks. It's only intended as a demo, but can be
installed using

```sh
go get github.com/aclements/go-rabin/cmd/cdsplit
```
