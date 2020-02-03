package main

import (
	"github.com/satyamtg/image-processing/pgmutil"
)

func FindOtsuThreshold(width int, height int, maxVal int, pix [][]int) int {
	histogramCounts := make([]int, maxVal+1)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			histogramCounts[pix[i][j]] += 1
		}
	}
	totalPixels := width * height
	sum := 0
	for i := 0; i <= maxVal; i++ {
		sum = sum + i*histogramCounts[i]
	}
	sumB := 0
	wB := 0
	wF := 0
	var varMax float32
	threshold := 0
	for t := 0; t <= maxVal; t++ {
		wB += histogramCounts[t]
		if wB == 0 {
			continue
		}
		wF = totalPixels - wB
		if wF == 0 {
			break
		}
		sumB += t * histogramCounts[t]
		mB := float32(sumB / wB)
		mF := float32((sum - sumB) / wF)
		varBetween := float32(wB) * float32(wF) * (mB - mF) * (mB - mF)
		if varBetween > float32(varMax) {
			varMax = varBetween
			threshold = t
		}

	}
	return threshold
}

func main() {
	var pgmPath = "pgm_files/pbmlib.ascii.pgm"
	var newFileName = "satyamtg.pgm"
	width, height, maxVal, pix := pgmutil.ReadPgm(pgmPath)
	threshold := FindOtsuThreshold(width, height, maxVal, pix)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if pix[i][j] >= threshold {
				pix[i][j] = maxVal
			} else {
				pix[i][j] = 0
			}
		}
	}
	pgmutil.WritePlainPgm(newFileName, pix, maxVal)
}
