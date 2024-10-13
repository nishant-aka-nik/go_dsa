package main

import (
	"fmt"
)

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	totalWater := 0

	for left < right {
		leftHeight := height[left]
		rightHeight := height[right]
		if leftHeight < rightHeight {
			// Move left pointer
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				totalWater += leftMax - height[left]
			}
			left++
		} else {
			// Move right pointer
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				totalWater += rightMax - height[right]
			}
			right--
		}
	}

	return totalWater
}

func main() {
	// Test cases
	height1 := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println("Output for height1:", trap(height1)) // Expected: 6

	height2 := []int{4,2,0,3,2,5}
	fmt.Println("Output for height2:", trap(height2)) // Expected: 9
}