package main

import (
	"fmt"
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string) // Hashmap for count
	for _, str := range strs {
		// A place for every single alphabet in our string
		count := make([]int, 26)
		// We will compute the frequency for every string
		for _, chr := range str {
			// Calculating the value from 1 to 26 for the alphabet
			index := chr - 'a'
			count[index] += 1 // Increasing its frequency in the hashmap
		}
		// Each element in this tuple represents the frequency of an
		// English letter in the corresponding title
		key := ""
		// mapping a string to key. This will generate indentical
		// vectors for strings that are anagrams
		for i := 0; i < 26; i++ {
			key += "#"
			key += strconv.Itoa(count[i])
		}
		if _, ok := res[key]; ok {
			res[key] = append(res[key], str)
		} else {
			// We add the string as an anagram if it matched the content
			// of our res hashmap
			res[key] = []string{str}
		}
	}
	group := make([][]string, 0)
	for _, value := range res {
		group = append(group, value)
	}
	return group
}

/*
	stringArrayToString() is used to convert a string array to

string. It is used for printing purposes in driver code.
*/
func stringArrayToString(str []string) string {
	res := "["
	for j, s := range str {
		res += "\"" + s + "\""
		if j < len(str)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}

/*
	doublyStringArrayToString() is used to convert a 2D string array to

string. It is used for printing purposes in driver code.
*/
func doublyStringArrayToString(str [][]string) string {
	res := "["
	for i, st := range str {
		res += "["
		for j, s := range st {
			res += "\"" + s + "\""
			if j < len(st)-1 {
				res += ", "
			}
		}
		res += "]"
		if i < len(str)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}

// main() is used for the driver code
func main() {
	titles := [][]string{{"eat", "beat", "neat", "tea"},
		{"duel", "dule", "speed", "spede", "deul", "cars"},
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""}, {"sword", "swords"}, {"pot", "top", "opt"}}
	for i, title := range titles {
		fmt.Printf("%d.\tThe Grouped Anagrams for the array %s%s\n", i+1, stringArrayToString(title), " are:")
		fmt.Printf("\t%s\n", doublyStringArrayToString(groupAnagrams(title)))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
