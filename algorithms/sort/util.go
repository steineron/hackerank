package sort

func swap(list []int, i, j int) {
	if i < len(list) && j < len(list) && i >= 0 && j >= 0 {
		tmp := list[i]
		list[i] = list[j]
		list[j] = tmp
	}
}
