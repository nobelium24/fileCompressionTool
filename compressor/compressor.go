package compressor

import (
	"bufio"
	comparator "compressionTool/comparator"
	priorityQueue "compressionTool/priorityQueue"
	types "compressionTool/types"
	"container/heap"
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
	pq := &priorityQueue.PriorityQueue{}
	heap.Init(pq)
	for _, charFreq := range sortedChar {
		node := &types.CharFrequency{
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
	// root := heap.Pop(pq).(*types.HuffmanNode)
	return sortedChar, nil
}
