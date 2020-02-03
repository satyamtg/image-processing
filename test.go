package main

import (
	"fmt"

	"github.com/satyamtg/image-processing/pgmutil"
)

func main() {
	var pgmPath = "pgm_files/balloons_noisy.ascii.pgm"
	var newFileName = "satyamtg.pgm"
	width, height, maxVal, pix := pgmutil.ReadPgm(pgmPath)
	fmt.Println(width)
	fmt.Println(height)
	fmt.Println(maxVal)
	fmt.Println(pix)
	pgmutil.WritePlainPgm(newFileName, pix, 255)
	fmt.Println("Done")
}
