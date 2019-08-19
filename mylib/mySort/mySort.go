package mySort

func BubbleSort(array []int) []int {
	for i:= 0; i < len(array); i++ {
		for j:= i; j < len(array); j++ {
			if array[i] > array[j] {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
	return array
}

func SelectionSort(array []int) []int {

	for i:= 0; i < len(array); i++ {
		var min int = array[i]
		var index int = i
		for j:= i; j < len(array); j++ {
			if min > array[j] {
				min = array[j]
				index = j
			}
		}
		temp := array[index]
		array[index] = array[i]
		array[i] = temp
	}
	return array
}