package leetcode

/*
# 定义

递归分治法是一种解决问题的策略，将一个大问题分解为多个较小的子问题，然后递归地解决这些子问题，最后将结果合并起来。在大整数乘法中，递归分治法可以通过将大整数拆分为更小的部分，分别进行计算并合并结果，从而提高计算效率。

## 示例

大整数乘法的经典分治算法是 **Karatsuba 算法**。它通过将两个大整数分解为两部分，利用递归和分治思想，减少了传统大整数乘法的计算复杂度。

以下是一个基于分治法的 Karatsuba 算法的 Golang 示例：
*/

import (
	"fmt"
	"math"
)

// Karatsuba 算法计算两个大整数的乘积
func karatsuba(x, y int) int {
	// 基本情况：当数字足够小，直接返回乘积
	if x < 10 || y < 10 {
		return x * y
	}

	// 计算数字的长度
	// fmt.Sprint(x) 将整数 x 转换为字符串表示形式。例如，1234 会被转换为字符串 "1234"。
	// m 是分割点，m 是 x 和 y 的最大长度的一半。例如，12345678 的最大长度是 10，因此 m 等于 5。
	m := int(math.Max(float64(len(fmt.Sprint(x))), float64(len(fmt.Sprint(y))))) / 2
	// 计算分割的基数，比如 pow = 10^m
	pow := int(math.Pow(10, float64(m)))

	// 分割数字为两部分 x = x1 * pow + x0; y = y1 * pow + y0
	x1 := x / pow
	x0 := x % pow

	y1 := y / pow
	y0 := y % pow

	// 递归地计算三个子问题
	// x * y = (x1 * pow + x0) * (y1 * pow + y0) = (x1 * y1) * pow * pow + (x1*y0 + x0*y1) * pow + (x0*y0)
	// 其中 (x1*y0 + x0*y1) = (x1+x0) * (y1+y0) - (x0*y0) - (x1*y1) --- 有效减少多项式的乘法计算量
	z2 := karatsuba(x1, y1)
	z0 := karatsuba(x0, y0)
	z1 := karatsuba(x1+x0, y1+y0) - z2 - z0

	// 合并结果
	return z2*int(math.Pow(10, float64(2*m))) + z1*pow + z0
}

func main() {
	x := 1234
	y := 5678
	result := karatsuba(x, y)
	fmt.Printf("Result of %d * %d is: %d\n", x, y, result)
}
