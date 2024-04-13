package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// EdgeCaseResult reprezentuje uśredniony wynik pomiarów przypadku szczególnego dla danego algorytmu sortowania.
type EdgeCaseResult struct {
  Algorithm string
  AvgTime  float64
}

// ToStrings zwraca wartości pól struktury EdgeCaseResult w postaci []string.
func (ecr EdgeCaseResult) ToStrings() []string {
  return []string{ecr.Algorithm, fmt.Sprintf("%.3f", ecr.AvgTime)}
}

// TestEdgeCases sprawdza przypadki szczególne.
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
  results = append(results, TestPartiallySorted33("heap", "", "").ToStrings())
  results = append(results, TestPartiallySorted33("quick", "lewy", "").ToStrings())
  results = append(results, TestPartiallySorted33("quick", "prawy", "").ToStrings())
  results = append(results, TestPartiallySorted33("quick", "środkowy", "").ToStrings())
  results = append(results, TestPartiallySorted33("quick", "losowy", "").ToStrings())
  results = append(results, TestPartiallySorted33("shell", "", "shell").ToStrings())
  results = append(results, TestPartiallySorted33("shell", "", "franklazarus").ToStrings())
  results = append(results, TestPartiallySorted33("insertion", "", "").ToStrings())

  utils.SaveCSV(filepath.Join(OutputDirectory, "edge_cases", "partially33.csv"), results)

  results = nil

  // Częściowo posortowana (w 66%)
  results = append(results, TestPartiallySorted66("heap", "", "").ToStrings())
  results = append(results, TestPartiallySorted66("quick", "lewy", "").ToStrings())
  results = append(results, TestPartiallySorted66("quick", "prawy", "").ToStrings())
  results = append(results, TestPartiallySorted66("quick", "środkowy", "").ToStrings())
  results = append(results, TestPartiallySorted66("quick", "losowy", "").ToStrings())
  results = append(results, TestPartiallySorted66("shell", "", "shell").ToStrings())
  results = append(results, TestPartiallySorted66("shell", "", "franklazarus").ToStrings())
  results = append(results, TestPartiallySorted66("insertion", "", "").ToStrings())

  utils.SaveCSV(filepath.Join(OutputDirectory, "edge_cases", "partially66.csv"), results)

  results = nil
}

// TestFullRandom mierzy czas tablicy z w pełni losowym uporządkowaniem elementów.
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
    algorithm += "_" + pivot
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
    algorithm += "_" + gap
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

// TestSortedAscending mierzy czasy sortowania dla tablicy posortowanej rosnąco
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
    algorithm += "_" + pivot
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
    algorithm += "_" + gap
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

// TestSortedDescending mierzy czas sortowania tablicy posortowanej malejąco.
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
    algorithm += "_" + pivot
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
    algorithm += "_" + gap
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

// TestPartiallySorted33 mierzy czas sortowania tablicy posortowanej rosnąco w 33%.
func TestPartiallySorted33(algorithm, pivot, gap string) EdgeCaseResult {
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
    algorithm += "_" + pivot
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
    algorithm += "_" + gap
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

// TestPartiallySorted66 mierzy czas sortowania tablicy posortowanej rosnąco w 66%.
func TestPartiallySorted66(algorithm, pivot, gap string) EdgeCaseResult {
  var timesSum float64
  size := 100000

  switch algorithm {
  case "insertion":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GeneratePartiallySortedArrayInt(size, 66, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      _, t := sortAnalyzer.InsertionSortAnalysis()
      timesSum += t
    }
  case "heap":
    for i := 0; i < int(TestNum); i++ {
      array := generator.GeneratePartiallySortedArrayInt(size, 66, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      _, t := sortAnalyzer.HeapSortAnalysis()
      timesSum += t
    }
  case "quick":
    algorithm += "_" + pivot
    for i := 0; i < int(TestNum); i++ {
      array := generator.GeneratePartiallySortedArrayInt(size, 66, true)
      sortAnalyzer := analysis.NewSortAnalyzer(array)
      log.Println(utils.BlueColor(fmt.Sprintf("[*] Pomiar %v/%v", i + 1, TestNum)))
      a, t := sortAnalyzer.QuickSortAnalysis(pivot)
      if a == nil {
        return EdgeCaseResult{}
      }
      timesSum += t
    }
  case "shell":
    algorithm += "_" + gap
    for i := 0; i < int(TestNum); i++ {
      array := generator.GeneratePartiallySortedArrayInt(size, 66, true)
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
