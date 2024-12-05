package utils

func Contain(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func IndexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return -1
}
