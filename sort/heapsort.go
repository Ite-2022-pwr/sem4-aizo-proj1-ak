package sort

import "golang.org/x/exp/constraints"

func HeapSort[T constraints.Ordered](array []T) {
  buildHeap(array)

  for i := len(array) - 1; i > 0; i-- {
    array[0], array[i] = array[i], array[0]
    heapify(array, i, 0)
  }
}

func buildHeap[T constraints.Ordered](array []T) {
  n := len(array)
  for i := n / 2; i >= 0; i-- {
    heapify(array, n, i)
  }
}

func heapify[T constraints.Ordered](array []T, size, root int) {
  left, right := 2 * root + 1, 2 * root + 2
  largest := root

  if left < size && array[left] > array[largest] {
    largest = left
  }

  if right < size && array[right] > array[largest] {
    largest = right
  }

  if largest != root {
    array[root], array[largest] = array[largest], array[root]
    heapify(array, size, largest)
  }
}
