package libs

func SwitchToMax(current, max int) int {
	if current < max {
		return current + 1
	}
	return max
}

func SwitchToMin(current, min int) int {
	if current > min {
		return current - 1
	}
	return min
}
