package leetcode

/*
思路： 双指针法，
先做一个试验,
lmax______l_____r____rmax
lmax 和rmax 代表最高的挡板， 现在计算l 点累计雨水，
当 lmax < rmax. 就意味着桶的短板找到了（无需再比较，我们只关心 lmax）， 累计的雨水 就是 lmax - l
当 lmax > rmax, 此时 l 点无法计算，因为不清楚，右边是否还存在更高的挡板。
--- 但是  r 点就可以计算了，计算之后 rmax 左移，直到找到 比 lmax 更高的点，l 点就又可以计算了
*/
func trap(height []int) int {
	length := len(height)
	sum, left, leftmax, right, rightmax := 0, 0, 0, length-1, 0

	for left < right {
		leftmax = max(leftmax, height[left])
		rightmax = max(rightmax, height[right])

		if leftmax < rightmax {
			sum += leftmax - height[left]
			left++
		} else {
			sum += rightmax - height[right]
			right--
		}

	}
	return sum

}
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
