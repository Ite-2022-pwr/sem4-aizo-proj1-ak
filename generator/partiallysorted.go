package generator

import (
	"math/rand"
	"time"
)

func GeneratePartiallySortedArrayInt(n, sortedPercentage int) []int {
  array := make([]int, n)

  sortedNum := n * sortedPercentage / 100

  // log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v%% tablicy typu %v o rozmiarze %v", utils.YellowColor(sortedPercentage), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))

  sortedPart := GenerateAscendingSortedArrayInt(sortedNum)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Int() % 100
    if r < array[sortedNum - 1] {
      r += array[sortedNum - 1]
    }
    array[i] = r
  }

  return array
}


func GeneratePartiallySortedArrayInt32(n, sortedPercentage int) []int32 {
  array := make([]int32, n)

  sortedNum := n * sortedPercentage / 100

  // log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v%% tablicy typu %v o rozmiarze %v", utils.YellowColor(sortedPercentage), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))

  sortedPart := GenerateAscendingSortedArrayInt32(sortedNum)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Int31() % 100
    if r < array[sortedNum - 1] {
      r += array[sortedNum - 1]
    }
    array[i] = r
  }

  return array
}


func GeneratePartiallySortedArrayInt64(n, sortedPercentage int) []int64 {
  array := make([]int64, n)

  sortedNum := n * sortedPercentage / 100

  // log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v%% tablicy typu %v o rozmiarze %v", utils.YellowColor(sortedPercentage), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))

  sortedPart := GenerateAscendingSortedArrayInt64(sortedNum)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Int63() % 1000
    if r < array[sortedNum - 1] {
      r += array[sortedNum - 1]
    }
    array[i] = r
  }

  return array
}


func GeneratePartiallySortedArrayFloat32(n, sortedPercentage int) []float32 {
  array := make([]float32, n)

  sortedNum := n * sortedPercentage / 100

  // log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v%% tablicy typu %v o rozmiarze %v", utils.YellowColor(sortedPercentage), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))

  sortedPart := GenerateAscendingSortedArrayFloat32(sortedNum)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Float32()
    if r < array[sortedNum - 1] {
      r += array[sortedNum - 1]
    }
    array[i] = r
  }

  return array
}


func GeneratePartiallySortedArrayFloat64(n, sortedPercentage int) []float64 {
  array := make([]float64, n)

  sortedNum := n * sortedPercentage / 100

  // log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v%% tablicy typu %v o rozmiarze %v", utils.YellowColor(sortedPercentage), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))

  sortedPart := GenerateAscendingSortedArrayFloat64(sortedNum)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Float64()
    if r < array[sortedNum - 1] {
      r += array[sortedNum - 1]
    }
    array[i] = r
  }

  return array
}


func GeneratePartiallySortedArrayByte(n, sortedPercentage int) []byte {
  array := make([]byte, n)

  sortedNum := n * sortedPercentage / 100

  // log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v%% tablicy typu %v o rozmiarze %v", utils.YellowColor(sortedPercentage), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))

  sortedPart := GenerateAscendingSortedArrayByte(sortedNum)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Int()
    if byte(r) < array[sortedNum - 1] {
      r = min(int(array[sortedNum - 1]) + r, 255)
    }
    array[i] = byte(r)
  }

  return array
}


func GeneratePartiallySortedArrayString(n, sortedPercentage int) []string {
  array := make([]string, n)

  sortedNum := n * sortedPercentage / 100

  // log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v%% tablicy typu %v o rozmiarze %v", utils.YellowColor(sortedPercentage), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))

  sortedPart := GenerateAscendingSortedArrayString(sortedNum)
  copy(array[:sortedNum], sortedPart)

  for i := sortedNum; i < n; i++ {
    r := GenerateRandomString(5)
    if r < array[sortedNum - 1] {
      r = array[sortedNum - 1] + r
    }
    array[i] = r
  }

  return array
}
