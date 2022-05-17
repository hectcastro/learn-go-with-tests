package main

import (
	"os"
	"time"

	"github.com/hectcastro/learn-go-with-tests/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
