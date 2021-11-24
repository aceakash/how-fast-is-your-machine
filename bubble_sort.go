package sort

func BubbleSort(list []int) {
	for i := len(list) - 1; i >= 0; i-- { // 3 times
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
