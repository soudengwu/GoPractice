package main

import (
	"os"
	"time"

	clockFace "github.com/soudengwu/GoPractice/015_maths"
)

func main() {
	t := time.Now()
	clockFace.SVGWriter(os.Stdout, t)
}
