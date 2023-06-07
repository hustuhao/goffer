package window

// 回顾一下，遇到子数组/子串相关的问题，你只要能回答出来以下几个问题，就能运用滑动窗口算法：
// 1、什么时候应该扩大窗口？2、什么时候应该缩小窗口？3、什么时候应该更新答案？

// 题目:无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	window := make(map[uint8]int)
	left, right, valid, max := 0, 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		if n := window[c]; n == 1 {
			valid++
		} else {
			valid--
		}
		if valid > max {
			max = valid
		}
		for right-left > valid {
			cc := s[left]
			left++
			window[cc]--
			if n := window[cc]; n == 1 {
				valid++
			} else {
				valid--
			}
			if valid > max {
				max = valid
			}
		}
	}
	return max
}

// Optimised by ChatGPT
func lengthOfLongestSubstring2(s string) int {
	window := make(map[uint8]int)
	left, right, max := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		if window[c] > 1 {
			for s[left] != c {
				window[s[left]]--
				left++
			}
			window[s[left]]--
			left++
		}
		if right-left > max {
			max = right - left
		}
	}
	return max
}
