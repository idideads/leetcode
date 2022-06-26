package leetcode

/*
给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。

解题思路：如按题目描述，正向实现提交只会超过时间限制。需要逆向思考，每操作一次等于缩小了第n个元素和n-1个元素之间的相对差异。
让所有元素都变成相等的最少步数，等价于让所有元素差异减到最小的那个数，也就是【求和 - 最小元素*元素个数】
*/
func minMoves(nums []int) int {
	sum, min, length := nums[0], nums[0], len(nums)
	for i := 1; i < length; i++ {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}
	return sum - min*length
}
