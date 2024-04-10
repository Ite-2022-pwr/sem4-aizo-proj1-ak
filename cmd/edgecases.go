package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

type EdgeCaseResult struct {
  Algorithm string
  AvgTime  float64
}

func (ecr EdgeCaseResult) ToStrings() []string {
  return []string{ecr.Algorithm, fmt.Sprintf("%.3f", ecr.AvgTime)}
}

func TestEdgeCases() {
  var results [][]string

  // W pełni losowa
  results = append(results, TestFullRandom("heap", "", "").ToStrings())
  results = append(results, TestFullRandom("quick", "lewy", "").ToStrings())
  results = append(results, TestFullRandom("quick", "prawy", "").ToStrings())
  results = append(results, TestFullRandom("quick", "środkowy", "").ToStrings())
  results = append(results, TestFullRandom("quick", "losowy", "").ToStrings())
  results = append(results, TestFullRandom("shell", "", "shell").ToStrings())
  results = append(results, TestFullRandom("shell", "", "franklazarus").ToStrings())
  results = append(results, TestFullRandom("insertion", "", "").ToStrings())

  utils.SaveCSV(filepath.Join(OutputDirectory, "edge_cases", "random.csv"), results)

  results = nil

  // Posortowana rosnąco
  results = append(results, TestSortedAscending("heap", "", "").ToStrings())
  results = append(results, TestSortedAscending("quick", "lewy", "").ToStrings())
  results = append(results, TestSortedAscending("quick", "prawy", "").ToStrings())
  results = append(results, TestSortedAscending("quick", "środkowy", "").ToStrings())
  results = append(results, TestSortedAscending("quick", "losowy", "").ToStrings())
  results = append(results, TestSortedAscending("shell", "", "shell").ToStrings())
  results = append(results, TestSortedAscending("shell", "", "franklazarus").ToStrings())
  results = append(results, TestSortedAscending("insertion", "", "").ToStrings())

  utils.SaveCSV(filepath.Join(OutputDirectory, "edge_cases", "ascending.csv"), results)

  results = nil

  // Posortowana malejąco
  results = append(results, TestSortedDescending("heap", "", "").ToStrings())
  results = append(results, TestSortedDescending("quick", "lewy", "").ToStrings())
  results = append(results, TestSortedDescending("quick", "prawy", "").ToStrings())
  results = append(results, TestSortedDescending("quick", "środkowy", "").ToStrings())
  results = append(results, TestSortedDescending("quick", "losowy", "").ToStrings())
  results = append(results, TestSortedDescending("shell", "", "shell").ToStrings())
  results = append(results, TestSortedDescending("shell", "", "franklazarus").ToStrings())
  results = append(results, TestSortedDescending("insertion", "", "").ToStrings())

  utils.SaveCSV(filepath.Join(OutputDirectory, "edge_cases", "descending.csv"), results)

  results = nil

  // Częściowo posortowana (w 33%)
  results = append(results, TestPartiallySorted("heap", "", "").ToStrings())
  results = append(results, TestPartiallySorted("quick", "lewy", "").ToStrings())
  results = append(results, TestPartiallySorted("quick", "prawy", "").ToStrings())
  results = append(results, TestPartiallySorted("quick", "środkowy", "").ToStrings())
  results = append(results, TestPartiallySorted("quick", "losowy", "").ToStrings())
  results = append(results, TestPartiallySorted("shell", "", "shell").ToStrings())
  results = append(results, TestPartiallySorted("shell", "", "franklazarus").ToStrings())
  results = append(results, TestPartiallySorted("insertion", "", "").ToStrings())

  utils.SaveCSV(filepath.Join(OutputDirectory, "edge_cases", "partially.csv"), results)

  results = nil
}

func TestFullRandom(algorithm, pivot, gap string) EdgeCaseResult {
  var timesSum float64
  size := 100000

  switch algorithm {
  case "insertion":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateRandomArrayInt(size, 1000, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      _, t := sortAnalyzer.InsertionSortAnalysis()
      timesSum += t
    }
  case "heap":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateRandomArrayInt(size, 1000, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      _, t := sortAnalyzer.HeapSortAnalysis()
      timesSum += t
    }
  case "quick":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateRandomArrayInt(size, 1000, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      a, t := sortAnalyzer.QuickSortAnalysis(pivot)
      if a == nil {
        return EdgeCaseResult{}
      }
      timesSum += t
    }
  case "shell":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateRandomArrayInt(size, 1000, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      a, t := sortAnalyzer.ShellSortAnalysis(gap)
      if a == nil {
        return EdgeCaseResult{}
      }
      timesSum += t
    }
  default:
    log.Println(utils.RedColor("[!!] Nieznany algorytm sortowania: ", algorithm))
    return EdgeCaseResult{}
  }

  avg := timesSum / TestNum
  return EdgeCaseResult{Algorithm: algorithm, AvgTime: avg}

}

func TestSortedAscending(algorithm, pivot, gap string) EdgeCaseResult {
  var timesSum float64
  size := 100000

  switch algorithm {
  case "insertion":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateAscendingSortedArrayInt(size, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      _, t := sortAnalyzer.InsertionSortAnalysis()
      timesSum += t
    }
  case "heap":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateAscendingSortedArrayInt(size, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      _, t := sortAnalyzer.HeapSortAnalysis()
      timesSum += t
    }
  case "quick":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateAscendingSortedArrayInt(size, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      a, t := sortAnalyzer.QuickSortAnalysis(pivot)
      if a == nil {
        return EdgeCaseResult{}
      }
      timesSum += t
    }
  case "shell":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateAscendingSortedArrayInt(size, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      a, t := sortAnalyzer.ShellSortAnalysis(gap)
      if a == nil {
        return EdgeCaseResult{}
      }
      timesSum += t
    }
  default:
    log.Println(utils.RedColor("[!!] Nieznany algorytm sortowania: ", algorithm))
    return EdgeCaseResult{}
  }

  avg := timesSum / TestNum
  return EdgeCaseResult{Algorithm: algorithm, AvgTime: avg}
}

func TestSortedDescending(algorithm, pivot, gap string) EdgeCaseResult {
  var timesSum float64
  size := 100000

  switch algorithm {
  case "insertion":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateDescendingSortedArrayInt(size, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      _, t := sortAnalyzer.InsertionSortAnalysis()
      timesSum += t
    }
  case "heap":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateDescendingSortedArrayInt(size, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      _, t := sortAnalyzer.HeapSortAnalysis()
      timesSum += t
    }
  case "quick":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateDescendingSortedArrayInt(size, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      a, t := sortAnalyzer.QuickSortAnalysis(pivot)
      if a == nil {
        return EdgeCaseResult{}
      }
      timesSum += t
    }
  case "shell":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GenerateAscendingSortedArrayInt(size, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      a, t := sortAnalyzer.ShellSortAnalysis(gap)
      if a == nil {
        return EdgeCaseResult{}
      }
      timesSum += t
    }
  default:
    log.Println(utils.RedColor("[!!] Nieznany algorytm sortowania: ", algorithm))
    return EdgeCaseResult{}
  }

  avg := timesSum / TestNum
  return EdgeCaseResult{Algorithm: algorithm, AvgTime: avg}

}

func TestPartiallySorted(algorithm, pivot, gap string) EdgeCaseResult {
  var timesSum float64
  size := 100000

  switch algorithm {
  case "insertion":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GeneratePartiallySortedArrayInt(size, 33, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      _, t := sortAnalyzer.InsertionSortAnalysis()
      timesSum += t
    }
  case "heap":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GeneratePartiallySortedArrayInt(size, 33, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      _, t := sortAnalyzer.HeapSortAnalysis()
      timesSum += t
    }
  case "quick":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GeneratePartiallySortedArrayInt(size, 33, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      a, t := sortAnalyzer.QuickSortAnalysis(pivot)
      if a == nil {
        return EdgeCaseResult{}
      }
      timesSum += t
    }
  case "shell":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GeneratePartiallySortedArrayInt(size, 33, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      a, t := sortAnalyzer.ShellSortAnalysis(gap)
      if a == nil {
        return EdgeCaseResult{}
      }
      timesSum += t
    }
  default:
    log.Println(utils.RedColor("[!!] Nieznany algorytm sortowania: ", algorithm))
    return EdgeCaseResult{}
  }

  avg := timesSum / TestNum
  return EdgeCaseResult{Algorithm: algorithm, AvgTime: avg}

}
