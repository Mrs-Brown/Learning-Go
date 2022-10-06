// Count frequency of words and return a map of words with their counts. Convert the text to lowercase and trim punctuation
package main

import (
	"fmt"
	"strings"
)

// Function to count frequency of words, trim punctuation, and convert to lowercase
func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text)) //Convert text to lowercase
	frequency := make(map[string]int, len(words))  //Make map
	for _, word := range words {
		word = strings.Trim(word, `.,;"-`) //Trim punctuation from word
		frequency[word]++                  //Add to map of word frequency
	}
	return frequency
}

func main() {
	text := `As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple 
  transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground 
  continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red
  creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear-except the fact of wandering unprovisioned
  and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man.`

	frequency := countWords(text)
	for word, count := range frequency {
		if count > 1 {
			fmt.Printf("%d %v\n", count, word)
		}
	}
}
