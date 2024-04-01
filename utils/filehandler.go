package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// SaveArray zapisuje tablicę do pliku
func SaveArray[T constraints.Ordered](filepath string, array []T) {
  f, err := os.Create(filepath)
  if err != nil {
    log.Fatal(RedColor(fmt.Sprintf("[!!] Nie można utworzyć pliku: %v", err)))
  }
  defer f.Close()

  wrtr := bufio.NewWriter(f)

  n := len(array)
  line := fmt.Sprintf("%v\n", n)
  _, err = wrtr.WriteString(line)
  if err != nil {
    log.Fatal(RedColor(fmt.Sprintf("[!!] Nie można zapisać do pliku: %v", err)))
  }

  for i := 0; i < n; i++ {
    line = fmt.Sprintf("%v\n", array[i])
    _, err = wrtr.WriteString(line)
    if err != nil {
      log.Fatal(RedColor(fmt.Sprintf("[!!] Nie można zapisać do pliku: %v", err)))
    }
    wrtr.Flush()
  }
}

// InputFileHandler odpowiada za wczytanie i parsowanie pliku z danymi wejściowymi
type InputFileHandler struct {
  FilePath string
  Data []string
  DataLength int
}

// NewInputFileHandler otwiera i czyta plik wejściowy.
// Zwraca obiekt InputFileHandler
func NewInputFileHandler(filepath string) *InputFileHandler {
  fh := &InputFileHandler{FilePath: filepath}
  fh.LoadData()
  return fh
}

// LoadData wczytuje dane z pliku InputFileHandler.FilePath
func (fh *InputFileHandler) LoadData() {
  f, err := os.Open(fh.FilePath)
  if err != nil {
    log.Fatal(RedColor(fmt.Sprintf("[!!] Nie można otworzyć pliku: %v", err)))
  }
  defer f.Close()

  rdr := bufio.NewReader(f)
  var line string
  line, err = rdr.ReadString('\n')
  if err != nil {
    log.Fatal(RedColor(fmt.Sprintf("[!!] Błąd czytania linii '%v': %v", line, err)))
  }
  line = strings.TrimSpace(line)

  var arraySize int
  arraySize, err = strconv.Atoi(line)
  if err != nil {
    log.Fatal(RedColor(fmt.Sprintf("[!!] Błąd konwersji '%v': %v", line, err)))
  }

  fh.Data = make([]string, arraySize)
  fh.DataLength = arraySize

  for i := 0; i < arraySize; i++ {
    line, err = rdr.ReadString('\n')
    if err != nil {
      log.Fatal(RedColor(fmt.Sprintf("[!!] Błąd czytania linii '%v': %v", line, err)))
    }
    fh.Data[i] = strings.TrimSpace(line)
  }
}

// TryParseInt próbuje przekonwertować dane z tekstu na int.
func (fh *InputFileHandler) TryParseInt() []int {
  array := make([]int, fh.DataLength)

  for i := 0; i < fh.DataLength; i++ {
    num, err := strconv.Atoi(fh.Data[i])
    if err != nil {
      log.Fatal(RedColor(fmt.Sprintf("[!!] Błąd konwersji '%v': %v", fh.Data[i], err)))
    }
    array[i] = num
  }

  return array
}

// TryParseInt32 próbuje przekonwertować dane z tekstu na int32.
func (fh *InputFileHandler) TryParseInt32() []int32 {
  array := make([]int32, fh.DataLength)

  for i := 0; i < fh.DataLength; i++ {
    num, err := strconv.ParseInt(fh.Data[i], 10, 32)
    if err != nil {
      log.Fatal(RedColor(fmt.Sprintf("[!!] Błąd konwersji '%v': %v", fh.Data[i], err)))
    }
    array[i] = int32(num)
  }

  return array
}

// TryParseInt64 próbuje przekonwertować dane z tekstu na int64.
func (fh *InputFileHandler) TryParseInt64() []int64 {
  array := make([]int64, fh.DataLength)

  for i := 0; i < fh.DataLength; i++ {
    num, err := strconv.ParseInt(fh.Data[i], 10, 64)
    if err != nil {
      log.Fatal(RedColor(fmt.Sprintf("[!!] Błąd konwersji '%v': %v", fh.Data[i], err)))
    }
    array[i] = num
  }

  return array
}

// TryParseFloat32 próbuje przekonwertować dane z tekstu na float32.
func (fh *InputFileHandler) TryParseFloat32() []float32 {
  array := make([]float32, fh.DataLength)

  for i := 0; i < fh.DataLength; i++ {
    num, err := strconv.ParseFloat(fh.Data[i], 32)
    if err != nil {
      log.Fatal(RedColor(fmt.Sprintf("[!!] Błąd konwersji '%v': %v", fh.Data[i], err)))
    }
    array[i] = float32(num)
  }

  return array
}

// TryParseFloat64 próbuje przekonwertować dane z tekstu na float64.
func (fh *InputFileHandler) TryParseFloat64() []float64 {
  array := make([]float64, fh.DataLength)

  for i := 0; i < fh.DataLength; i++ {
    num, err := strconv.ParseFloat(fh.Data[i], 64)
    if err != nil {
      log.Fatal(RedColor(fmt.Sprintf("[!!] Błąd konwersji '%v': %v", fh.Data[i], err)))
    }
    array[i] = num
  }

  return array
}

// TryParseByte próbuje przekonwertować dane z tekstu na bajty.
func (fh *InputFileHandler) TryParseByte() []byte {
  array := make([]byte, fh.DataLength)

  for i := 0; i < fh.DataLength; i++ {
    num, err := strconv.ParseUint(fh.Data[i], 0, 8)
    if err != nil {
      log.Fatal(RedColor(fmt.Sprintf("[!!] Błąd konwersji: %v", fh.Data[i], err)))
    }
    array[i] = byte(num)
  }

  return array
}
