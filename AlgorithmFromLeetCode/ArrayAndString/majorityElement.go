package main

import "fmt"

// 多数元素

/**
	给定一个大小为 n 的数组 nums ，返回其中的多数元素。
    多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

    你可以假设数组是非空的，并且给定的数组总是存在多数元素。

	示例 1：

		输入：nums = [3,2,3]
		输出：3

	示例 2：

		输入：nums = [2,2,1,1,1,2,2]
		输出：2
*/

func main() {
	fmt.Println(majorityElement([]int{0, 2, 2}))
}

func majorityElement(nums []int) int {
	val, ant := 0, 0
	for _, v := range nums {
		if ant == v {
			val++
		} else if val == 0 {
			ant = v
		} else {
			val--
		}
	}
	return ant
}
