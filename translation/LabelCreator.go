package translation
import "strconv"

type LabelIndexer struct {
	indexes map[string]int
}

func NewLabelIndexer() LabelIndexer {
	return LabelIndexer{make(map[string]int)}
}

func (labelIndexer *LabelIndexer) AddIndex(label string) string {
	index := labelIndexer.indexes[label]
	labelIndexer.indexes[label]++
	return label + "." + strconv.FormatInt(int64(index), 10)
}