// 20:56 Strings  https://www.youtube.com/watch?v=YzLrWHZa-Kc

package main

import ("fmt" 	 // format  (Print)
"strings"
)
var pl = fmt.Println // macro


func main () {
	sV1 := "A word"
	// Replacer that can be used on multiple strings
	// to replace one string with another
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)

	// Get length
	pl("Length : ", len(sV2))

	// Contains string
	pl("Contains Another :", strings.Contains(sV2, "Another"))

	// Get first index match
	pl("o index :", strings.Index(sV2, "o"))

	// Replace all matches with 0
	// If -1 was 2 it would replace the 1st 2 matches
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))

	// Remove whitespace characters from beginning and end of string
	sV3 := "\nSome words\n"
	sV3 = strings.TrimSpace(sV3)
}