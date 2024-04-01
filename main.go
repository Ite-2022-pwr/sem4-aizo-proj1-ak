// Package main stanowi punkt wejścia do całego programu
package main

import (
	"fmt"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/generator"
)

func main() {
  // test sort analysis
  n := 100000
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

}
