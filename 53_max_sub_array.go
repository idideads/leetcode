package leetcode

/**
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。
示例1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例2：
输入：nums = [1]
输出：1
示例3：
输入：nums = [5,4,-1,7,8]
输出：23
*/
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res, p := nums[0], 0, 0
	for p < len(nums) {
		res += nums[p]    // 累加
		if res > maxSum { // 遇到比上次大的累加，就替换
			maxSum = res
		}
		if res < 0 { // 负数就重置累加
			res = 0
		}
		p++
	}
	return maxSum
}
