package cmd

import (
	"fmt"
	"log"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// VerifySorting przeprowadza weryfikację poprawności działania zaimplementowanych
// algorytmów sortujących.
func VerifySorting() {
  log.Println(utils.YellowColor("[*] Weryfikacja poprawności działania algorytmów sortowania"))
  fmt.Println()

  array := generator.GenerateRandomArrayInt(15, 10, true)

  sortAnalyzer := analysis.NewSortAnalyzer(array)

  sortAnalyzer.InsertionSortAnalysis()
  fmt.Println()

  sortAnalyzer.HeapSortAnalysis()
  fmt.Println()

  sortAnalyzer.ShellSortAnalysis("shell")
  fmt.Println()

  sortAnalyzer.ShellSortAnalysis("franklazarus")
  fmt.Println()

  sortAnalyzer.QuickSortAnalysis("lewy")
  fmt.Println()

  sortAnalyzer.QuickSortAnalysis("prawy")
  fmt.Println()

  sortAnalyzer.QuickSortAnalysis("środkowy")
  fmt.Println()

  sortAnalyzer.QuickSortAnalysis("losowy")
  fmt.Println()

}
