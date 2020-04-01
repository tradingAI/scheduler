package server

func uint64ArrToInt64Arr(source []uint64) (target []int64) {
	for _, s := range source {
		target = append(target, int64(s))
	}

	return
}

func int32ArrToInt64Arr(source []int32) (target []int64) {
	for _, s := range source {
		target = append(target, int64(s))
	}

	return
}
