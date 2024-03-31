package analysis

import (
	"fmt"
	"log"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/sort"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// ShellSortAnalysis mierzy czas sortowania Shella.
// Parametr gapCalculatingFormula musi przyjmować jedną z wartości: "shell", "franklazarus" lub "hibbard". 
func (sa *SortAnalyzer[T]) ShellSortAnalysis(gapCalculatingFormula string) []T {
  array := sa.GetDataCopy()

  var calculateGap sort.GapCalculatingFormula
  switch gapCalculatingFormula {
  case "shell":
    gapCalculatingFormula, calculateGap = "Shella", sort.CalculateShellGap
  case "franklazarus":
    gapCalculatingFormula, calculateGap = "Franka i Lazarusa", sort.CalculateFrankLazarusGap
  case "hibbard":
    gapCalculatingFormula, calculateGap = "Hibbarda", sort.CalculateHibbardGap
  default:
    log.Fatalf("[!!] Błąd: Nieznana formuła wyboru odstępu: %s. Dopuszczalne: shell, franklazarus, hibbard", gapCalculatingFormula)
  }
 
  start := time.Now()
  prompt := fmt.Sprintf("Sortowanie Shella (wybór odstępu wg wzoru %s) dla typu danych %T", gapCalculatingFormula, array)
  log.Printf("[*] Rozpoczynanie: %s\n", prompt)

  sort.ShellSort(array, calculateGap)

  utils.PrintTimeElapsed(start, prompt)

  if IsArraySorted(array) {
    log.Println("[+] Tablica posortowana poprawnie")
  } else {
    log.Fatal("[!!] Tablica nie została posortowana poprawnie!")
  }

  return array
}
