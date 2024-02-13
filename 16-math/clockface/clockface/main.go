package main

import (
	"os"
	"time"

	"learn-go-with-tests/16-math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
