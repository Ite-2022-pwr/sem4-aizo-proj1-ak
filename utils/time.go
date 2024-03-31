// Package utils zawiera funkcje ogólnego użytku
package utils

import (
	"log"
	"time"
)

// TimeElapsed wypisuje ile czasu minęło od danego momentu
func PrintTimeElapsed(start time.Time, prompt string) {
  elapsed := time.Since(start)
  log.Printf("[+] %s zajęło %s\n", prompt, GreenColor(elapsed))
}
