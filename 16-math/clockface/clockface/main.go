// Writes an SVG clockface of the current time to Stdout.
package main

import (
	"learn-go-with-tests/16-math/clockface/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
