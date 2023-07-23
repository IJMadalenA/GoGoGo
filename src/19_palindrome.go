package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	var textReverse string

	// normalise the text.
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("The text is a palindrome.")
	} else {
		fmt.Println("The text is not a palindrome.")
	}
}

func main() {
	isPalindrome("kayak")
}
