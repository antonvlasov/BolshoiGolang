package benchpress

func SliceFunc(s []int) {
	var sum int

	for i, e := range s[:len(s)-1] {
		sum += e * s[i+1]
	}
}
