package main

import "fmt"

func main() {
	input := [][]int{
		{1, 2, 0, 4},
		{5, 8, 9, 7},
		{3, 0, 8, 6},
		{4, 8, 9, 8},
		{6, 1, 4, 9},
	}

	iPair := make(map[int]struct{})
	jPair := make(map[int]struct{})

	fmt.Println(input)

	lenOfRows := 5
	lenOfColumns := 4

	var i int
	var j int
	for i = 0; i < lenOfRows; i++ {
		for j = 0; j < lenOfColumns; j++ {
			if input[i][j] == 0 {
				iPair[i] = struct{}{}
				jPair[j] = struct{}{}
			}
		}
	}

	//saare zero ko find krke unki indexing ko save kr lo
	// phir dobara run krke row or column k basis pe zero kr do

	fmt.Println(iPair)
	fmt.Println(jPair)

	for i = 0; i < lenOfRows; i++ {
		for j = 0; j < lenOfColumns; j++ {
			_, ok := iPair[i]
			if ok {
				input[i][j] = 0
			} else {
				_, ok := jPair[j]
				if ok {
					input[i][j] = 0
				}
			}
		}
	}

	fmt.Printf("new input: %v", input)
}
