// Copyright (c) 2015, Arbo von Monkiewitsch All rights reserved.
// Copyright (c) 2017, Lukas Rist All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdeep

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

func ExampleFuzzyFilename() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h, err := FuzzyFile(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(h)
}

func ExampleFuzzyBytes() {
	buffer := make([]byte, 4097)
	rand.Read(buffer)
	h, err := FuzzyBytes(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(h)
}
