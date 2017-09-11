// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command cdsplit divides a file into variable-sized, content-defined
// chunks that are robust to insertions, deletions, and changes to the
// input file.
//
// It is a demo for the Go rabin package.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aclements/go-rabin/cmd/internal/cdflags"
	"github.com/aclements/go-rabin/rabin"
)

func main() {
	// Parse and validate flags
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [flags] in-file\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Divide in-file into variable-sized, content-defined chunks that are robust to\n")
		fmt.Fprintf(os.Stderr, "insertions, deletions, and changes to in-file.\n\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	window := cdflags.FlagBytes("window", 64, "use a rolling hash with window size `w`")
	avg := cdflags.FlagBytes("avg", 4<<10, "average chunk `size`; must be a power of 2")
	min := cdflags.FlagBytes("min", 512, "minimum chunk `size`")
	max := cdflags.FlagBytes("max", 32<<10, "maximum chunk `size`")
	outBase := flag.String("out", "", "write output to `base`.NNNNNN")
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
	}
	if *min > *max {
		log.Fatal("-min must be <= -max")
	}
	if *avg&(*avg-1) != 0 {
		log.Fatal("-avg must be a power of two")
	}
	if *min < *window {
		log.Fatal("-min must be >= -window")
	}
	inFile := flag.Arg(0)
	if *outBase == "" {
		*outBase = inFile
	}

	// Open input file
	f, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Chunk and write output files.
	copy := new(bytes.Buffer)
	r := io.TeeReader(f, copy)
	c := rabin.NewChunker(rabin.NewTable(rabin.Poly64, int(*window)), r, int(*min), int(*avg), int(*max))
	for i := 0; ; i++ {
		clen, err := c.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		name := fmt.Sprintf("%s.%06d", *outBase, i)
		fOut, err := os.Create(name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(fOut, copy, int64(clen))
		if err == nil {
			err = fOut.Close()
		}
		if err != nil {
			log.Fatalf("error writing %s: %s", name, err)
		}
	}
}
