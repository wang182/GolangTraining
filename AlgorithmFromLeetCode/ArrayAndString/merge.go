package main

import (
	"fmt"
	"sort"
)

// 合并两个有序数组

/**
  给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，
  分别表示 nums1 和 nums2 中的元素数目。

  请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

  注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，
  nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/
func main() {
	merge1([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums1 = append(nums1, nums2[:n]...)
	sort.Ints(nums1)
	fmt.Println(nums1)
}

/**
 * 时间复杂度：O(m+n)

 * 空间复杂度：O(m+n)
 */
func merge1(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	var p1, p2 = 0, 0 // p1, p2分别指向nums1和nums2的起始位置
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
	fmt.Println(nums1)
}
