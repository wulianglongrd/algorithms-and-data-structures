package sort

// InsertionSort 关键实现点：
// 分为已排序和未排序两个区间，初始已排序区间只有一个元素，就是第一个元素。
// 每次排序在已排序区间查找位置插入当前要排序的元素
func InsertionSort(a []int) []int {
	length := len(a)

	// 从小到大排序
	// 初始第一个元素已经有序
	for i := 1; i < length; i++ {
		j := i - 1
		value := a[i] // 当前要排序的元素
		for ; j >= 0; j-- {
			// 已排序区间：如果第一个元素小于当前元素，则左侧的元素都小于当前排序元素，直接break
			if a[j] > value {
				a[j+1] = a[j] // 当前元素向后移一位
			} else {
				break
			}
		}

		// 因为 j-- 跳出条件时已经减1，这里加回1
		a[j+1] = value
	}

	return a
}
