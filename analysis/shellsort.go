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
func (sa *SortAnalyzer[T]) ShellSortAnalysis(gapCalculatingFormula string) ([]T, float64) {
  array := sa.GetDataCopy()

  log.Println("Typ danych:", utils.YellowColor(sa.DataTypyName))
  log.Println("Rozmiar Tablicy:", utils.YellowColor(sa.DataLength))
  if sa.DataLength > 20 {
    log.Println("Tablica (pierwsze 20 elementów):", utils.YellowColor(array[:20]))
  } else {
    log.Println("Tablica:", utils.YellowColor(array))
  }

  var calculateGap sort.GapCalculatingFormula
  switch gapCalculatingFormula {
  case "shell":
    gapCalculatingFormula, calculateGap = "Shella", sort.CalculateShellGap
  case "franklazarus":
    gapCalculatingFormula, calculateGap = "Franka i Lazarusa", sort.CalculateFrankLazarusGap
  case "hibbard":
    gapCalculatingFormula, calculateGap = "Hibbarda", sort.CalculateHibbardGap
  default:
    log.Println(utils.RedColor(fmt.Sprintf("[!!] Błąd: Nieznana formuła wyboru odstępu: %s. Dopuszczalne: shell, franklazarus, hibbard", gapCalculatingFormula)))
    return nil, 0.0
  }
 
  prompt := utils.BlueColor(fmt.Sprintf("Sortowanie Shella (wybór odstępu wg wzoru %s)", gapCalculatingFormula))
  log.Printf("[*] Rozpoczynanie: %s\n", prompt)
  start := time.Now()

  sort.ShellSort(array, calculateGap)

  timeElapsed := utils.PrintTimeElapsed(start, prompt)

  AssertSortedAscendingArray(array)

  return array, timeElapsed
}
