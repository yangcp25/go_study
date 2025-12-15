package main

import "math"

func main() {

}

func findMedianSortedArrays(A []int, B []int) float64 {
	// 确保 A 是较短的数组
	if len(A) > len(B) {
		return findMedianSortedArrays(B, A)
	}

	m, n := len(A), len(B)
	total := m + n
	half := total / 2

	left, right := 0, m

	for left <= right {
		i := (left + right) / 2 // A 的切点
		j := half - i           // B 的切点（由 i 推出）

		// 四个关键变量 —— 划分点两侧的值
		Aleft := math.Inf(-1)
		Aright := math.Inf(1)
		Bleft := math.Inf(-1)
		Bright := math.Inf(1)

		if i > 0 {
			Aleft = float64(A[i-1])
		}
		if i < m {
			Aright = float64(A[i])
		}
		if j > 0 {
			Bleft = float64(B[j-1])
		}
		if j < n {
			Bright = float64(B[j])
		}

		// ① 划分正确：左边的最大 <= 右边的最小
		if Aleft <= Bright && Bleft <= Aright {

			if total%2 == 1 { // 奇数，取右边最小
				return math.Min(Aright, Bright)
			}

			// 偶数，取左右最大与右最小平均数
			leftMax := math.Max(Aleft, Bleft)
			rightMin := math.Min(Aright, Bright)
			return (leftMax + rightMin) / 2
		}

		// ② 划分不正确：需要调整 i
		if Aleft > Bright {
			// 左半边太大，i 往左
			right = i - 1
		} else {
			// 左半边太小，i 往右
			left = i + 1
		}
	}

	return 0 // 理论不会走到这里
}
