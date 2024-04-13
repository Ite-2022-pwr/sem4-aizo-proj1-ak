package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

const arraySize = 1000000

// TypeAnalysisResult to struktura przechowująca uśredniony wynik pomiarów czasu sortowania
// dla danego typu danych.
type TypeAnalysisResult struct {
  DataType string
  AvgTime float64
}

// ToStrings zwraca wartości pól TypeAnalysisResult w postaci []string.
func (tar TypeAnalysisResult) ToStrings() []string {
  return []string{tar.DataType, fmt.Sprintf("%.3f", tar.AvgTime)}
}

// AnalyzeDifferentDataTypesHeapsort przeprowadza pomiary sortowania przez kopcowanie
// dla różnych typów danych.
func AnalyzeDifferentDataTypesHeapsort() {
  var results [][]string

  results = append(results, AnalyzeHeapsortInt().ToStrings())
  results = append(results, AnalyzeHeapsortInt32().ToStrings())
  results = append(results, AnalyzeHeapsortInt64().ToStrings())
  results = append(results, AnalyzeHeapsortFloat32().ToStrings())
  results = append(results, AnalyzeHeapsortFloat64().ToStrings())
  results = append(results, AnalyzeHeapsortByte().ToStrings())
  results = append(results, AnalyzeHeapsortString().ToStrings())

  utils.SaveCSV(filepath.Join(OutputDirectory, "types", "heapsort.csv"), results)
}

// AnalyzeHeapsortInt przeprowadza pomiary sortowania przez kopcowanie dla typu int.
func AnalyzeHeapsortInt() TypeAnalysisResult {
  var timesSum float64

  for i := 0; i < int(TestNum); i++ {
    log.Println(utils.BlueColor(fmt.Sprintf("[int] Pomiar %v/%v", i + 1, int(TestNum))))
    array := generator.GenerateRandomArrayInt(arraySize, 1000, true)
    sortAnalyzer := analysis.NewSortAnalyzer(array)
    _, t := sortAnalyzer.HeapSortAnalysis()
    timesSum += t
  }

  return TypeAnalysisResult{DataType: "int", AvgTime: timesSum / TestNum}
}

// AnalyzeHeapsortInt32 przeprowadza pomiary sortowania przez kopcowanie dla typu int32.
func AnalyzeHeapsortInt32() TypeAnalysisResult {
  var timesSum float64

  for i := 0; i < int(TestNum); i++ {
    log.Println(utils.BlueColor(fmt.Sprintf("[int32] Pomiar %v/%v", i + 1, int(TestNum))))
    array := generator.GenerateRandomArrayInt32(arraySize, 1000, true)
    sortAnalyzer := analysis.NewSortAnalyzer(array)
    _, t := sortAnalyzer.HeapSortAnalysis()
    timesSum += t
  }

  return TypeAnalysisResult{DataType: "int32", AvgTime: timesSum / TestNum}
}

// AnalyzeHeapsortInt64 przeprowadza pomiary sortowania przez kopcowanie dla typu int64.
func AnalyzeHeapsortInt64() TypeAnalysisResult {
  var timesSum float64

  for i := 0; i < int(TestNum); i++ {
    log.Println(utils.BlueColor(fmt.Sprintf("[int64] Pomiar %v/%v", i + 1, int(TestNum))))
    array := generator.GenerateRandomArrayInt64(arraySize, 1000, true)
    sortAnalyzer := analysis.NewSortAnalyzer(array)
    _, t := sortAnalyzer.HeapSortAnalysis()
    timesSum += t
  }

  return TypeAnalysisResult{DataType: "int64", AvgTime: timesSum / TestNum}
}

// AnalyzeHeapsortFloat32 przeprowadza pomiary sortowania przez kopcowanie dla typu float32.
func AnalyzeHeapsortFloat32() TypeAnalysisResult {
  var timesSum float64

  for i := 0; i < int(TestNum); i++ {
    log.Println(utils.BlueColor(fmt.Sprintf("[float32] Pomiar %v/%v", i + 1, int(TestNum))))
    array := generator.GenerateRandomArrayFloat32(arraySize, 1000, true)
    sortAnalyzer := analysis.NewSortAnalyzer(array)
    _, t := sortAnalyzer.HeapSortAnalysis()
    timesSum += t
  }

  return TypeAnalysisResult{DataType: "float32", AvgTime: timesSum / TestNum}
}

// AnalyzeHeapsortFloat64 przeprowadza pomiary sortowania przez kopcowanie dla typu float64.
func AnalyzeHeapsortFloat64() TypeAnalysisResult {
  var timesSum float64

  for i := 0; i < int(TestNum); i++ {
    log.Println(utils.BlueColor(fmt.Sprintf("[float64] Pomiar %v/%v", i + 1, int(TestNum))))
    array := generator.GenerateRandomArrayFloat32(arraySize, 1000, true)
    sortAnalyzer := analysis.NewSortAnalyzer(array)
    _, t := sortAnalyzer.HeapSortAnalysis()
    timesSum += t
  }

  return TypeAnalysisResult{DataType: "float64", AvgTime: timesSum / TestNum}
}

// AnalyzeHeapsortByte przeprowadza pomiary sortowania przez kopcowanie dla typu byte.
func AnalyzeHeapsortByte() TypeAnalysisResult {
  var timesSum float64

  for i := 0; i < int(TestNum); i++ {
    log.Println(utils.BlueColor(fmt.Sprintf("[byte] Pomiar %v/%v", i + 1, int(TestNum))))
    array := generator.GenerateRandomArrayByte(arraySize, true)
    sortAnalyzer := analysis.NewSortAnalyzer(array)
    _, t := sortAnalyzer.HeapSortAnalysis()
    timesSum += t
  }

  return TypeAnalysisResult{DataType: "byte", AvgTime: timesSum / TestNum}
}

// AnalyzeHeapsortString przeprowadza pomiary sortowania przez kopcowanie dla typu string.
func AnalyzeHeapsortString() TypeAnalysisResult {
  var timesSum float64

  for i := 0; i < int(TestNum); i++ {
    log.Println(utils.BlueColor(fmt.Sprintf("[string] Pomiar %v/%v", i + 1, int(TestNum))))
    array := generator.GenerateRandomArrayString(arraySize, 20, true)
    sortAnalyzer := analysis.NewSortAnalyzer(array)
    _, t := sortAnalyzer.HeapSortAnalysis()
    timesSum += t
  }

  return TypeAnalysisResult{DataType: "string", AvgTime: timesSum / TestNum}
}
