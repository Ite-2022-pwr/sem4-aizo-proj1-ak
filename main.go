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
  fmt.Println(f)
}
