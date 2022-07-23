package main

import (
	"clockface/logic"
	"os"
	"time"
)

func main() {
	t := time.Now()
	logic.SVGWriter(os.Stdout, t)
}
