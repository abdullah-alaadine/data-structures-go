package internal

func ReverseSlice[K comparable](slc []K) {
	for i := 0; i < len(slc); i++ {
		slc[i], slc[len(slc)-i-1] = slc[len(slc)-i-1], slc[i]
	}
}
