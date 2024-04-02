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
  case "generate":
    ParseGenerateOptions(os.Args[2:])
  default:
    log.Fatal(utils.RedColor(fmt.Sprintf("[!!] Nieznane polecenie: %v", os.Args[1])))
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
  fmt.Println("\tgenerate <typ danych> <rozmiar>\tgenerowanie plików z losowymi danymi wejściowymi")
  fmt.Println("\thelp\t\t\t\twyświetlanie tej wiadomości")

  fmt.Println()
  fmt.Println("Aby wyświetlić więcej informacji nt. danej komendy należy wpisać:")
  fmt.Printf("$ %v <cmd> -help\n", os.Args[0])
  fmt.Println()
  fmt.Println("Przykłady:")
  fmt.Printf("$ %v generate float32 100\n", os.Args[0])

  fmt.Println(exitCode)

  os.Exit(0)
}
