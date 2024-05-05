package priorityqueue

// by default it is max priority queue / max heap
// ----------------- Type Declarations --------------------
type PriorityQueue[T comparable] struct {
	arr  []T
	size int
	max  bool
}

// -------------------- Operations--------------------------------------

func NewPriorityQueue[T comparable]() *PriorityQueue[T] {
	return &PriorityQueue[T]{arr: make([]T, 0), size: 0, max: true}
}

func (pq *PriorityQueue[T]) heapify(arr []T, n int, i int ) {
	maxIdx, left, right := i, 2*i+1, 2*i+2 ; 
	if left < n && arr[i] < arr[le]
}

func 
