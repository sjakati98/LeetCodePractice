package main

import "fmt"

func bubbleSort(nums []int) {
	/*
		This function takes in an slice of integers
		and returns it in sorted ascending order
	*/

	var i, j int
	for i = 0; i < len(nums); i++ {
		for j = 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				// swap the elements of the slice
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func main() {
	nums := []int{3, 4, 2, 1, 7, 8}
	// testing out pass by reference
	bubbleSort(nums)
	// this will sort the nums slice "in place" in memory
	fmt.Println(nums)
}
