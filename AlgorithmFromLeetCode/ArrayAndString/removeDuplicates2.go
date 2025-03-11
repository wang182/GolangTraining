package main

import "fmt"

// 删除有序数组中的重复元素

/**
    给你一个有序数组 nums ，请你 原地 删除重复出现的元素，
    使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

	不要使用额外的数组空间，你必须在 原地 修改输入数组
	并在使用 O(1) 额外空间的条件下完成。
*/

func main() {
	fmt.Println(removeDuplicatesWithDoublePoint([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}

func removeDuplicates2(nums []int) int {
	val, index := 0, 0
	for i, v := range nums {
		if i < 2 {
			val = v
			index++
			continue
		}
		if (val != v) || (val == v && val != nums[index-2]) {
			nums[index] = v
			index++
			val = v
		}
	}
	return index
}

// 双指针
func removeDuplicatesWithDoublePoint(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	// 快慢指针
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
		fmt.Println(nums, "----", slow, fast)
	}
	return slow
}
