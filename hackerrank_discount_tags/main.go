package main

import (
	"fmt"
	"sort"
)

// solution algorithm
// find the smallest odd number
// find the sum
// if the sum is even return sum
// if sum is odd - substract the smallest odd from the sum to make it even


var coupons = []int{2, 3, 6, -5, 10, 1, 1}

//answer is 22

func main() {
	fmt.Printf("coupons: %#v\n", coupons)

	//decending sort
	sort.Slice(coupons, func(i, j int) bool {
		return coupons[i] > coupons[j]
	})

	fmt.Printf("sorted coupons: %#v\n", coupons)

	sum := 0
	smallestODD := coupons[0]
	for i := 0; i < len(coupons); i++ {
		sum = sum + coupons[i]
		if coupons[i] < smallestODD {
			smallestODD = coupons[i]
		}
	}

}
