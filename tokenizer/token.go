package tokenizer

import "github.com/wangbin/jiebago"

var seg jiebago.Segmenter

func init() {
	seg.LoadDictionary("dict.txt")
}

func Cut_word(dsc string) []string{
	var	words []string
	for word := range seg.CutForSearch(dsc, true) {
		words = append(words, word)
	}
	return words
}

func TF_cal(word string, dscList []string) float32 {
	var cnt float32 = 0
	for _, e := range dscList {
		if e == word {
			cnt++
		}
	}
	return cnt / (float32)(len(dscList))
}

