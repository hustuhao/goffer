package array

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//
//
//注意：
//
//对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//示例 1：
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//示例 2：
//
//输入：s = "a", t = "a"
//输出："a"
//解释：整个字符串 s 是最小覆盖子串。
//示例 3:
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//提示：
//
//m == s.length
//n == t.length
//1 <= m, n <= 105
//s 和 t 由英文字母组成
//

// 我写的算法
func minWindow(s string, t string) string {
	window := make(map[uint8]int)
	tm := make(map[uint8]int)
	for i := range t {
		tm[t[i]]++
	}

	var r string
	var minLen int
	left, right := 0, 0
	for right < len(s) {
		// c 是移入窗口的字符
		c := s[right]
		right++
		//进行窗口内数据的一系列更新
		window[c]++

		windowNeedsShrink := true // 判断左侧窗口是否要收缩
		for u, n := range tm {
			if window[u] < n {
				windowNeedsShrink = false
			}
		}
		if minLen == 0 && windowNeedsShrink {
			r = s[left:right]
			minLen = len(r)
		}

		if windowNeedsShrink && len(s[left:right]) < minLen {
			r = s[left:right]
			minLen = len(r)
		}

		for windowNeedsShrink {
			d := s[left]
			left++
			window[d]--

			for u, n := range tm {
				if window[u] < n {
					windowNeedsShrink = false
				}
			}
			if windowNeedsShrink && len(s[left:right]) < minLen {
				r = s[left:right]
				minLen = len(r)
			}
		}
	}
	return r
}

func minWindow1(s string, t string) string {
	window := make(map[uint8]int)
	need := make(map[uint8]int)
	for i := range t {
		need[t[i]]++
	}

	var r string
	var minLen int
	left, right := 0, 0
	var valid int
	for right < len(s) {
		// c 是移入窗口的字符
		c := s[right]
		right++
		//进行窗口内数据的一系列更新
		window[c]++

		if n := need[c]; n > 0 { // 优化:当有需要的字符进入时，才统计计数
			if window[c] == n {
				valid++
			}
		}

		for valid == len(need) {
			tmp := s[left:right]
			if minLen == 0 || minLen > len(tmp) {
				r = tmp
				minLen = len(r)
			}

			d := s[left]
			left++
			window[d]--

			if n := need[d]; n > 0 { // 优化:当需要的字符移出时，才统计计数
				if window[d] < n {
					valid--
				}
			}
		}
	}
	return r
}

// 滑动窗口模板
//func slidingWindow(s string) {
//	window := make(map[rune]int)
//	left, right := 0, 0
//
//	for right < len(s) {
//		// C 是将移入窗口的字符
//		c := s[right]
//		right++
//		fmt.Printf("window: [%d, %d)\n", left, right)
//		// //进行窗口内数据的一系列更新
//
//		// 判断左侧窗口是否要收缩
//		var windowNeedsShrink bool
//		for windowNeedsShrink {
//			// d 是将移出窗口的字符
//			d := s[left]
//			// 缩小窗口
//			left++
//			// 进行窗口内数据的一系列更新
//		}
//	}
