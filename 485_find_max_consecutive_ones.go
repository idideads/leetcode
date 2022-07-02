package leetcode

func findMaxConsecutiveOnes(nums []int) int {
    sum, max := 0, 0
	for i, n := range nums {
		sum += n
		if n == 0 || i == len(nums)-1 {
			if sum > max {
				max = sum
			}
			sum = 0
		}
	}
	return max
}