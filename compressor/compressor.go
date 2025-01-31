package compressor

import (
	"bufio"
	comparator "compressionTool/comparator"
	types "compressionTool/types"
	"io"
)

func Compressor(file io.Reader) ([]types.CharFrequency, error) {
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
	return sortedChar, nil
}
