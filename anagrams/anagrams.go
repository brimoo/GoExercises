package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a word: ")
	word, _ := reader.ReadString('\n')
	word = word[:len(word)-1]

	words := make([]string, 0)

	fmt.Print("Enter word to check(enter 0 to finish): ")

	temp, _ := reader.ReadString('\n')
	temp = temp[:len(temp)-1]

	for temp != "0" {
		words = append(words, temp)
		fmt.Print("Enter word to check(enter 0 to finish): ")
		temp, _ = reader.ReadString('\n')
		temp = temp[:len(temp)-1]
	}

	anagrams(word, words)

}

func anagrams(word string, words []string) {
	fmt.Println("--Anagrams--")
	for i := 0; i < len(words); i++ {
		if isAnagram(word, words[i]) {
			fmt.Println(words[i])
		}
	}
}

func isAnagram(word, toCheck string) bool {
	if len(word) != len(toCheck) {
		return false
	}

	for i := 0; i < len(word); i++ {
		if !strings.Contains(word, string(toCheck[i])) {
			return false
		}
	}

	return true
}
