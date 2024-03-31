// Package main stanowi punkt wejścia do całego programu
package main

import (
	"fmt"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/generator"
)

func main() {
  // test sort analysis
  arrayInt := generator.GenerateRandomArrayInt(10, 10)
  fmt.Println("Array:", arrayInt)

  sortMeter := analysis.NewSortMeter(arrayInt)
  // sortedArr := sortMeter.QuickSortAnalysis("lewy")
  // sortedArr := sortMeter.QuickSortAnalysis("prawy")
  // sortedArr := sortMeter.QuickSortAnalysis("środkowy")
  sortedArr := sortMeter.QuickSortAnalysis("losowy")
  fmt.Println("QuickSort:", sortedArr)

  sortedArr = sortMeter.InsertionSortAnalysis()
  fmt.Println("InsertionSort:", sortedArr)

  sortedArr = sortMeter.HeapSortAnalysis()
  fmt.Println("HeapSort:", sortedArr)

  sortedArr = sortMeter.ShellSortAnalysis("shell")
  fmt.Println("ShellSort:", sortedArr)

}
