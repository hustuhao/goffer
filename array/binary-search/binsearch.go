package binary_search

// 核心在于：1.什么时候退出for循环。2.缩小边界
// 二分思维的精髓就是：通过已知信息尽可能多地收缩（折半）搜索空间，从而增加穷举效率，快速找到目标。
func binarySearch(list []int, target int) int {
	left, right := 0, len(list)-1
	for left <= right {
		mid := left + (right-left)/2
		if list[mid] == target {
			return mid
		} else if list[mid] > target {
			right = mid - 1
		} else if list[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func leftBound(list []int, target int) int {
	left, right := 0, len(list)-1
	for left <= right {
		mid := left + (right-left)/2
		if list[mid] == target {
			right = mid - 1 // 收缩右边
		} else if list[mid] > target {
			right = mid - 1
		} else if list[mid] < target {
			left = mid + 1
		}
	}

	if left > len(list)-1 {
		return -1
	}

	if list[left] == target {
		return left
	}
	return -1
}

func rightBound(list []int, target int) int {
	left, right := 0, len(list)-1
	for left <= right {
		mid := left + (right-left)/2
		if list[mid] == target { // 123 12
			left = mid + 1 // 收缩左边
		} else if list[mid] > target {
			right = mid - 1
		} else if list[mid] < target {
			left = mid + 1
		}
	}

	if right < 0 {
		return -1
	}

	if list[right] == target {
		return right
	}
	return -1
}
