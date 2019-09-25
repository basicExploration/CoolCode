package main

var (
	money = []int{12,1,43,23,54,23,64}
)
func PackageTwo()(G int){
	results := make([][]int,0)
	t := make([]int,maxGoods+1)
	t[1] =  12
	results = append(results,t)
	for i := 1; i < len(goods);i++ { // 背包问题每一个子问题的分割都是装的步骤。
		result := make([]int,maxGoods+1)
		results = append(results,result)
		// 如果测试到goods已经等于最大限制了，或者是说当我所有的物品都用光了都没有达到最大的限度的时候就可以返回结束了。
		for j := 0;j <= maxGoods;j++ { // 不装
			if results[i-1][j] > 0 {
				results[i][j] = results[i-1][j]
			}
		}
		for k := 0;k <= maxGoods - goods[i];k++ { // 装
			v := results[i-1][k] + money[i] //
			if v > results[i][k+goods[i]]  {//  这里主要是因为不同的上一层有可能是在这一层是一样的重量但是钱却不一样，那么我们只要把这个重量的最大值搞定即可。
				results[i][k+goods[i]] = v
			}
		}
	}
	
	// 倒叙输出这个最大值。
	m := 0
	for i := 0;i <  maxGoods;i++ {
		if results[len(goods)-1][i] > m {
			m = results[len(goods)-1][i]
		}
	}
	return m
}

func main(){
	PackageTwo()
}