package array

// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。

// 示例 1：
// 输入：s1 = "ab" s2 = "eidbaooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("ba").
// 示例 2：
//
// 输入：s1= "ab" s2 = "eidboaoo"
// 输出：false

// 提示：
// 1 <= s1.length, s2.length <= 104
// s1 和 s2 仅包含小写字母
func checkInclusion(s1 string, s2 string) bool {
	window := make(map[uint8]int) // 想清楚 window 里面装的是什么:装的是子串组成
	need := make(map[uint8]int)
	for i := range s1 {
		need[s1[i]]++
	}

	left, right := 0, 0
	var valid int

	for right < len(s2) {
		c := s2[right]
		right++
		if n := need[c]; n > 0 {
			window[c]++
			if n == window[c] {
				valid++
			}
		}

		for right-left >= len(s1) { // 窗口左边收缩条件
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if n := need[d]; n > 0 {
				if n == window[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
