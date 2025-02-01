package types

type CharFrequency struct {
	Char      string
	Frequency int
}

type HuffmanNode struct {
	Char      string
	Frequency int
	Left      *HuffmanNode
	Right     *HuffmanNode
}
