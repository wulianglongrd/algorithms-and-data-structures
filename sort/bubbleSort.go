package sort

// BubbleSort 实现关键点
// j 从最左边开始，比较j和j-1的数据
// 当某次冒泡操作已经没有数据交换时，说明已经达到完全有序。
func BubbleSort(a []int) []int {
	length := len(a)
	flag := false
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				flag = true
			}
		}

		// 如果没有数据交换 -> 任意两个相邻的元素都是有序的
		// 如果任意两个相邻的元素都是有序的 -> 数据整体是有序的
		// => 如果没有数据交换 -> 数据整体是有序的
		if !flag {
			break
		}
	}

	return a
}
