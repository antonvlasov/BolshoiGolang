package benchpress

func FindSum(list []int) int {
	sum := 0

	for _, number := range list {
		sum += number
	}

	return sum
}

func FindSumSlow(list []*int) int {
	sum := 0

	for _, number := range list {
		sum += *number
	}

	return sum
}
