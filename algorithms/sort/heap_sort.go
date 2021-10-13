package sort

/*
Heaps - near complete binary tree (may be missing some leaves)
Heaps as arrays:

parent(i)=A[i/2]
left(i)=A[2*i]
right(i)=A[2*i+1]

max (min) heap property:
for every i>0, A[parent(i)] >= A[i]
*/

type heap struct {
}

func (h *heap) String() string {
	return "Heap Sort"
}

func (h *heap) Sort(list []int) []int {

	mh := maxHeap(list)
	for i := mh.size - 1; i >= 1; i-- {

		swap(mh.heap, 0, i)
		mh.size--
		mh.maxHeapify(0)
	}
	return mh.heap
}

type heapsort struct {
	heap []int
	size int
}

func parent(i int) int {
	return i / 2
}
func (h *heapsort) left(i int) int {
	if i*2 > h.size-1 {
		return -1
	}
	return i * 2
}
func (h *heapsort) right(i int) int {
	if i*2+1 > h.size-1 {
		return -1
	}
	return i*2 + 1
}

func (h *heapsort) maxHeapify(i int) {

	largest := i // the index of the largest of the current, left and right nodes
	l := h.left(i)
	if l >= 0 && h.heap[l] > h.heap[largest] {
		largest = l
	}
	r := h.right(i)
	if r >= 0 && h.heap[r] > h.heap[largest] {
		largest = r
	}
	if largest != i {
		swap(h.heap, largest, i)
		h.maxHeapify(largest)
	}
}

func maxHeap(array []int) *heapsort {
	h := &heapsort{
		heap: array,
		size: len(array),
	}

	for i := h.size / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
	return h
}
