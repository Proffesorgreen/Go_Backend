package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your String: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)
	word = strings.ToLower(word)

	left, right := 0, len(word)-1
	for left < right {
		if word[left] != word[right] {
			fmt.Print("Not a Palindrome")
			return
		}
		left += 1
		right -= 1
	}
	fmt.Print("Palindrome")
}
