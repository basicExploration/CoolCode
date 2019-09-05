package main

import "sort"

// 这道题的核心其实就是排序，只有排序了才能看出哪些数据是一样的。

func Three(value []int) (res [][]int) {
	sort.Ints(value) // 为了区分相同的数据，我们要将数据进行排序。
	for i := range value {
    if i > 0 && value[i] = value[i-1] {
      continue
    }
		// 我们使用三个字母来标定三个变量
		l, i, r := 0, i, len(value)
		for l < r {
			values := value[l] + value[i] + value[r]
			switch {
				// switch v := ins.(type) { type 是固定用法，只能是单词type 这里的v是类型，
			case values > 0:
        l++
			case values < 0:
        r --
			default:
        res = append(res, []int{nums[i], nums[l], nums[r]})
        l,r =  next(value,l,r) // 每次都是0的时候避免产生重复数据。
			}
		}
	}

	return result
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}
