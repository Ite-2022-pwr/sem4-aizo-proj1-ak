package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

const (
  // TestNum mówi ile razy ma zostać wykonane sortowanieTestNum
  TestNum float64 = 100

  // OutputDirectory to folder do zapisywania wynikowych plików csv
  OutputDirectory = "data/output"
)

// SizeAnalysisResult przechowuje średni czas sortowania dla danego rozmiaru tablicy
type SizeAnalysisResult struct {
  Size int
  AvgTime float64
}

//ToStrings zwraca dane struktury w postaci tablicy ciągów znaków
func (sar SizeAnalysisResult) ToStrings() []string {
    return []string{fmt.Sprintf("%v", sar.Size), fmt.Sprintf("%.3f", sar.AvgTime)}
}

// MeasureTimesSizeDependend uruchamia pomiary sortowania ze względu na czas
func MeasureTimesSizeDependend() {
  sizes := []int{10000, 30000, 50000, 70000, 90000, 110000, 130000}
  algorithms := []string{"quick", "heap", "shell", "insertion"}
  pivots := []string{"lewy", "prawy", "środkowy", "losowy"}
  gaps := []string{"shell", "franklazarus"}

  dir := filepath.Join(OutputDirectory, "sizes")
  finalResults := make(map[string][][]string)
  
  for a := 0; a < len(algorithms); a++ {
    log.Println(utils.YellowColor("[*] Sortowanie: ", algorithms[a]))
    for s := 0; s < len(sizes); s++ {
      log.Println(utils.YellowColor("[*] Rozmiar: ", sizes[s]))
      if algorithms[a] == "quick" {
        for p := 0; p < len(pivots); p++ {
          log.Println(utils.YellowColor("[*] Pivot: ", pivots[p]))
          results := MeasureTimeForSize(sizes[s], algorithms[a], pivots[p], "").ToStrings()
          finalResults[algorithms[a] + "_" + pivots[p]] = append(finalResults[algorithms[a]+"_"+pivots[p]], results)
        }
      } else if algorithms[a] == "shell" {
        for g := 0; g < len(gaps); g++ {
          log.Println(utils.YellowColor("[*] Gap: ", gaps[g]))
          results := MeasureTimeForSize(sizes[s], algorithms[a], "", gaps[g]).ToStrings()
          finalResults[algorithms[a] + "_" + gaps[g]] = append(finalResults[algorithms[a]+"_"+gaps[g]], results)
        }
      } else {
        results := MeasureTimeForSize(sizes[s], algorithms[a], "", "").ToStrings()
        finalResults[algorithms[a]] = append(finalResults[algorithms[a]], results)
      }
    }
  }

  fmt.Println(utils.YellowColor("[*] Zapisywanie wyników"))

  for alg, res := range finalResults {
    utils.SaveCSV(filepath.Join(dir, alg + ".csv"), res)
  }
}

// MeasureTimeForSize mierzy średni czas sortowania dla danego algorytmu i rozmiaru tablicy
func MeasureTimeForSize(size int, algorithm, pivot, gap string) SizeAnalysisResult {
  var timesSum float64

  switch algorithm {
  case "insertion":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateRandomArrayInt(size, 1000, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i, TestNum)))
      _, t := sortAnalyzer.InsertionSortAnalysis()
      timesSum += t
    }
  case "heap":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateRandomArrayInt(size, 1000, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i, TestNum)))
      _, t := sortAnalyzer.HeapSortAnalysis()
      timesSum += t
    }
  case "quick":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateRandomArrayInt(size, 1000, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i, TestNum)))
      a, t := sortAnalyzer.QuickSortAnalysis(pivot)
      if a == nil {
        return SizeAnalysisResult{}
      }
      timesSum += t
    }
  case "shell":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateRandomArrayInt(size, 1000, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i, TestNum)))
      a, t := sortAnalyzer.ShellSortAnalysis(gap)
      if a == nil {
        return SizeAnalysisResult{}
      }
      timesSum += t
    }
  default:
    log.Println(utils.RedColor("[!!] Nieznany algorytm sortowania: ", algorithm))
    return SizeAnalysisResult{}
  }

  avg := timesSum / TestNum
  return SizeAnalysisResult{Size: size, AvgTime: avg}
}

