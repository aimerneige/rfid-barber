package pkg

import "strconv"

// String2Uint64 convert string to uint64
func String2Uint64(str string) (uint64, error) {
	if str == "" {
		return 0, nil
	}
	return strconv.ParseUint(str, 10, 64)
}

// String2Int convert string to int
func String2Int(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	return strconv.Atoi(str)
}
