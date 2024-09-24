package main

import (
	"fmt"
	"sort"
)

//wrong way for var initialisation
// var words []string{"eat", "beat", "neat", "tea"}

// learning slice initialisation is different with equal sign after variable
var words = []string{"eat", "beat", "neat", "tea"}

func main() {

	groupMap := make(map[string][]string)

	for _, word := range words {
		wordRune := []rune(word)

		sort.Slice(wordRune, func(i int, j int) bool {
			return wordRune[i] < wordRune[j]
		})

		_, ok := groupMap[string(wordRune)]
		if !ok {
			groupMap[string(wordRune)] = append(groupMap[string(wordRune)], word)
		} else {
			groupMap[string(wordRune)] = append(groupMap[string(wordRune)], word)
		}

		fmt.Printf("word: %#v\n", word)

	}

	for _, group := range groupMap{
		fmt.Println(group)
	}
}

// sort.Slice(runes, func(i, j int) bool {
//     if ascending {
//         return runes[i] < runes[j]
//     }
//     return runes[i] > runes[j]
// })

// Visual Example: Sorting in Ascending Order

// Let’s walk through a simple example to see how true and false affect sorting.

// Original String: "cba"

// Converted to Runes:
// runes := []rune{'c', 'b', 'a'}

// Sorting Process:

// 	1.	First Comparison (i=0, j=1): ‘c’ vs ‘b’
// 	•	Condition: c < b → false
// 	•	Action: Since false, the algorithm determines that ‘c’ should not come before ‘b’.
// 	•	Result: Swap ‘c’ and ‘b’ → ['b', 'c', 'a']
// 	2.	Next Comparison (i=1, j=2): ‘c’ vs ‘a’
// 	•	Condition: c < a → false
// 	•	Action: ‘c’ should not come before ‘a’.
// 	•	Result: Swap ‘c’ and ‘a’ → ['b', 'a', 'c']
// 	3.	Re-compare (i=0, j=1): ‘b’ vs ‘a’
// 	•	Condition: b < a → false
// 	•	Action: ‘b’ should not come before ‘a’.
// 	•	Result: Swap ‘b’ and ‘a’ → ['a', 'b', 'c']
// 	4.	Final State: ['a', 'b', 'c'] → "abc"

// Summary:

// 	•	Each time the comparison returns false, the algorithm checks whether a swap is needed.
// 	•	The algorithm continues comparing and swapping until the entire slice is sorted.

//------------------------------------------------------------------------------------------------------------

// groupMap[string(wordRune)] = append(groupMap[string(wordRune)], word)

// This single line does the following:
// 	•	If string(wordRune) does not exist in groupMap, groupMap[string(wordRune)] returns nil, and append will create a new slice with word as its first element.
// 	•	If string(wordRune) does exist, append adds word to the existing slice.