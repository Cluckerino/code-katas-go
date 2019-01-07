package comdenom

// MultiLcm finds the LCM of a slice of integers.
func MultiLcm(nums []int) int {
	lastLcm := nums[0]
	for _, num := range nums[1:] {
		lastLcm = lastLcm * num / Gcd(lastLcm, num)
	}
	return lastLcm
}
