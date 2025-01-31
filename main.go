package main

import (
	compressor "compressionTool/compressor"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: compressionTool <filename>")
		return
	}

	var input io.Reader
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("An error occurred while opening the file: %v", err)
		return
	}
	defer file.Close()
	input = file
	charMap, err := compressor.Compressor(input)
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		return
	}
	for _, freq := range charMap {
		fmt.Printf("%s appears %d times in %s\n", freq.Char, freq.Frequency, fileName)
	}
}
