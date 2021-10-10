package sort

type merge struct {
}

func (m *merge) String() string {
	return "Merge Sort"
}

func (m *merge) Sort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	n := len(list) / 2
	return mergeSorted(m.Sort(list[:n]), m.Sort(list[n:]))

}

func mergeSorted(left, right []int) []int {

	merged := make([]int, len(left)+len(right))

	for i, l, r := 0, 0, 0; l < len(left) || r < len(right); {
		if l < len(left) && r < len(right)  {
			if left[l] <= right[r] {
				merged[i] = left[l]
				l++
			} else {
				merged[i] = right[r]
				r++
			}

		} else if r < len(right) {
			merged[i] = right[r]
			r++

		} else {
			merged[i] = left[l]
			l++
		}
		i++
	}
	return merged

}
