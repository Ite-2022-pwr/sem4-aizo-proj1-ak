// Package utils zawiera funkcje ogólnego użytku
package utils

import (
	"fmt"
	"log"
	"time"
)

// TimeElapsed wypisuje ile czasu minęło od danego momentu
func PrintTimeElapsed(start time.Time, prompt string) {
  elapsed := time.Since(start)
  elapsedStr := GreenColor(fmt.Sprintf("%v", elapsed))
  log.Printf("[+] %s zajęło %s\n", prompt, elapsedStr)
}
