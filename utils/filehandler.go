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

type FileHandler struct {
  FilePath string
  Data []string
}

func NewFileHandler(filepath string) *FileHandler {
  fh := &FileHandler{FilePath: filepath}
  fh.LoadData()
  return fh
}

func (fh *FileHandler) LoadData() {
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

  for i := 0; i < arraySize; i++ {
    line, err = rdr.ReadString('\n')
    if err != nil {
      log.Fatal(RedColor(fmt.Sprintf("[!!] Błąd czytania linii '%v': %v", line, err)))
    }
    fh.Data[i] = strings.TrimSpace(line)
  }
}
