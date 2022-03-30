package gg

// Unique returns slice with unique values.
// Example: Unique([]int{8, 7, 7, 8}) == []int{8, 7}
func Unique[C comparable](stack []C) []C {
	visited := make(map[C]bool, len(stack))
	resStack := make([]C, 0, len(stack))

	for _, item := range stack {
		if _, ok := visited[item]; !ok {
			visited[item] = true
			resStack = append(resStack, item)
		}
	}

	return resStack
}
