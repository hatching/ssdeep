// Copyright (c) 2015, Arbo von Monkiewitsch All rights reserved.
// Copyright (c) 2017, Lukas Rist All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/hatching/ssdeep"
)

var (
	// VERSION is set by the makefile
	VERSION = "v0.0.0"
	// BUILDDATE is set by the makefile
	BUILDDATE = ""
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please provide a file path: ./ssdeep /tmp/file")
		os.Exit(1)
	}

	h1, err := ssdeep.FuzzyFilename(args[0])
	if err != nil && !ssdeep.Force {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(args) == 2 {
		var h2 string
		h2, err = ssdeep.FuzzyFilename(args[1])
		if err != nil && !ssdeep.Force {
			fmt.Println(err)
			os.Exit(1)
		}

		var score int
		score, err = ssdeep.Distance(h1, h2)
		if score != 0 {
			fmt.Printf("%s matches %s (%d)\n", args[0], args[1], score)
		} else if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("The files don't match")
		}
	} else {
		fmt.Println(h1, args[0])
		if err != nil {
			fmt.Println(err)
		}
	}
}
