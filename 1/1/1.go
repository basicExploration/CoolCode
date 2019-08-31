package main

import "math"

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	value := make(map[float64]int)
	result := make([][]int, 0)       // 初始化数据结构
	for i := 0; i < len(nums); i++ { // 三层数据遍历
		for j := i + 1; j < len(nums); j++ {
			for q := j + 1; q < len(nums); q++ {
				if nums[i]+nums[j]+nums[q] == 0 { //判断数据是否等于0
					resultN := make([]int, 3)
					resultN[0], resultN[1], resultN[2] = nums[i], nums[j], nums[q]
					if Pan(math.Abs(float64(resultN[0]))+math.Abs(float64(resultN[1]))+math.Abs(float64(resultN[2])), value) {
						result = append(result, resultN) // 将数据合成
					}
				}
			}
		}
	}
	return result
}

func Pan(v float64, value map[float64]int) bool {
	if _, ok := value[v]; ok {
		return false
	}
	value[v]++
	return true
}
