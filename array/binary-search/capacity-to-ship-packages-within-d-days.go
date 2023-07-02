package binary_search

// leetcode 1011

// 传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。
//
//传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量（weights）的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
//
//返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。

// 提示：
//
//1 <= days <= weights.length <= 5 * 104
//1 <= weights[i] <= 500

// 要点： 二分法+边界条件

// 优化二分查找的版本
func shipWithinDays(weights []int, days int) int {
	var maxCapacity int
	minCapacity := 500
	for _, w := range weights {
		maxCapacity += w
		if minCapacity > w {
			minCapacity = w
		}
	}

	if days == 1 { // 注意边界，否则会超时
		return maxCapacity
	}

	left, right := minCapacity, maxCapacity
	for left <= right {
		mid := left + (right-left)/2
		if isValid(weights, days, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func isValid(weights []int, days int, cap int) bool {
	var sum int
	var count int
	for k := 0; k < len(weights); k++ {
		if sum+weights[k] > cap {
			count++ // 装满了，剩余的放到第二天装
			sum = 0 // 重量消耗掉
			k--
		} else if sum+weights[k] == cap {
			count++
			sum = 0
		} else {
			sum += weights[k] // 没有装满，可以继续尝试装
			if k == len(weights)-1 {
				count++
			}
		}
		if count > days {
			break
		}
	}
	if count <= days {
		return true
	}
	return false
}

// 未优化二分查找的版本
func shipWithinDays1(weights []int, days int) int {
	var maxCapacity int
	minCapacity := 500
	for _, w := range weights {
		maxCapacity += w
		if minCapacity > w {
			minCapacity = w
		}
	}

	if days == 1 { // 注意边界，否则会超时
		return maxCapacity
	}

	for i := minCapacity; i <= maxCapacity; i++ {
		var sum int
		var count int
		for k := 0; k < len(weights); k++ {
			if sum+weights[k] > i {
				count++ // 装满了，剩余的放到第二天装
				sum = 0 // 重量消耗掉
				k--
			} else if sum+weights[k] == i {
				count++
				sum = 0
			} else {
				sum += weights[k] // 没有装满，可以继续尝试装
				if k == len(weights)-1 {
					count++
				}
			}
			if count > days {
				break
			}
		}
		if count <= days {
			return i
		}
	}
	return 0
}
