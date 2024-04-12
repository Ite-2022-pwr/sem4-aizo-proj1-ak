// Package main stanowi punkt wejścia do całego programu
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/cmd"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

func main() {
  fmt.Println(utils.MagentaColor("AZO - zadanie projektowe nr 1"))
  fmt.Println((utils.MagentaColor("Badanie efektywności wybranych algorytmów sortowania ze względu na złożoność obliczeniową")))
  fmt.Println(utils.BlueColor("Autor: Artur Kręgiel, 2024r."))
  fmt.Println()

  if len(os.Args) < 2 {
    PrintGeneralHelp(1)
  }

  switch os.Args[1] {
  case "help":
    PrintGeneralHelp(0)
  case "verify":
    cmd.VerifySorting()
  case "generate":
    ParseGenerateOptions(os.Args[2:])
  case "run":
    cmd.Run()
  case "interactive":
    cmd.Menu()
  case "sizes":
    cmd.MeasureTimesSizeDependend()
  case "types":
    cmd.AnalyzeDifferentDataTypesHeapsort()
  case "edge":
    cmd.TestEdgeCases()
  default:
    log.Println(utils.RedColor(fmt.Sprintf("[!!] Nieznane polecenie: %v", os.Args[1])))
    PrintGeneralHelp(1)
  }
}

// ParseGenerateOptions przetwarza argumenty polecenia generate
func ParseGenerateOptions(args []string) {
  if len(args) != 2 {
    cmd.PrintGenerateHelp(1)
  }

  dataType := args[0]
  n, err := strconv.Atoi(args[1])
  if err != nil {
    log.Fatal(utils.RedColor(fmt.Sprintf("[!!] Błąd konwersji: %v", err)))
  }

  cmd.GenerateInputFile(dataType, n)
}

// PrintGeneralHelp wypisuje generalną pomoc i sposób użycia programu.
func PrintGeneralHelp(exitCode int) {
  fmt.Printf("Sposób użycia:\n$ %v <cmd> [options]\n\n", os.Args[0])
  fmt.Println("\tverify\t\t\t\tweryfikacja poprawności zaimplementowanych algorytmów sortowania")
  fmt.Println("\tgenerate <typ danych> <rozmiar>\tgenerowanie plików z losowymi danymi wejściowymi")
  fmt.Println("\trun [options]\t\t\turuchamianie wybranego algorytmu sortowania")
  fmt.Println("\tinteractive\t\t\ttryb interaktywny (menu)")
  fmt.Println("\thelp\t\t\t\twyświetlanie tej wiadomości")

  fmt.Println()
  fmt.Println("Przykłady:")
  fmt.Printf("$ %v generate float32 100\n", os.Args[0])
  fmt.Printf("$ %v run -alg=quick -type=string -generate=100000 -pivot=losowy\n", os.Args[0])
  fmt.Printf("$ %v run -alg=quick -type=string -pivot=prawy -input-file=data/input/input_int_1000000.txt\n", os.Args[0])
  fmt.Println()

  os.Exit(exitCode)
}
