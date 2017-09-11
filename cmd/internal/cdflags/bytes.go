// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cdflags

import (
	"flag"
	"fmt"
)

// Bytes is a byte size.
//
// Bytes implements flag.Value.
type Bytes int64

var si = []string{"", "k", "M", "G", "T", "P", "E", "Z", "Y"}

// String pretty-prints b using an SI prefix.
func (b Bytes) String() string {
	f := float64(b)
	for i, s := range si {
		if f < 1024 || i == len(si)-1 {
			return fmt.Sprintf("%g%s", f, s)
		}
		f /= 1024
	}
	panic("not reached")
}

// Set parses s into bytes, accepting SI prefixes.
func (b *Bytes) Set(s string) error {
	var num float64
	var unit string
	_, err := fmt.Sscanf(s, "%g%s", &num, &unit)
	if err == nil {
		for _, s := range si {
			if unit == s || unit == s+"B" || unit == s+"iB" {
				*b = Bytes(num)
				return nil
			}
			num *= 1024
		}
	}
	return fmt.Errorf("expected <num> or <num>[%s]", si)
}

// FlagBytes defines a Bytes flag and returns a pointer to the
// variable where the value of the flag will be stored.
func FlagBytes(name string, value Bytes, usage string) *Bytes {
	flag.Var(&value, name, usage)
	return &value
}

// FlagBytesVar defines a Bytes flag. The argument p points to a Bytes
// variable in which to sore the value of the flag.
func FlagBytesVar(p *Bytes, name string, value Bytes, usage string) {
	flag.Var(p, name, usage)
	*p = value
}
