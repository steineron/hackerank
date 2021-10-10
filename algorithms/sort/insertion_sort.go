package sort

type insertion struct {
}

func (ins *insertion) String() string {
	return "Insertion Sort"
}

func (ins *insertion) Sort(list []int) []int {

	for j := 1; j < len(list); j++ {
		for i := j; i > 0; i-- {
			if list[i-1] > list[i] {
				swap(list, i, i-1)
			}
		}
	}
	return list
}
