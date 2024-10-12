package main

import (
	"fmt"
	"strings"
)

type pair struct {
	char  rune
	index int
}

func minRemoveParentheses(s string) string {
	stack := make([]pair, 0)

	for index, character := range s {

		if character == '(' {
			stack = append(stack, pair{
				char:  character,
				index: index,
			})
		} else if character == ')' {
			stackLength := len(stack)
			if stackLength > 0 && stack[stackLength-1].char == '(' {
				stack = stack[:stackLength-1]
			} else {
				stack = append(stack, pair{
					char:  character,
					index: index,
				})
			}
		}
	}

	runeString := []rune(s)
	removeCount := 0
	for _, pair := range stack {
		indexToRemove := pair.index - removeCount
		runeString = append(runeString[:indexToRemove], runeString[indexToRemove+1:]...)
		//as slice is getting shorter the index needs to be adjusted
		removeCount++
	}

	return string(runeString)
}

// Driver code
func main() {
	inputs := []string{"((ar)ab(abc)abd(", "a)rt)lm(ikgh)", "aq)xy())qf(a(ba)q)", "(aw))kk())(w(aa)(bv(wt)r)", "(qi)(kl)((y(yt))(r(q(g)s)"}
	for i, input := range inputs {
		fmt.Printf("%d.\tInput: \"%s\"\n", i+1, input)
		fmt.Printf("\tValid parentheses, after minimum removal: \"%s\"\n", minRemoveParentheses(input))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
