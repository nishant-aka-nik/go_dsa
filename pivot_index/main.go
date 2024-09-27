package main

import "fmt"

var pivotAray = []int{1, 7, 3, 6, 5, 6}
// var pivotAray = []int{1,2,3}
// var pivotAray = []int{2,1,-1}


func main() {
	totalSum := 0

	for _, value := range pivotAray {
		totalSum += value
	}
	//got total sum in O(n)

	//now we have to iterate over the array
	//sum the elements one by one into beforePivotSum
	// rightArray = totalSum - beforePivotSum 
	// and to skip the last element substract it from above expression
	// rightArray = totalSum - beforePivotSum - pivotAray[i]

	beforePivotSum := 0

	for i := 1; i < len(pivotAray); i++ {
		beforePivotSum += pivotAray[i-1]

		rightArraySum := totalSum - beforePivotSum - pivotAray[i]

		if rightArraySum == beforePivotSum {
			fmt.Printf("Pivot Index is %v", i)
			return
		}
	}

	fmt.Println("No pivot index found")

}
