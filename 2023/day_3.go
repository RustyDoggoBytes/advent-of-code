package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func main() {
	filename := os.Args[1]
	dat, _ := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	defer dat.Close()

	rd := bufio.NewReader(dat)

	matrix := make([][]rune, 0)
	for {
		lineBytes, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(lineBytes)

		columns := make([]rune, 0)
		for _, char := range line {
			columns = append(columns, char)
		}
		matrix = append(matrix, columns)
	}

	checkEngineParts(matrix)
	checkGearRatio(matrix)
}

func checkGearRatio(matrix [][]rune) {
	sum := 0

	for rowIdx, column := range matrix {
		for columnIdx, symbol := range column {
			digitSet := make(map[string]bool)

			if symbol != '*' {
				continue
			}

			for i := rowIdx - 1; i <= (rowIdx + 1); i++ {
				if i < 0 || i >= len(matrix) {
					continue
				}

				row := matrix[i]
				for j := columnIdx - 1; j <= (columnIdx + 1); j++ {
					if j < 0 || j >= len(row) {
						continue
					}

					surroundingChar := matrix[i][j]
					if unicode.IsNumber(surroundingChar) {
						digit := string(surroundingChar)

						k := j + 1
						for {
							if k >= len(row) || !unicode.IsNumber(matrix[i][k]) {
								break
							}
							fmt.Printf("digit %s, %c row=%d col=%d\n", digit, matrix[i][k], i, k)
							digit += string(matrix[i][k])
							k++
						}

						k = j - 1
						for {
							if k < 0 || !unicode.IsNumber(matrix[i][k]) {
								break
							}
							fmt.Printf("digit %s, %c row=%d col=%d\n", digit, matrix[i][k], i, k)
							digit = string(matrix[i][k]) + digit
							k--
						}
						digitSet[digit] = true
						fmt.Println(digit)
						// digitInt, _ := strconv.Atoi(digit)
					}
				}
			}
			digitCount := len(digitSet)

			if digitCount == 2 {
				setSum := 1
				for digit, _ := range digitSet {
					digitInt, _ := strconv.Atoi(digit)
					setSum *= digitInt
				}
				sum += setSum
			}
		}
	}
	fmt.Println("gear ratio sum", sum)
}

func checkEngineParts(matrix [][]rune) {
	sum := 0

	for rowIdx, column := range matrix {
		digitStr := ""
		startIdx := -1

		for columnIdx, char := range column {
			if unicode.IsNumber(char) {
				if len(digitStr) == 0 {
					startIdx = columnIdx
				}
				digitStr += string(char)

				if len(column)-1 != columnIdx {
					continue
				}
			}

			if startIdx != -1 {
				isAdjacentToSymbol := checkSurroundings(matrix, startIdx, columnIdx-1, rowIdx)
				if isAdjacentToSymbol {
					digit, _ := strconv.Atoi(digitStr)
					sum += digit
				}

				digitStr = ""
				startIdx = -1
			}
		}
	}
	fmt.Println("engine part sum", sum)
}

func checkSurroundings(matrix [][]rune, startIdx int, endIdx int, rowIdx int) bool {
	for i := rowIdx - 1; i <= (rowIdx + 1); i++ {
		if i < 0 || i >= len(matrix) {
			continue
		}

		row := matrix[i]
		for j := startIdx - 1; j <= (endIdx + 1); j++ {
			if j < 0 || j >= len(row) {
				continue
			}

			char := matrix[i][j]
			isSymbol := char != '.' && (unicode.IsSymbol(char) || unicode.IsPunct(char))
			// fmt.Printf("char=%c row=%d col=%d symbol=%t\n", char, i, j, isSymbol)
			if isSymbol {
				return true
			}
		}
	}
	return false
}
