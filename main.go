// MIT License
//
// Copyright (c) 2017 Jewel Mahanta
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// For now, we only need the Files
type Metadata struct {
	Files []map[string]string `json:"files"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Reading metadata.json")
	file, err := ioutil.ReadFile("metadata.json")
	check(err)
	var metadata Metadata
	err = json.Unmarshal(file, &metadata)
	check(err)
	files, err := ioutil.ReadDir(".")
	check(err)
	for _, file := range files {
		for i := range metadata.Files {
			if file.Name() == metadata.Files[i]["original"] {
				log.Println("Renaming: ", metadata.Files[i]["original"], " > ", metadata.Files[i]["real"])
				os.Rename(file.Name(), metadata.Files[i]["real"])
			}
		}
	}
	log.Println("All done! Exiting...")
}
