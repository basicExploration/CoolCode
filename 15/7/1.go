package main

var (
	goods = []int{1,1,1,1,1,13,7}
	maxGoods = 27
)

// 动态规划解决 背包问题。背包问题的分解是每次的添加与否。
// 背包问题的子问题就是是否要加入到背包。有两个子问题，那么就是有两个子for
func PackageOne()(G int){
	results := make([][]bool,0)
	t := make([]bool,maxGoods+1)
	t[1] =  true
	results = append(results,t)
	for i := 1; i < len(goods);i++ { // 背包问题每一个子问题的分割都是装的步骤。
		result := make([]bool,maxGoods+1)
		results = append(results,result)
		// 如果测试到goods已经等于最大限制了，或者是说当我所有的物品都用光了都没有达到最大的限度的时候就可以返回结束了。
		for j := 0;j <= maxGoods;j++ { // 不装
			if results[i-1][j] {
				results[i][j] = true
			}
		}
		for k := 0;k <= maxGoods - goods[i];k++ { // 装
			if results[i-1][k] {
				results[i][k+goods[i]] = true
			}
		}
	}

	// 倒叙输出这个最大值。
	for i := maxGoods;i >= 0;i-- {
		if results[len(goods)-1][i] {
			G = i
			break
		}
	}
	return
}


