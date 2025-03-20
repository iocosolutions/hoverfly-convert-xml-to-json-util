package main

import (
	"fmt"
	"os"
	"strings"

	//xj "github.com/basgys/goxml2json"
	xj "github.com/SpectoLabs/goxml2json"
)

func main() {
	// Check if the filename argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: convertxml <filename>")
		return
	}

	filename := os.Args[1]

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	// xml is an io.Reader
	xml := strings.NewReader(string(content))
	json, err := xj.Convert(xml)
	if err != nil {
		panic("An error occurred:  " + err.Error())
	}
	fmt.Println(json.String())
}
