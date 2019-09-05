package main

func majorityElement(nums []int) int {
	result := make(map[int]int)
	for _, v := range nums {
		result[v]++
	}
	for i, v := range result {
		if v > len(nums)/2 {
			return i
		}
	}
	return -1
}
