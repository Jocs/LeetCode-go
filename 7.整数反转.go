/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (32.60%)
 * Likes:    1177
 * Dislikes: 0
 * Total Accepted:    152.7K
 * Total Submissions: 467.2K
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *
 * 示例 1:
 *
 * 输入: 123
 * 输出: 321
 *
 *
 * 示例 2:
 *
 * 输入: -123
 * 输出: -321
 *
 *
 * 示例 3:
 *
 * 输入: 120
 * 输出: 21
 *
 *
 * 注意:
 *
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回
 * 0。
 *
 */

// @lc code=start
import (
	"strconv"
	"strings"
)

// 1534236469

func reverse(x int) int {
	s := strconv.Itoa(x)
	factor := 1
	if strings.HasPrefix(s, "-") {
		factor = -1
		s = strings.TrimPrefix(s, "-")
	}
	tokens := strings.Split(s, "")
	var rs []string
	len := len(tokens)
	for i := len - 1; i >= 0; i-- {
		rs = append(rs, tokens[i])
	}
	result, _ := strconv.Atoi(strings.Join(rs, ""))
	if result >= -(1<<31) && result <= (1<<31)-1 {
		return result * factor
	}
	return 0
}

// @lc code=end

