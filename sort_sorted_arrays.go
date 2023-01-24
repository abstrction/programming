// sort the array constructed from two sorted arrays, one in decending and the other in ascending order
package main

import "fmt"

func f(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i
		for {
			if j <= 0 || nums[j-1] < nums[j] {
				break
			}

			t := nums[j-1]
			nums[j-1] = nums[j]
			nums[j] = t
			j--
		}
	}
}

func main() {
	var a []int
	_ = []int{2, 9, 7, 4, 6, 5, 8, 1}
	a = []int{2, 9, 7, 5, 1, 4, 6, 8}
	fmt.Println(a)

	f(a)

	fmt.Println(a)
}
