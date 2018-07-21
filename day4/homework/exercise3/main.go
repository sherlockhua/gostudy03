package main

import (
	"fmt"
)

func addWord(charCount map[rune]int, char rune) {
	
	
		count, ok := charCount[char]
		if !ok {
			charCount[char] = 1
		} else {
			charCount[char] = count+1
		}
	
}

func main() {
	str := "how are   you! you are welcome;我可可"
	var charCount map[rune]int = make(map[rune]int, 10)

	var chars []rune = []rune(str)
	for i := 0; i < len(chars); i++ {
		
			addWord(charCount, chars[i])
		
	}

	for key, val := range charCount {
		fmt.Printf("key:%c val:%d\n", key, val)
	}
}