package b

func isOK(s []int) bool {
	for _, v := range s {
		if v == 0 {
			return true
		}
	}
	return false
}
