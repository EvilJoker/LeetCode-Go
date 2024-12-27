package leetcode

/*
思路：假设第一个足够大。类似插入排序。数组2的元素向数组1插入。

优化：
为了不大量移动元素，就要从2个数组长度之和的最后一个位置开始，依次选取两个数组中大的数，从第一个数组的尾巴开始往头放，只要循环一次以后，就生成了合并以后的数组了。

*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 新数组长度
	p := m + n
	for m > 0 && n > 0 {
		if nums1[m-1] <= nums2[n-1] {
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
		p--
	}

	for m > 0 {
		nums1[p-1] = nums1[m-1]
		p--
		m--
	}
	for n > 0 {
		nums1[p-1] = nums2[n-1]
		p--
		n--
	}

}
