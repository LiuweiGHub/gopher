package main

import (
	"fmt"
	"math"
	"sort"
)

/**
给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。

请你返回让 s 成为回文串的 最少操作次数 。

「回文串」是正读和反读都相同的字符串。



示例 1：

输入：s = "zzazz"
输出：0
解释：字符串 "zzazz" 已经是回文串了，所以不需要做任何插入操作。
示例 2：

输入：s = "mbadm"
输出：2
解释：字符串可变为 "mbdadbm" 或者 "mdbabdm" 。

示例 3：
输入：s = "leetcode"
输出：5
解释：插入 5 个字符后字符串变为 "leetcodocteel" 。
示例 4：

输入：s = "g"
输出：0
示例 5：

输入：s = "no"
输出：1


提示：

1 <= s.length <= 500
s 中所有字符都是小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	//[[5,4],[6,4],[6,7],[2,3]]
	s := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	res := maxEnvelopes(s)
	fmt.Println(res)

	fmt.Println(int(^int32(^uint32(0) >> 1)))
}

func bbc(nums []int) int {

	dp := make([]int, len(nums))

	for i, _ := range nums {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j])+1))
			}
		}
	}

	res := 0

	for i := 0; i < len(dp); i++ {
		res = int(math.Max(float64(res), float64(dp[i])))
	}

	return res

}

func maxEnvelopes(envelopes [][]int) int {

	//送分
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	//按宽度排序，从小到大
	//如果宽度一直，则按高度从大到小排
	//slice 切片排序 。 语法不熟练
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	fmt.Println(envelopes)

	//然后，求高度中的最长递增子序列(LIS问题)
	f := make([]int, n)
	for i, _ := range f {
		f[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	return max(f...)
}

func max(a ...int) int {
	fmt.Println(a)
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
