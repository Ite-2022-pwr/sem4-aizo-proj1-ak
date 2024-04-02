package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
	"golang.org/x/exp/constraints"
)

// Run przetwarza argumenty wiersza poleceń dla komendy run.
func Run() {
  // pozbycie się "run" z listy argumentów wiersza poleceń
  copy(os.Args[1:], os.Args[2:])
  os.Args = os.Args[:len(os.Args) - 1]

  flag.Usage = func() {
    fmt.Printf("Sposób użycia: %v run [options]\n\n", os.Args[0])

    flag.PrintDefaults()
  }

  // zdefiniowanie opcji programu dla komendy run
  help := flag.Bool("help", false, "wyświetlanie tej wiadomości")
  algorithm := flag.String("alg", "", "algorytm sortowania (insertion, heap, shell, quick) [WYMAGANY]")
  dataType := flag.String("type", "", "typ danych (int, int32, int64, float32, float64, byte, string) [WYMAGANY]")
  arraySize := flag.Int("generate", 1000, "ile liczb do wygenerowania (domyślnie 1000)")
  saveInput := flag.Bool("save", false, "czy zapisać wygenerowane dane do pliku")
  inputFile := flag.String("input-file", "", "czy pobrać dane wejściowe z pliku i jakiego (zamiast generate)")
  pivot := flag.String("pivot", "lewy", "jakiego pivota do sortowania szybkiego użyć (lewy, prawy, środkowy, losowy - domyślnie lewy)")
  gap := flag.String("gap", "shell", "jakiej metody obliczania odstępu w sortowaniu Shella użyć (shell, franklazarus, hibbard - domyślnie shell)")
  
  flag.Parse()

  if *help {
    flag.Usage()
    os.Exit(0)
  }

  if *algorithm == "" {
    log.Println(utils.RedColor("[!!] Nie podano algorytmu sortowania"))
    flag.Usage()
    os.Exit(1)
  }

  if *dataType == "" {
    log.Println(utils.RedColor("[!!] Nie podano typu danych"))
    flag.Usage()
    os.Exit(1)
  }

  algorithms := []string{"insertion", "heap", "shell", "quick"}
  if !slices.Contains(algorithms, *algorithm) {
    log.Println(utils.RedColor("[!!] Nieznany algorytm sortowania: ", *algorithm))
    flag.Usage()
    os.Exit(1)
  }

  switch *dataType {
  case "int":
    RunInt(algorithm, inputFile, pivot, gap, *arraySize, *saveInput)
  case "int32":
    RunInt32(algorithm, inputFile, pivot, gap, *arraySize, *saveInput)
  case "int64":
    RunInt64(algorithm, inputFile, pivot, gap, *arraySize, *saveInput)
  case "float32":
    RunFloat32(algorithm, inputFile, pivot, gap, *arraySize, *saveInput)
  case "float64":
    RunFloat64(algorithm, inputFile, pivot, gap, *arraySize, *saveInput)
  case "byte":
    RunByte(algorithm, inputFile, pivot, gap, *arraySize, *saveInput)
  case "string":
    RunString(algorithm, inputFile, pivot, gap, *arraySize, *saveInput)
  default:
    log.Println(utils.RedColor("[!!] Nieznany typ danych: ", *dataType))
    flag.Usage()
    os.Exit(1)
  }
}

// RunSort uruchamia wybrany algorytm sortowania.
func RunSort[T constraints.Ordered](alg, pivot, gap string, array []T) {
  sortAnalyzer := analysis.NewSortAnalyzer(array)
  switch alg {
  case "insertion":
    sortAnalyzer.InsertionSortAnalysis()
  case "heap":
    sortAnalyzer.HeapSortAnalysis()
  case "shell":
    sortAnalyzer.ShellSortAnalysis(gap)
  case "quick":
    sortAnalyzer.QuickSortAnalysis(pivot)
  default:
    log.Println(utils.RedColor("[!!] Nieznany algorytm sortowania:", alg))
  }
}

// RunInt uruchamia sortowanie dla typu int.
func RunInt(alg, inputFile, pivot, gap *string, arraySize int, save bool) {
  var array []int

  if *inputFile != "" {
    fileHandler := utils.NewInputFileHandler(*inputFile)
    array = fileHandler.TryParseInt()
  } else {
    array = generator.GenerateRandomArrayInt(arraySize, 1000, true)
    if save {
      utils.SaveArray(filepath.Join(InputFileDirectory, InputFileNamePrefix + fmt.Sprintf("int_%v", arraySize) + ".txt"), array)
    }
  }

  if array == nil {
    return
  }

  RunSort(*alg, *pivot, *gap, array)

}

// RunInt32 uruchamia sortowanie dla typu int32.
func RunInt32(alg, inputFile, pivot, gap *string, arraySize int, save bool) {
  var array []int32

  if *inputFile != "" {
    fileHandler := utils.NewInputFileHandler(*inputFile)
    array = fileHandler.TryParseInt32()
  } else {
    array = generator.GenerateRandomArrayInt32(arraySize, 1000, true)
    if save {
      utils.SaveArray(filepath.Join(InputFileDirectory, InputFileNamePrefix + fmt.Sprintf("int32_%v", arraySize) + ".txt"), array)
    }
  }

  if array == nil {
    return
  }

  RunSort(*alg, *pivot, *gap, array)

}

// RunInt64 uruchamia sortowanie dla typu int64.
func RunInt64(alg, inputFile, pivot, gap *string, arraySize int, save bool) {
  var array []int64

  if *inputFile != "" {
    fileHandler := utils.NewInputFileHandler(*inputFile)
    array = fileHandler.TryParseInt64()
  } else {
    array = generator.GenerateRandomArrayInt64(arraySize, 1000000, true)
    if save {
      utils.SaveArray(filepath.Join(InputFileDirectory, InputFileNamePrefix + fmt.Sprintf("int64_%v", arraySize) + ".txt"), array)
    }
  }

  if array == nil {
    return
  }

  RunSort(*alg, *pivot, *gap, array)

}

// RunFloat32 uruchamia sortowanie dla typu float32.
func RunFloat32(alg, inputFile, pivot, gap *string, arraySize int, save bool) {
  var array []float32

  if *inputFile != "" {
    fileHandler := utils.NewInputFileHandler(*inputFile)
    array = fileHandler.TryParseFloat32()
  } else {
    array = generator.GenerateRandomArrayFloat32(arraySize, 1000, true)
    if save {
      utils.SaveArray(filepath.Join(InputFileDirectory, InputFileNamePrefix + fmt.Sprintf("float32_%v", arraySize) + ".txt"), array)
    }
  }

  if array == nil {
    return
  }

  RunSort(*alg, *pivot, *gap, array)

}

// RunFloat64 uruchamia sortowanie dla typu float64.
func RunFloat64(alg, inputFile, pivot, gap *string, arraySize int, save bool) {
  var array []float64

  if *inputFile != "" {
    fileHandler := utils.NewInputFileHandler(*inputFile)
    array = fileHandler.TryParseFloat64()
  } else {
    array = generator.GenerateRandomArrayFloat64(arraySize, 1000000, true)
    if save {
      utils.SaveArray(filepath.Join(InputFileDirectory, InputFileNamePrefix + fmt.Sprintf("float64_%v", arraySize) + ".txt"), array)
    }
  }

  if array == nil {
    return
  }

  RunSort(*alg, *pivot, *gap, array)

}

// RunByte uruchamia sortowanie dla typu byte.
func RunByte(alg, inputFile, pivot, gap *string, arraySize int, save bool) {
  var array []byte

  if *inputFile != "" {
    fileHandler := utils.NewInputFileHandler(*inputFile)
    array = fileHandler.TryParseByte()
  } else {
    array = generator.GenerateRandomArrayByte(arraySize, true)
    if save {
      utils.SaveArray(filepath.Join(InputFileDirectory, InputFileNamePrefix + fmt.Sprintf("byte_%v", arraySize) + ".txt"), array)
    }
  }

  if array == nil {
    return
  }

  RunSort(*alg, *pivot, *gap, array)

}

// RunString uruchamia sortowanie dla typu string.
func RunString(alg, inputFile, pivot, gap *string, arraySize int, save bool) {
  var array []string

  if *inputFile != "" {
    fileHandler := utils.NewInputFileHandler(*inputFile)
    array = fileHandler.GetDataCopy()
  } else {
    array = generator.GenerateRandomArrayString(arraySize, 20, true)
    if save {
      utils.SaveArray(filepath.Join(InputFileDirectory, InputFileNamePrefix + fmt.Sprintf("string_%v", arraySize) + ".txt"), array)
    }
  }

  if array == nil {
    return
  }

  RunSort(*alg, *pivot, *gap, array)

}
