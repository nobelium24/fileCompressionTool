package comparator

import (
	types "compressionTool/types"
	"sort"
)

func Comparator(charFreqs []types.CharFrequency) {
	sort.Slice(charFreqs, func(i, j int) bool {
		return charFreqs[i].Frequency > charFreqs[j].Frequency
	})
}
