package main

import (
	"fmt"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/sort"
)

func main() {
  a := []int{2, 3, 1, 6}
  sort.InsertionSort(a)
  fmt.Println(a)

  f := []float32{2.45, 7.67, 1.34, 0.34, 5.67}
  sort.HeapSort(f)
  // sort.QuickSort(f, 0, len(f) - 1)
  fmt.Println(f)

  // s := []string{"afs", "eon", "asf", "sgd", "afsgd"}
  // sort.QuickSort(s, 0, len(s) - 1)
  // sort.InsertionSort(s)
  // fmt.Println(s)

  b := []byte{0x69, 0x21, 0x45, 0x23, 0x67}
  sort.QuickSort(b, 0, len(b) - 1)
  // sort.HeapSort(b)
  fmt.Println(b)
}
