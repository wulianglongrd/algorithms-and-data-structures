package sort

import "fmt"

func ExampleInsertionSort() {
	a := []int{4, 5, 6, 1, 3, 2}
	b := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(InsertionSort(a))
	fmt.Println(InsertionSort(b))
	// output:
	// [1 2 3 4 5 6]
	// [1 2 3 4 5 6]
}
