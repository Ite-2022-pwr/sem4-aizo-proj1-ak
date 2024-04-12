package cmd

import (
	"fmt"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/analysis"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/generator"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
	"golang.org/x/exp/constraints"
)

// Menu to główne menu trybu interaktywnego
func Menu() {
  for {
    fmt.Println("Menu:")
    fmt.Println("Wybierz typ danych:")
    fmt.Println("0 - Wyjście z programu")
    fmt.Println("1 - int")
    fmt.Println("2 - int32")
    fmt.Println("3 - int64")
    fmt.Println("4 - float32")
    fmt.Println("5 - float64")
    fmt.Println("6 - byte")
    fmt.Println("7 - string")
    fmt.Print("Wybór: ")

    var choice int
    if _, e := fmt.Scanf("%d", &choice); e != nil {
      fmt.Println(utils.RedColor("Błąd: ", e))
      continue
    }

    switch choice {
    case 0:
      return
    case 1:
      MenuInt()
    case 2:
      MenuInt32()
    case 3:
      MenuInt64()
    case 4:
      MenuFloat32()
    case 5:
      MenuFloat64()
    case 6:
      MenuByte()
    case 7:
      MenuString()
    default:
      fmt.Println("Wybierz numer od 0 do 7")
    }
  }
}

func chooseGap() string {
  for {
    fmt.Println("Wybierz metodę wyboru odstępu:")
    fmt.Println("0 - Shella")
    fmt.Println("1 - Franka i Lazarusa")
    fmt.Print("Wybór: ")

    var choice int
    if _, e := fmt.Scanf("%d", &choice); e != nil {
      fmt.Println(utils.RedColor("Błąd: ", e))
      continue
    }
    switch choice {
    case 0:
      return "shell"
    case 1:
      return "franklazarus"
    default:
      fmt.Println(utils.RedColor("Wybierz 0 lub 1"))
    }
  }
}

func choosePivot() string {
  for {
    fmt.Println("Wybierz pivot:")
    fmt.Println("0 - Skrajny lewy")
    fmt.Println("1 - Skrajny prawy")
    fmt.Println("2 - Środkowy")
    fmt.Println("3 - Losowy")
    fmt.Print("Wybór: ")

    var choice int
    if _, e := fmt.Scanf("%d", &choice); e != nil {
      fmt.Println(utils.RedColor("Błąd: ", e))
      continue
    }
    switch choice {
    case 0:
      return "lewy"
    case 1:
      return "prawy"
    case 2:
      return "środkowy"
    case 3:
      return "losowy"
    default:
      fmt.Println(utils.RedColor("Wybierz numer od 0 do 3"))
    }
  }
}

// MenuSort to menu z wyborem algorytmu sortującego
func MenuSort[T constraints.Ordered](sortAnalyzer *analysis.SortAnalyzer[T]) []T {
  if sortAnalyzer == nil {
    return nil
  }

  var sortedArray []T

  for {
    fmt.Println("Wybierz algorytm sortowania:")
    fmt.Println("0 - Sortowanie przez wstawianie")
    fmt.Println("1 - Sortowanie przez kopcowanie")
    fmt.Println("2 - Sortowanie Shella")
    fmt.Println("3 - Sortowanie szybkie")
    fmt.Print("Wybór: ")
    var choice int
    if _, e := fmt.Scanf("%d", &choice); e != nil {
      fmt.Println(utils.RedColor("Błąd: ", e))
      continue
    }
    switch choice {
    case 0:
      sortedArray, _ = sortAnalyzer.InsertionSortAnalysis()
      return sortedArray
    case 1:
      sortedArray, _ = sortAnalyzer.HeapSortAnalysis()
      return sortedArray
    case 2:
      gap := chooseGap()
      sortedArray, _ = sortAnalyzer.ShellSortAnalysis(gap)
      return sortedArray
    case 3:
      pivot := choosePivot()
      sortedArray, _ = sortAnalyzer.QuickSortAnalysis(pivot)
      return sortedArray
    default:
      fmt.Println(utils.RedColor("Wybierz numer od 0 do 3"))
    }
  }
}

// MenuInt to menu dla typu danych int
func MenuInt() {
  var sortAnalyzer *analysis.SortAnalyzer[int]
  var sortedArray []int

  for {
    fmt.Println("Wybierz opcję:")
    fmt.Println("0 - Wyjście do menu głównego")
    fmt.Println("1 - Wygeneruj nową tablicę [int]")
    fmt.Println("2 - Wczytaj dane z pliku")
    fmt.Println("3 - Wyświetl tablicę przed posortowaniem")
    fmt.Println("4 - Sortuj tablicę (rosnąco)")
    fmt.Println("5 - Wyświetl tablicę po posortowaniu")
    fmt.Print("Wybór: ")

    var choice int
    if _, e := fmt.Scanf("%d", &choice); e != nil {
      fmt.Println(utils.RedColor("Błąd: ", e))
      continue
    }
    switch choice {
    case 0:
      return
    case 1:
      var size int
      for {
        fmt.Print("Podaj rozmiar tablicy: ")
        if _, e := fmt.Scanf("%d", &size); e != nil {
          fmt.Println(utils.RedColor("Błąd: ", e))
          continue
        }
        if size <= 0 {
          fmt.Println("Rozmiar tablicy musi być liczbą dodatnią")
          continue
        }
        break
      }
      array := generator.GenerateRandomArrayInt(size, 1000, true)
      sortAnalyzer = analysis.NewSortAnalyzer(array)
    case 2:
      fmt.Print("Podaj nazwę pliku: ")
      var filename string
      fmt.Scanf("%s", &filename)
      fileHandler := utils.NewInputFileHandler(filename)
      if array := fileHandler.TryParseInt(); array != nil {
        sortAnalyzer = analysis.NewSortAnalyzer(array)
      }
    case 3:
      if sortAnalyzer != nil {
        fmt.Println(sortAnalyzer.Data)
      }
    case 4:
      sortedArray = MenuSort(sortAnalyzer)
    case 5:
      fmt.Println(sortedArray)
    default:
      fmt.Println(utils.RedColor("Wybierz numer od 0 do 5"))
    }
  }
}

// MenuInt32 to menu dla typu danych int32
func MenuInt32() {
  var sortAnalyzer *analysis.SortAnalyzer[int32]
  var sortedArray []int32

  for {
    fmt.Println("Wybierz opcję:")
    fmt.Println("0 - Wyjście do menu głównego")
    fmt.Println("1 - Wygeneruj nową tablicę [int32]")
    fmt.Println("2 - Wczytaj dane z pliku")
    fmt.Println("3 - Wyświetl tablicę przed posortowaniem")
    fmt.Println("4 - Sortuj tablicę (rosnąco)")
    fmt.Println("5 - Wyświetl tablicę po posortowaniu")
    fmt.Print("Wybór: ")

    var choice int
    if _, e := fmt.Scanf("%d", &choice); e != nil {
      fmt.Println(utils.RedColor("Błąd: ", e))
      continue
    }
    switch choice {
    case 0:
      return
    case 1:
      var size int
      for {
        fmt.Print("Podaj rozmiar tablicy: ")
        if _, e := fmt.Scanf("%d", &size); e != nil {
          fmt.Println(utils.RedColor("Błąd: ", e))
          continue
        }
        if size <= 0 {
          fmt.Println("Rozmiar tablicy musi być liczbą dodatnią")
          continue
        }
        break
      }
      array := generator.GenerateRandomArrayInt32(size, 1000, true)
      sortAnalyzer = analysis.NewSortAnalyzer(array)
    case 2:
      fmt.Print("Podaj nazwę pliku: ")
      var filename string
      fmt.Scanf("%s", &filename)
      fileHandler := utils.NewInputFileHandler(filename)
      if array := fileHandler.TryParseInt32(); array != nil {
        sortAnalyzer = analysis.NewSortAnalyzer(array)
      }
    case 3:
      if sortAnalyzer != nil {
        fmt.Println(sortAnalyzer.Data)
      }
    case 4:
      sortedArray = MenuSort(sortAnalyzer)
    case 5:
      fmt.Println(sortedArray)
    default:
      fmt.Println(utils.RedColor("Wybierz numer od 0 do 5"))
    }
  }
}

// MenuInt64 to menu dla typu danych int64
func MenuInt64() {
  var sortAnalyzer *analysis.SortAnalyzer[int64]
  var sortedArray []int64

  for {
    fmt.Println("Wybierz opcję:")
    fmt.Println("0 - Wyjście do menu głównego")
    fmt.Println("1 - Wygeneruj nową tablicę [int64]")
    fmt.Println("2 - Wczytaj dane z pliku")
    fmt.Println("3 - Wyświetl tablicę przed posortowaniem")
    fmt.Println("4 - Sortuj tablicę (rosnąco)")
    fmt.Println("5 - Wyświetl tablicę po posortowaniu")
    fmt.Print("Wybór: ")

    var choice int
    if _, e := fmt.Scanf("%d", &choice); e != nil {
      fmt.Println(utils.RedColor("Błąd: ", e))
      continue
    }
    switch choice {
    case 0:
      return
    case 1:
      var size int
      for {
        fmt.Print("Podaj rozmiar tablicy: ")
        if _, e := fmt.Scanf("%d", &size); e != nil {
          fmt.Println(utils.RedColor("Błąd: ", e))
          continue
        }
        if size <= 0 {
          fmt.Println("Rozmiar tablicy musi być liczbą dodatnią")
          continue
        }
        break
      }
      array := generator.GenerateRandomArrayInt64(size, 1000, true)
      sortAnalyzer = analysis.NewSortAnalyzer(array)
    case 2:
      fmt.Print("Podaj nazwę pliku: ")
      var filename string
      fmt.Scanf("%s", &filename)
      fileHandler := utils.NewInputFileHandler(filename)
      if array := fileHandler.TryParseInt64(); array != nil {
        sortAnalyzer = analysis.NewSortAnalyzer(array)
      }
    case 3:
      if sortAnalyzer != nil {
        fmt.Println(sortAnalyzer.Data)
      }
    case 4:
      sortedArray = MenuSort(sortAnalyzer)
    case 5:
      fmt.Println(sortedArray)
    default:
      fmt.Println(utils.RedColor("Wybierz numer od 0 do 5"))
    }
  }
}

// MenuFloat32 to menu dla typu danych float32
func MenuFloat32() {
  var sortAnalyzer *analysis.SortAnalyzer[float32]
  var sortedArray []float32

  for {
    fmt.Println("Wybierz opcję:")
    fmt.Println("0 - Wyjście do menu głównego")
    fmt.Println("1 - Wygeneruj nową tablicę [float32]")
    fmt.Println("2 - Wczytaj dane z pliku")
    fmt.Println("3 - Wyświetl tablicę przed posortowaniem")
    fmt.Println("4 - Sortuj tablicę (rosnąco)")
    fmt.Println("5 - Wyświetl tablicę po posortowaniu")
    fmt.Print("Wybór: ")

    var choice int
    if _, e := fmt.Scanf("%d", &choice); e != nil {
      fmt.Println(utils.RedColor("Błąd: ", e))
      continue
    }
    switch choice {
    case 0:
      return
    case 1:
      var size int
      for {
        fmt.Print("Podaj rozmiar tablicy: ")
        if _, e := fmt.Scanf("%d", &size); e != nil {
          fmt.Println(utils.RedColor("Błąd: ", e))
          continue
        }
        if size <= 0 {
          fmt.Println("Rozmiar tablicy musi być liczbą dodatnią")
          continue
        }
        break
      }
      array := generator.GenerateRandomArrayFloat32(size, 1000, true)
      sortAnalyzer = analysis.NewSortAnalyzer(array)
    case 2:
      fmt.Print("Podaj nazwę pliku: ")
      var filename string
      fmt.Scanf("%s", &filename)
      fileHandler := utils.NewInputFileHandler(filename)
      if array := fileHandler.TryParseFloat32(); array != nil {
        sortAnalyzer = analysis.NewSortAnalyzer(array)
      }
    case 3:
      if sortAnalyzer != nil {
        fmt.Println(sortAnalyzer.Data)
      }
    case 4:
      sortedArray = MenuSort(sortAnalyzer)
    case 5:
      fmt.Println(sortedArray)
    default:
      fmt.Println(utils.RedColor("Wybierz numer od 0 do 5"))
    }
  }
}

// MenuFloat64 to menu dla typu danych float64
func MenuFloat64() {
  var sortAnalyzer *analysis.SortAnalyzer[float64]
  var sortedArray []float64

  for {
    fmt.Println("Wybierz opcję:")
    fmt.Println("0 - Wyjście do menu głównego")
    fmt.Println("1 - Wygeneruj nową tablicę [float64]")
    fmt.Println("2 - Wczytaj dane z pliku")
    fmt.Println("3 - Wyświetl tablicę przed posortowaniem")
    fmt.Println("4 - Sortuj tablicę (rosnąco)")
    fmt.Println("5 - Wyświetl tablicę po posortowaniu")
    fmt.Print("Wybór: ")

    var choice int
    if _, e := fmt.Scanf("%d", &choice); e != nil {
      fmt.Println(utils.RedColor("Błąd: ", e))
      continue
    }
    switch choice {
    case 0:
      return
    case 1:
      var size int
      for {
        fmt.Print("Podaj rozmiar tablicy: ")
        if _, e := fmt.Scanf("%d", &size); e != nil {
          fmt.Println(utils.RedColor("Błąd: ", e))
          continue
        }
        if size <= 0 {
          fmt.Println("Rozmiar tablicy musi być liczbą dodatnią")
          continue
        }
        break
      }
      array := generator.GenerateRandomArrayFloat64(size, 1000, true)
      sortAnalyzer = analysis.NewSortAnalyzer(array)
    case 2:
      fmt.Print("Podaj nazwę pliku: ")
      var filename string
      fmt.Scanf("%s", &filename)
      fileHandler := utils.NewInputFileHandler(filename)
      if array := fileHandler.TryParseFloat64(); array != nil {
        sortAnalyzer = analysis.NewSortAnalyzer(array)
      }
    case 3:
      if sortAnalyzer != nil {
        fmt.Println(sortAnalyzer.Data)
      }
    case 4:
      sortedArray = MenuSort(sortAnalyzer)
    case 5:
      fmt.Println(sortedArray)
    default:
      fmt.Println(utils.RedColor("Wybierz numer od 0 do 5"))
    }
  }
}

// MenuByte to menu dla typu danych byte (uint8)
func MenuByte() {
  var sortAnalyzer *analysis.SortAnalyzer[byte]
  var sortedArray []byte

  for {
    fmt.Println("Wybierz opcję:")
    fmt.Println("0 - Wyjście do menu głównego")
    fmt.Println("1 - Wygeneruj nową tablicę [byte]")
    fmt.Println("2 - Wczytaj dane z pliku")
    fmt.Println("3 - Wyświetl tablicę przed posortowaniem")
    fmt.Println("4 - Sortuj tablicę (rosnąco)")
    fmt.Println("5 - Wyświetl tablicę po posortowaniu")
    fmt.Print("Wybór: ")

    var choice int
    if _, e := fmt.Scanf("%d", &choice); e != nil {
      fmt.Println(utils.RedColor("Błąd: ", e))
      continue
    }
    switch choice {
    case 0:
      return
    case 1:
      var size int
      for {
        fmt.Print("Podaj rozmiar tablicy: ")
        if _, e := fmt.Scanf("%d", &size); e != nil {
          fmt.Println(utils.RedColor("Błąd: ", e))
          continue
        }
        if size <= 0 {
          fmt.Println("Rozmiar tablicy musi być liczbą dodatnią")
          continue
        }
        break
      }
      array := generator.GenerateRandomArrayByte(size, true)
      sortAnalyzer = analysis.NewSortAnalyzer(array)
    case 2:
      fmt.Print("Podaj nazwę pliku: ")
      var filename string
      fmt.Scanf("%s", &filename)
      fileHandler := utils.NewInputFileHandler(filename)
      if array := fileHandler.TryParseByte(); array != nil {
        sortAnalyzer = analysis.NewSortAnalyzer(array)
      }
    case 3:
      if sortAnalyzer != nil {
        fmt.Println(sortAnalyzer.Data)
      }
    case 4:
      sortedArray = MenuSort(sortAnalyzer)
    case 5:
      fmt.Println(sortedArray)
    default:
      fmt.Println(utils.RedColor("Wybierz numer od 0 do 5"))
    }
  }
}

// MenuString to menu dla typu danych string
func MenuString() {
  var sortAnalyzer *analysis.SortAnalyzer[string]
  var sortedArray []string

  for {
    fmt.Println("Wybierz opcję:")
    fmt.Println("0 - Wyjście do menu głównego")
    fmt.Println("1 - Wygeneruj nową tablicę [string]")
    fmt.Println("2 - Wczytaj dane z pliku")
    fmt.Println("3 - Wyświetl tablicę przed posortowaniem")
    fmt.Println("4 - Sortuj tablicę (rosnąco)")
    fmt.Println("5 - Wyświetl tablicę po posortowaniu")
    fmt.Print("Wybór: ")

    var choice int
    if _, e := fmt.Scanf("%d", &choice); e != nil {
      fmt.Println(utils.RedColor("Błąd: ", e))
      continue
    }
    switch choice {
    case 0:
      return
    case 1:
      var size int
      for {
        fmt.Print("Podaj rozmiar tablicy: ")
        if _, e := fmt.Scanf("%d", &size); e != nil {
          fmt.Println(utils.RedColor("Błąd: ", e))
          continue
        }
        if size <= 0 {
          fmt.Println("Rozmiar tablicy musi być liczbą dodatnią")
          continue
        }
        break
      }
      array := generator.GenerateRandomArrayString(size, 20, true)
      sortAnalyzer = analysis.NewSortAnalyzer(array)
    case 2:
      fmt.Print("Podaj nazwę pliku: ")
      var filename string
      fmt.Scanf("%s", &filename)
      fileHandler := utils.NewInputFileHandler(filename)
      if array := fileHandler.GetDataCopy(); array != nil {
        sortAnalyzer = analysis.NewSortAnalyzer(array)
      }
    case 3:
      if sortAnalyzer != nil {
        fmt.Println(sortAnalyzer.Data)
      }
    case 4:
      sortedArray = MenuSort(sortAnalyzer)
    case 5:
      fmt.Println(sortedArray)
    default:
      fmt.Println(utils.RedColor("Wybierz numer od 0 do 5"))
    }
  }
}
