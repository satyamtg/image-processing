package pgmutil

import (
	"bufio"
	"os"
	"strconv"
)

func WritePlainPgm(filepath string, pix [][]int, maxval int) {
	width := len(pix[0])
	height := len(pix)
	maxLengthOfLineInFile := 70
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	noBytes, err := w.WriteString("P2\n")
	noBytes, err = w.WriteString("#File created by satyamtg\n")
	noBytes, err = w.WriteString(strconv.Itoa(width) + " " + strconv.Itoa(height) + "\n")
	noBytes, err = w.WriteString(strconv.Itoa(maxval) + "\n")
	currLineLength := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if currLineLength+len(strconv.Itoa(pix[i][j])) < maxLengthOfLineInFile-1 {
				noBytes, err = w.WriteString(strconv.Itoa(pix[i][j]) + " ")
				currLineLength = currLineLength + noBytes
			} else {
				noBytes, err = w.WriteString("\n")
				noBytes, err = w.WriteString(strconv.Itoa(pix[i][j]) + " ")
				currLineLength = noBytes
			}
		}
	}
	w.Flush()
}
