package leetcode

/*
题目： 给孩子发糖。 评分高的孩子比左右两边糖多

思路：
从左往右，下一个孩子评分多，多发一颗。保证左侧满足要求
从右往走，下一个孩子评分多，多发一颗。保证右侧满足要求
*/

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)

	// 左边
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 右边
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] { // 这个条件很重要。当已经比这个孩子高时，就不用再发了。
			candies[i] = candies[i+1] + 1
		}
	}

	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum + n

}
