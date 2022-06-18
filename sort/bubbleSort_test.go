package sort

import (
	"fmt"
)

func ExampleBubbleSort() {
	a := []int{4, 5, 6, 3, 2, 1}
	b := []int{1, 2, 3, 4, 6, 5}
	fmt.Println(BubbleSort(a))
	fmt.Println(BubbleSort(b))
	// output:
	// [1 2 3 4 5 6]
	// [1 2 3 4 5 6]
}
