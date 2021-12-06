package main

import "fmt"

/**

给定一个整数数组，判断是否存在重复元素。

如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。



示例 1:

输入: [1,2,3,1]
输出: true
示例 2:

输入: [1,2,3,4]
输出: false
示例 3:

输入: [1,1,1,3,3,4,3,2,4,2]
输出: true

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x248f5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func main() {
	arr := []int{1, 2, 1, 1}
	res := eleIsUniq(arr)
	fmt.Println(res)
}

func eleIsUniq(nums []int) bool {

	//set := map[int]struct{}{}
	//for _, v := range nums {
	//	if _, has := set[v]; has {
	//		return true
	//	}
	//	set[v] = struct{}{}
	//}
	//return false

	m := make(map[int]int)
	for k, v := range nums {
		if _, has := m[v]; has {
			return true
		}
		fmt.Println(v)
		m[v] = k
	}
	return false
}
