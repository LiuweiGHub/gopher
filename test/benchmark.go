package test

// MakeSliceWithoutAlloc 不预分配
func MakeSliceWithoutAlloc() []int {

	var newSlice []int

	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, i)
	}

	return newSlice
}

// MakeSliceWithAlloc 预分配
func MakeSliceWithAlloc() []int {
	newSlice := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, i)
	}
	return newSlice
}
