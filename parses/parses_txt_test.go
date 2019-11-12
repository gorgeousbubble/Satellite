package parses

import "testing"

func TestParseTxtSimpleByWords(t *testing.T) {
	file := "../test/data/parses/Alice’s Adventures in Wonderland.txt"
	freq := map[string]int{}
	upateWordsFreq(file, freq)
	reportByWords(freq)
}

func BenchmarkParseTxtSimpleByWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file := "../test/data/parses/Alice’s Adventures in Wonderland.txt"
		freq := map[string]int{}
		upateWordsFreq(file, freq)
		reportByWords(freq)
	}
}

func TestParseTxtSimpleByFreq(t *testing.T) {
	file := "../test/data/parses/Alice’s Adventures in Wonderland.txt"
	freq := map[string]int{}
	upateWordsFreq(file, freq)
	reportByFreq(invertMap(freq))
}

func BenchmarkParseTxtSimpleByFreq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file := "../test/data/parses/Alice’s Adventures in Wonderland.txt"
		freq := map[string]int{}
		upateWordsFreq(file, freq)
		reportByFreq(invertMap(freq))
	}
}
