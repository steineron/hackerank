package sort

type quick struct {
}

func (q *quick) String() string {
	return "Quick Sort"
}

func (q *quick) Sort(list []int) []int {

	p := partition(list)
	if p >= 1 {
		q.Sort(list[:p])
	}
	if p+1 < len(list) {
		q.Sort(list[p+1:])
	}
	return list
}

func partition(list []int) int {
	if len(list) <= 1 {
		return 0
	}
	n := len(list)
	pivot := list[n-1]
	partition := 0
	for seeker := 0; seeker < n-1; seeker++ {

		if list[seeker] <= pivot {
			swap(list, seeker, partition)
			partition++
		}
	}
	swap(list, partition, n-1)
	return partition
}
