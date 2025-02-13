package traversal

import (
	types "compressionTool/types"
)

func GenerateCode(node *types.HuffmanNode, code string, codes map[string]string) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		codes[node.Char] = code
	}
	GenerateCode(node.Left, code+"0", codes)
	GenerateCode(node.Right, code+"1", codes)
}
