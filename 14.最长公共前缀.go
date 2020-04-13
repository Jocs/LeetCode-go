/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (33.89%)
 * Likes:    625
 * Dislikes: 0
 * Total Accepted:    99.8K
 * Total Submissions: 292.2K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */

// @lc code=start

func longestCommonPrefix(strs []string) string {
	var result string
	strsLen := len(strs)
	var minLen int
	for i := 0; i < strsLen; i++ {
		strLen := len(strs[i])
		if i == 0 {
			minLen = strLen
		} else {
			if strLen < minLen {
				minLen = strLen
			}
		}
	}

	for i := 0; i < minLen; i++ {
		var count int
		for j := 1; j < strsLen; j++ {
			if strs[j][i] == strs[j-1][i] {
				count++
			} else {
				break
			}
		}
		if count == strsLen-1 {
			result += string(strs[0][i])
		} else {
			break
		}
	}
	return result
}

// @lc code=end

