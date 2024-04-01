// Package main stanowi punkt wejścia do całego programu
package main

import (
	"fmt"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/generator"
)

func main() {
  // test sort analysis
  n := 1000
  array := generator.GenerateRandomArrayInt(n, 1000, true)
  // array := generator.GenerateRandomArrayByte(n, true)
  // array := generator.GenerateRandomArrayFloat64(n, true)
  // array := generator.GenerateRandomArrayString(n, 10, true)
  // fmt.Println("Array:", arrayInt[:20])

  sortMeter := analysis.NewSortMeter(array)
  // sortedArr := sortMeter.QuickSortAnalysis("lewy")
  sortedArr := sortMeter.QuickSortAnalysis("prawy")
  // sortedArr := sortMeter.QuickSortAnalysis("środkowy")
  // sortedArr := sortMeter.QuickSortAnalysis("losowy")
  fmt.Println("QuickSort:", sortedArr[0])

  sortedArr = sortMeter.InsertionSortAnalysis()
  // fmt.Println("InsertionSort:", sortedArr)

  sortedArr = sortMeter.HeapSortAnalysis()
  // fmt.Println("HeapSort:", sortedArr)

  sortedArr = sortMeter.ShellSortAnalysis("shell")
  // fmt.Println("ShellSort:", sortedArr)

  fmt.Println("asc int:", generator.GenerateAscendingSortedArrayInt(10, true))
  fmt.Println("desc int:", generator.GenerateDescendingSortedArrayInt(10, true))
  fmt.Println("asc int32:", generator.GenerateAscendingSortedArrayInt32(10, true))
  fmt.Println("desc int32:", generator.GenerateDescendingSortedArrayInt32(10, true))
  fmt.Println("asc int64:", generator.GenerateAscendingSortedArrayInt64(10, true))
  fmt.Println("desc int64:", generator.GenerateDescendingSortedArrayInt64(10, true))
  fmt.Println("asc float32:", generator.GenerateAscendingSortedArrayFloat32(5, true))
  fmt.Println("desc float32:", generator.GenerateDescendingSortedArrayFloat32(5, true))
  fmt.Println("asc float64:", generator.GenerateAscendingSortedArrayFloat64(5, true))
  fmt.Println("desc float64:", generator.GenerateDescendingSortedArrayFloat64(5, true))
  fmt.Println("asc byte:", generator.GenerateAscendingSortedArrayByte(10, true))
  fmt.Println("desc byte:", generator.GenerateDescendingSortedArrayByte(10, true))
  fmt.Println("asc string:", generator.GenerateAscendingSortedArrayString(5, true))
  fmt.Println("desc string:", generator.GenerateDescendingSortedArrayString(5, true))

  // ba := generator.GenerateAscendingSortedArrayByte(10000)
  // fmt.Println(analysis.IsArraySortedAscending(ba))
  //
  // ba = generator.GenerateDescendingSortedArrayByte(10000)
  // fmt.Println(analysis.IsArraySortedDescending(ba))

  fmt.Println("part int:", generator.GeneratePartiallySortedArrayInt(10, 33, true))
  fmt.Println("part int32:", generator.GeneratePartiallySortedArrayInt32(10, 33, true))
  fmt.Println("part int64:", generator.GeneratePartiallySortedArrayInt64(10, 33, true))
  fmt.Println("part float32:", generator.GeneratePartiallySortedArrayFloat32(10, 33, true))
  fmt.Println("part float64:", generator.GeneratePartiallySortedArrayFloat64(10, 33, true))
  fmt.Println("part byte:", generator.GeneratePartiallySortedArrayByte(10, 33, true))
  fmt.Println("part string:", generator.GeneratePartiallySortedArrayString(10, 33, true))
}
