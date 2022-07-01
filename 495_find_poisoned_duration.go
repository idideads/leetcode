package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
	if duration == 0 {
		return 0
	}
	totalPoisoneSeconds := duration
	for i := 1; i < len(timeSeries); i++ {
		totalPoisoneSeconds += duration
		if timeSeries[i]-timeSeries[i-1] <= duration {
			totalPoisoneSeconds -= (duration - (timeSeries[i] - timeSeries[i-1]))
		}
	}

	return totalPoisoneSeconds
}
