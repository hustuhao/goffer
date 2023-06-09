package binary_search

// leetcode 1482. 制作 m 束花所需的最少天数
// 给你一个整数数组 bloomDay，以及两个整数 m 和 k 。
//
//现需要制作 m 束花。制作花束时，需要使用花园中 相邻的 k 朵花 。
//
//花园中有 n 朵花，第 i 朵花会在 bloomDay[i] 时盛开，恰好 可以用于 一束 花中。
//
//请你返回从花园中摘 m 束花需要等待的最少的天数。如果不能摘到 m 束花则返回 -1 。
func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	var max int
	for _, d := range bloomDay { // 确定天数范围
		if d > max {
			max = d
		}
	}

	min := 1
	left, right := min, max
	for left <= right {
		mid := left + (right-left)/2
		if isValid1(bloomDay, mid, m, k) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func isValid1(bloomDay []int, day, m, k int) bool {
	var totalCnt int
	var cnt int
	for _, d := range bloomDay {
		if d <= day {
			cnt++
		} else {
			cnt = 0
		}
		if cnt == k {
			totalCnt++
			cnt = 0
		}
		if totalCnt == m {
			return true
		}

	}
	return false
}
