package sort

type bubbler struct {
}

func (b *bubbler) String() string {
	return "Bubble Sort"
}

func (b *bubbler) Sort(list []int) []int {

	for {
		swaps := 0
		for i := 1; i < len(list); i++ {

			if list[i-1] > list[i] {
				// swap
				swap(list, i, i-1)

				swaps++
			}
		}
		if swaps == 0 {
			break
		}
	}

	return list
}
