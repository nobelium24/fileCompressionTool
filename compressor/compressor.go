package compressor

import (
	"bufio"
	comparator "compressionTool/comparator"
	priorityQueue "compressionTool/priorityQueue"
	"compressionTool/traversal"
	types "compressionTool/types"
	"container/heap"
	"io"
	"os"
)

// func Compressor(file io.Reader) ([]types.CharFrequency, error) {

func Compressor(file io.ReadSeeker) (map[string]string, error) {
	frequencyMap := make(map[string]int)
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanRunes)
	for input.Scan() {
		rune := input.Text()
		frequencyMap[rune]++
	}
	if err := input.Err(); err != nil {
		return nil, err
	}
	var sortedChar []types.CharFrequency
	for char, freq := range frequencyMap {
		sortedChar = append(sortedChar, types.CharFrequency{
			Char:      char,
			Frequency: freq,
		})
	}
	comparator.Comparator(sortedChar)
	pq := &priorityQueue.PriorityQueue{}
	heap.Init(pq)
	for _, charFreq := range sortedChar {
		node := &types.HuffmanNode{
			Char:      charFreq.Char,
			Frequency: charFreq.Frequency,
		}
		heap.Push(pq, node)
	}

	for pq.Len() > 1 {
		left := heap.Pop(pq).(*types.HuffmanNode)
		right := heap.Pop(pq).(*types.HuffmanNode)

		newNode := &types.HuffmanNode{
			Frequency: left.Frequency + right.Frequency,
			Left:      left,
			Right:     right,
		}
		heap.Push(pq, newNode)
	}
	root := heap.Pop(pq).(*types.HuffmanNode)
	codes := make(map[string]string)
	traversal.GenerateCode(root, "", codes)

	file.Seek(0, io.SeekStart) // Reset the input reader to the beginning
	var encodedText string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		char := scanner.Text()
		encodedText += codes[char]
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Write the encoded text to an output file
	outputFile, err := os.Create("encoded_output.txt")
	if err != nil {
		return nil, err
	}
	defer outputFile.Close()
	_, err = outputFile.WriteString(encodedText)
	if err != nil {
		return nil, err
	}

	return codes, nil

	// return sortedChar, nil
}
