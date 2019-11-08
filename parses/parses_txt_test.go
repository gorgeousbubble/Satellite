package parses

import "testing"

func TestParseTxtSimple(t *testing.T) {
	file := "../test/data/parses/Alice’s Adventures in Wonderland.txt"
	freq := map[string]int{}
	upateWordsFreq(file, freq)
	reportByWords(freq)
}
