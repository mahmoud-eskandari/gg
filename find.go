package gg

// LastIndexOf returns last index of needle in heyStack.
// Example: gg.IndexOf([]int{1, 2, 3, 4, 5, 6, 5, 0}, 5) == 4
func IndexOf[C comparable](heyStack []C, needle C) int {
	for i, item := range heyStack {
		if item == needle {
			return i
		}
	}

	return -1
}

// LastIndexOf returns last index of needle in heyStack.
// Example: gg.LastIndexOf([]int{1, 2, 3, 4, 5, 6, 5, 0}, 5) == 6
func LastIndexOf[C comparable](heyStack []C, needle C) int {
	res := -1
	for i, item := range heyStack {
		if item == needle {
			res = i
		}
	}

	return res
}

// AllIndexesOf returns all index of needle in heyStack.
// Example: gg.AllIndexesOf([]int{1, 2, 3, 4, 5, 6, 5, 0}, 5) == []int{4, 6}
func AllIndexesOf[C comparable](heyStack []C, needle C) []int {
	res := make([]int, 0, len(heyStack))

	for i, item := range heyStack {
		if item == needle {
			res = append(res, i)
		}
	}

	return res
}

// CountByValue returns count of needle in heyStack.
// Example: gg.CountByValue([]int{5, 6, 5}, 5) == 2
func CountByValue[C comparable](heyStack []C, needle C) int {
	res := 0

	for _, item := range heyStack {
		if item == needle {
			res++
		}
	}

	return res
}
