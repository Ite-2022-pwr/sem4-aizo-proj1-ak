// Package cmd przetwarza argumenty, z jakimi został uruchomiony program.
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

const (
  // InputFileDirectory folder do przechowywania wygenerowanych plików
  InputFileDirectory = "data/input"

  // InputFileNamePrefix to prefiks nazwy pliku z wygenerowanymi danymi
  InputFileNamePrefix = "input_"
)

// GenerateInputFile generuje plik z losowymi danymi wejściowymi typu dataType o ilości liczb size.
// Parametr dataType przyjmuje wartości: "int", "int32", "int64", "float32", "float64", "byte" lub "string".
// Parametr size przyjmuje wartości dodatnie.
func GenerateInputFile(dataType string, size int) {
  filename := InputFileNamePrefix + dataType + fmt.Sprintf("_%v", size) + ".txt"
  path := filepath.Join(InputFileDirectory, filename)

  if size < 1 {
    log.Fatal(utils.RedColor("[!!] Rozmiar generowanej tablicy nie musi być dodatni"))
  }

  switch dataType {
  case "int":
    data := generator.GenerateRandomArrayInt(size, 10000, true)
    utils.SaveArray(path, data)
  case "int32":
    data := generator.GenerateRandomArrayInt32(size, 10000, true)
    utils.SaveArray(path, data)
  case "int64":
    data := generator.GenerateRandomArrayInt64(size, 100000, true)
    utils.SaveArray(path, data)
  case "float32":
    data := generator.GenerateRandomArrayFloat32(size, true)
    utils.SaveArray(path, data)
  case "float64":
    data := generator.GenerateRandomArrayFloat64(size, true)
    utils.SaveArray(path, data)
  case "byte":
    data := generator.GenerateRandomArrayByte(size, true)
    utils.SaveArray(path, data)
  case "string":
    data := generator.GenerateRandomArrayString(size, 20, true)
    utils.SaveArray(path, data)
  default:
    log.Fatal(utils.RedColor(fmt.Sprintf("[!!] Nieznany typ danych: %v. Dopuszczalne: int, int32, int64, float32, float64, byte lub string", dataType)))
  }
}

// PrintGenerateHelp wypisuje pomoc i sposób użycia programu z komendą generate.
func PrintGenerateHelp(exitCode int) {
  fmt.Printf("$ %v generate <typ danych> <rozmiar>\n", os.Args[0])
  fmt.Println()

  os.Exit(exitCode)
}
