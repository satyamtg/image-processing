package main

import (
	"fmt"

	"github.com/satyamtg/image-processing/pgmutil"
)

func main() {
	var pgmPath = "pgm_files/balloons_noisy.ascii.pgm"
	width, height, maxVal, pix:= pgmutil.ReadPgm(pgmPath)
	fmt.Println(width)
	fmt.Println(height)
	fmt.Println(maxVal)
	fmt.Println(pix)
	fmt.Println("Done")
}
