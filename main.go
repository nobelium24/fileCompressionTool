package main

import (
	"compressionTool/compressor"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: compressionTool <filename>")
		return
	}

	// var input io.Reader
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("An error occurred while opening the file: %v", err)
		return
	}
	defer file.Close()
	// input = file
	charMap, err := compressor.Compressor(file)
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		return
	}
	// for _, freq := range charMap {
	// 	fmt.Printf("%s appears %d times in %s\n", freq., freq.Frequency, fileName)
	// }
	for key, value := range charMap {
		fmt.Printf("Key Value\n")
		fmt.Printf("%s, %s\n", key, value)
	}
}
