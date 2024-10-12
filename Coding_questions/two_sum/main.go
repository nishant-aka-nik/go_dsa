package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

func main() {
	sumArray := []int{3, 2, 4}
	sum := 6

	diffComplement := 0

	complementMap := make(map[int]int)
	for idx, value := range sumArray {
		diffComplement = sum - value

		// this is the mistake happened here i am find the complement while iterating over the value
		// the solution would be
		// find the value
		// store the complement

		// _, ok := complementMap[diffComplement]
		// if !ok {
		// 	complementMap[diffComplement] = idx
		// } else {
		// 	fmt.Println("This index for element is ", complementMap[diffComplement], idx)
		// }

		_, ok := complementMap[value]
		if !ok {
			complementMap[diffComplement] = idx
		} else {
			fmt.Println("This index for element is ", complementMap[value], idx)
		}
	}

}
