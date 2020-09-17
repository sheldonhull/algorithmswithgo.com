package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {

	var match bool = false
	for _, i := range list {
		if i == num {
			match = true
			break
		}
	}
	return match
}
