package main

import (
	"fmt"
)

func addWord(wordCount map[string]int, chars []rune) {
	word := string(chars)
	if len(word) > 0 {
		count, ok := wordCount[word]
		if !ok {
			wordCount[word] = 1
		} else {
			wordCount[word] = count+1
		}
	}
}

func main() {
	str := "how are   you! you are welcome"

	var tmp []rune
	var wordCount map[string]int = make(map[string]int, 10)

	var chars []rune = []rune(str)
	for i := 0; i < len(chars); i++ {
		if (str[i] >= 'a' && str[i] <= 'z' || str[i] >= 'A' && str[i] <= 'Z') {
			tmp = append(tmp, chars[i])
		} else {
			addWord(wordCount, tmp)
			tmp = tmp[0:0]
		}
	}

	if len(tmp) > 0 {
		addWord(wordCount, tmp)
	}

	for key, val := range wordCount {
		fmt.Printf("key:%s val:%d\n", key, val)
	}
}