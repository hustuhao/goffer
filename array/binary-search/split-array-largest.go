package binary_search

// https://leetcode.cn/problems/split-array-largest-sum/

// 给定一个非负整数数组 nums 和一个整数 m ，你需要将这个数组分成 m 个非空的连续子数组。
//
//设计一个算法使得这 m 个子数组各自和的最大值最小。
//
//
//
//示例 1：
//
//输入：nums = [7,2,5,10,8], m = 2
//输出：18
//解释：
//一共有四种方法将 nums 分割为 2 个子数组。
//其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
//因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
func splitArray(nums []int, m int) int {
	maxSum := 0
	sum := 0
	for _, num := range nums {
		maxSum = max(maxSum, num)
		sum += num
	}

	left := maxSum
	right := sum

	for left < right {
		mid := left + (right-left)/2
		if isValid2(nums, m, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func isValid2(nums []int, m, target int) bool {
	count := 1 // 为什么从一开始：因为 count 表示分割的子数组的个数，最开始起码有一个数组，后面被切割一次，子数组的数量+1
	sum := 0
	for _, num := range nums {
		sum += num
		if sum > target {
			count++
			sum = num
			if count > m { // 是否满足要求:要求数组的和不能大于 target
				return false
			}
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
