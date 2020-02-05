package main

import (
	"fmt"

	"github.com/satyamtg/image-processing/imgalgo"
	"github.com/satyamtg/image-processing/pgmutil"
)

func main() {
	var pgmPath = "pgm_files/venus1.ascii.pgm"
	var newFileName = "satyamtg.pgm"
	imgalgo.BinarizeGlobalOtsu(pgmPath, newFileName)
	width, height, maxVal, pix := pgmutil.ReadPgm(pgmPath)
	fmt.Println(width)
	fmt.Println(height)
	fmt.Println(maxVal)
	fmt.Println(pix)
	//pgmutil.WritePlainPgm(newFileName, pix, 255)
	fmt.Println("Done")
}
