package main

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	fmt.Println("正常测试结果", ThreeSum([]int{1, -1, 0, 2, 4, -2}))
	fmt.Println("返回nil的结果", ThreeSum([]int{1, 2}))
	fmt.Println("返回空的结果", ThreeSum([]int{1, 1, 1, 1, 1}))
}
