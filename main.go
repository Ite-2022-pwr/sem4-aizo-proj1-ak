// Package main is an entry point to the program
package main

import (
	"fmt"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/sort"
)

func main() {
  a := []int{2, 3, 1, 6}

  // sort.InsertionSort(a)
  sort.QuickSort(a, 0, len(a) - 1, sort.GetPivotHigh)
  fmt.Println(a)

  f := []float32{2.45, 7.67, 1.34, 0.34, 5.67}
  // sort.HeapSort(f)
  sort.QuickSort(f, 0, len(f) - 1, sort.GetPivotRandom)
  fmt.Println(f)

  s := []string{"afs", "eon", "asf", "sgd", "afsgd"}
  sort.QuickSort(s, 0, len(s) - 1, sort.GetPivotLow)
  // sort.InsertionSort(s)
  fmt.Println(s)

  b := []byte{0x69, 0x21, 0x45, 0x23, 0x67}
  sort.QuickSort(b, 0, len(b) - 1, sort.GetPivotMiddle)
  // sort.HeapSort(b)
  fmt.Println(b)

  a1 := []int32{-12, 4, -3, 14, 7, 0, 2, 9}
  sort.ShellSort(a1, sort.CalculateShellGap)
  fmt.Println(a1)

  a2 := []float64{-3.14, 0.65, 5.76, -6.38, 7.12, 4.01}
  sort.ShellSort(a2, sort.CalculateFrankLazarusGap)
  fmt.Println(a2)

  // test sortmeter
  arr := []float64{-3.14, 0.65, 5.76, -6.38, 7.12, 4.01}
  sortMeter := analysis.NewSortMeter(arr)
  sortedArr := sortMeter.QuickSort("lewy")
  fmt.Println("QuickSort:", sortedArr)
  fmt.Println(sortMeter.Data)
  fmt.Println("Tablica posortowana:", analysis.IsArraySorted(sortedArr))

  sortedArr = sortMeter.InsertionSort()
  fmt.Println("InsertionSort:", sortedArr)

  sortedArr = sortMeter.ShellSort("hibbard")
  fmt.Println("ShellSort:", sortedArr)

}
