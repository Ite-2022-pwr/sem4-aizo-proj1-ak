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
  // array := generator.GenerateRandomArrayInt(n, 1000)
  // array := generator.GenerateRandomArrayByte(n)
  // array := generator.GenerateRandomArrayFloat64(n)
  array := generator.GenerateRandomArrayString(n, 10)
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

  fmt.Println("asc int:", generator.GenerateAscendingSortedInt(10))
  fmt.Println("desc int:", generator.GenerateDescendingSortedInt(10))
  fmt.Println("asc int32:", generator.GenerateAscendingSortedInt32(10))
  fmt.Println("desc int32:", generator.GenerateDescendingSortedInt32(10))
  fmt.Println("asc int64:", generator.GenerateAscendingSortedInt64(10))
  fmt.Println("desc int64:", generator.GenerateDescendingSortedInt64(10))
  fmt.Println("asc float32:", generator.GenerateAscendingSortedFloat32(5))
  fmt.Println("desc float32:", generator.GenerateDescendingSortedFloat32(5))
  fmt.Println("asc float64:", generator.GenerateAscendingSortedFloat64(5))
  fmt.Println("desc float64:", generator.GenerateDescendingSortedFloat64(5))

}
