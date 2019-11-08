package parses

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func parseCmdFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name)
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

func upateWordsFreq(name string, freq map[string]int) {
	var file *os.File
	var err error
	if file, err = os.Open(name); err != nil {
		log.Println("Error open the file:", err)
		return
	}
	defer file.Close()
	readWordsFreq(bufio.NewReader(file), freq)
}

func readWordsFreq(reader *bufio.Reader, freq map[string]int) {
	for {
		line, err := reader.ReadString('\n')
		for _, word := range splitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				freq[strings.ToLower(word)]++
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Println("Error finish reading the file:", err)
			}
			break
		}
	}
}

func splitOnNonLetters(s string) []string {
	noALetter := func(char rune) bool {
		return !unicode.IsLetter(char)
	}
	return strings.FieldsFunc(s, noALetter)
}

func invertMap(nor map[string]int) map[int][]string {
	inv := make(map[int][]string, len(nor))
	for k, v := range nor {
		inv[v] = append(inv[v], k)
	}
	return inv
}

func reportByWords(freq map[string]int) {
	words := make([]string, 0, len(freq))
	wordWidth, freqWidth := 0, 0
	for k, v := range freq {
		words = append(words, k)
		if width := utf8.RuneCountInString(k); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(v)); width > freqWidth {
			freqWidth = width
		}
	}
	sort.Strings(words)
	gap := wordWidth + freqWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, freqWidth, freq[word])
	}
}

func reportByFreq(freq map[int][]string) {
	freqs := make([]int, 0, len(freq))
	for v := range freq {
		freqs = append(freqs, v)
	}
	sort.Ints(freqs)
	width := len(fmt.Sprint(freqs[len(freqs)-1]))
	fmt.Println("Frequency → Words")
	for _, v := range freqs {
		words := freq[v]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, v, strings.Join(words, ", "))
	}
}
