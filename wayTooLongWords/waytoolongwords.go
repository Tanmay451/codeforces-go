// https://codeforces.com/problemset/problem/71/A
// 71A - Way Too Long Words

package main

import "fmt"

// processWord takes a word as input and processes it according to the specified rules.
// If the length of the word is greater than 10, it shortens the word by replacing
// characters between the first and last characters with the count of removed characters.
func processWord(word string) string {
	if len(word) > 10 {
		return fmt.Sprintf("%c%d%c", word[0], len(word)-2, word[len(word)-1])
	}
	return word
}

func main() {
	var n int

	// Read the number of words to process from the user input.
	fmt.Scanf("%d", &n)

	// Loop through each word and process it using the processWord function.
	for i := 0; i < n; i++ {
		var word string

		// Read each word from the user input.
		fmt.Scan(&word)

		// Process the word and print the result.
		result := processWord(word)
		fmt.Println(result)
	}
}
