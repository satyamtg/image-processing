package pgmutil

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filepath string) []string {
	data, err := ioutil.ReadFile(filepath)
	check(err)
	var dataString = string(data)
	var dataArray []string
	var lines = strings.Split(dataString, "\n")
	for i := 0; i < len(lines); i++ {
		var tempArray = strings.Split(lines[i], " ")
		for j := 0; j < len(tempArray); j++ {
			if tempArray[j] != "" {
				dataArray = append(dataArray, tempArray[j])
			}
		}
	}
	return(dataArray)
}

func readPlainPgm(fileTokens []string) (int, int, int, [][]int){
	width, err := strconv.Atoi(fileTokens[0])
	check(err)
	height, err := strconv.Atoi(fileTokens[1])
	check(err)
	maxVal, err := strconv.Atoi(fileTokens[2])
	check(err)
	pix := make([][]int, height)
	for i, k := 0, 3; i<height; i++ {
		pix[i] = make([]int, width)
		for j := 0; j<width; j,k = j+1, k+1 {
			data, err := strconv.Atoi(fileTokens[k])
			check(err)
			pix[i][j] = data
		}
	}
	return width, height, maxVal, pix
}

func readExPgm(fileTokens []string) (int, int, int, [][]int){
	fmt.Println("Will be implemented later")
	width, err := strconv.Atoi(fileTokens[0])
	check(err)
	height, err := strconv.Atoi(fileTokens[1])
	check(err)
	maxVal, err := strconv.Atoi(fileTokens[2])
	check(err)
	var pix [][]int
	return width, height, maxVal, pix
}

func ReadPgm(filepath string) (int, int, int, [][]int){
	fileTokensArray := readFile(filepath)
	if fileTokensArray[0] == "P2" {
		fileTokensArray = fileTokensArray[1:]
		return readPlainPgm(fileTokensArray)
	} else if fileTokensArray[0] == "P5" {
		fileTokensArray = fileTokensArray[1:]
		return readExPgm(fileTokensArray)
	} else {
		panic("Mimetype not present. File may be corrupt")
	}
}