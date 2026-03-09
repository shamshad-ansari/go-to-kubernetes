package main

import (
	"fmt"
)

// ReverseString reverses a string, correctly handling multi-byte Unicode
// characters.
//
// Examples:
//
//	ReverseString("hello")      → "olleh"
//	ReverseString("résumé")     → "émusér"
//	ReverseString("Hello, 世界") → "界世 ,olleH"
func ReverseString(s string) string {
	// TODO: Convert the string to a []rune.
	// TODO: Reverse the rune slice (swap from both ends toward the middle).
	// TODO: Convert back to string and return.

	return "" // replace this
}

// IsPalindrome checks if a string reads the same forwards and backwards,
// ignoring case. It must handle Unicode characters correctly.
//
// Examples:
//
//	IsPalindrome("racecar")  → true
//	IsPalindrome("Madam")    → true
//	IsPalindrome("hello")    → false
//	IsPalindrome("Abba")     → true
func IsPalindrome(s string) bool {
	// TODO: Convert to lowercase for case-insensitive comparison.
	// TODO: Convert to []rune for correct Unicode handling.
	// TODO: Compare characters from both ends moving inward.
	//       If any pair doesn't match, return false.
	// TODO: Return true if all pairs match.

	return false // replace this
}

// CensorWord replaces all occurrences of `word` in `text` with asterisks.
// The number of asterisks should match the rune-length of the word (not byte
// length — this matters for Unicode words).
// Matching is case-sensitive.
//
// Examples:
//
//	CensorWord("Go is great, Go is fast", "Go")     → "** is great, ** is fast"
//	CensorWord("I love résumé writing", "résumé")    → "I love ****** writing"
func CensorWord(text, word string) string {
	// TODO: Count the number of runes in `word` (not bytes!).
	//       Hint: len([]rune(word))
	// TODO: Build a replacement string of that many asterisks.
	//       Hint: strings.Repeat("*", count)
	// TODO: Replace all occurrences of `word` in `text` with the asterisks.
	//       Hint: strings.ReplaceAll(text, word, replacement)
	// TODO: Return the censored text.

	return "" // replace this
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
