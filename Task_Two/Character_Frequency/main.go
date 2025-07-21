package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	m := make(map[rune]int)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your String: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)
	word = strings.ToLower(word)

	for _, ch := range word {
		m[ch] += 1
	}

	fmt.Println("Character Frequency: ")
	for k, v := range m {
		fmt.Printf("%q: %d\n", k, v)
	}
}
