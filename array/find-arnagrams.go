package array

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
//异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
//
//
//
//示例 1:
//
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
// 示例 2:
//
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

func findAnagrams(s string, p string) []int {
	window := make(map[uint8]int) // 窗口中的字母和数量
	need := make(map[uint8]int)
	for i := range p {
		need[p[i]]++
	}
	res := make([]int, 0)
	left, right, valid := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		if n := need[c]; n > 0 {
			window[c]++
			if window[c] == n {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) { // 满足输出条件
				res = append(res, left)
			}
			c2 := s[left]
			left++ // 左边收缩
			if n := need[c2]; n > 0 {
				if window[c2] == n {
					valid--
				}
				window[c2]--
			}
		}
	}
	return res
}
