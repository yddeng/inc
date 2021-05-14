package inc

import "github.com/yddeng/utils/strutil"

func ReadWords(line string) ([]string, int) {
	words := strutil.Str2Slice(line)
	wordsLen := len(words)
	return words, wordsLen
}
