package leetcode

/*
	二分查找，参考sort.Search
*/
func firstBadVersion(n int) int {
	lowPtr, highPtr := 0, n
	for lowPtr < highPtr {
		mid := (lowPtr + highPtr) / 2
		if isBadVersion(mid) {
			highPtr = mid
		} else {
			lowPtr = mid + 1
		}
	}
	return lowPtr
}

func isBadVersion(version int) bool {
	return version > 5
}
