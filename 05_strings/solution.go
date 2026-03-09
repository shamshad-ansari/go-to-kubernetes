//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPalindrome(s string) bool {
	runes := []rune(strings.ToLower(s))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func CensorWord(text, word string) string {
	runeCount := len([]rune(word))
	replacement := strings.Repeat("*", runeCount)
	return strings.ReplaceAll(text, word, replacement)
}

func main() {
	fmt.Println("=== Reverse String ===")
	fmt.Printf("  %-20s → %s\n", "hello", ReverseString("hello"))
	fmt.Printf("  %-20s → %s\n", "résumé", ReverseString("résumé"))
	fmt.Printf("  %-20s → %s\n", "Hello, 世界", ReverseString("Hello, 世界"))

	fmt.Println("\n=== Is Palindrome ===")
	words := []string{"racecar", "Madam", "hello", "Abba", "A"}
	for _, w := range words {
		fmt.Printf("  %-10s → %t\n", w, IsPalindrome(w))
	}

	fmt.Println("\n=== Censor Word ===")
	fmt.Println(" ", CensorWord("Go is great, Go is fast", "Go"))
	fmt.Println(" ", CensorWord("I love résumé writing", "résumé"))
	fmt.Println(" ", CensorWord("secret secret agent", "secret"))
}
