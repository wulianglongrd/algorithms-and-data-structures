package sort

func SelectionSort(a []int) []int {
	length := len(a)
	
	for i := 0; i < length-1; i++ {
		idx := i
		for j := i + 1; j < length; j++ {
			if a[j] < a[idx] {
				idx = j
			}
		}
		a[i], a[idx] = a[idx], a[i]
	}

	return a
}
