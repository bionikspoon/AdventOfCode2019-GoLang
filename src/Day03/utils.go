package day03

func minInt(slice []int) int {
	var min int
	for index, element := range slice {
		if index == 0 || element < min {
			min = element
		}
	}
	return min
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
