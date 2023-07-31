package gfn

// Contains returns true if the array contains the value.
func Contains[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
