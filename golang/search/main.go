package main

import "fmt"

func main() {
	res := searchInsert([]int{1, 3, 5, 6}, 4)
	fmt.Println(res)
}

func searchInsert(nums []int, target int) int {
	return middleSearch(nums, 0, len(nums)-1, target)
}

func middleSearch(nums []int, start, end, target int) int {
	if target > nums[end] {
		return end + 1
	} else if target < nums[start] {
		if start-1 < 0 {
			return 0
		}
		return start - 1
	}

	m := (end - start) / 2
	if nums[m] == target {
		return m
	} else if nums[m] > target {
		return middleSearch(nums, 0, m-1, target)
	} else {
		return middleSearch(nums, m+1, end, target)
	}
}
