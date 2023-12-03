package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
) 

var numberMap = map[string]string {
	"one": "1",
	"two":"2",
	"three":"3",
	"four":"4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
}

func getDigit(number string) string {
	if _, err := strconv.Atoi(number); err == nil {
		return number
	} else if val, ok := numberMap[number]; ok {
		return val
	}

	return ""
}

func main() {
	dat, _ := os.OpenFile("day_1_input.txt", os.O_RDONLY, os.ModePerm)
	defer dat.Close()
	reader := bufio.NewReader(dat)
	sum := 0
	re := regexp.MustCompile("[1-9]|one|two|three|four|five|six|seven|eight|nine")
	for {
		lineByte, _, err := reader.ReadLine() 
		line := string(lineByte)
		if err != nil {
			if err == io.EOF {
				break 
			}
			log.Fatalln("unexpected error", err)
			return
		}
		first := "" 
		last := "" 

		for idx, _ := range line {
			digit := re.FindString(line[:idx+1])
			if digit != "" {
				first = getDigit(digit)
				break;
			}
		}

		for idx, _ := range line {
			digit := re.FindString(line[len(line)-idx - 1:])
			if digit != "" {
				last = getDigit(digit)
				break;
			}
		}
		fmt.Println(line + "->" + "(" + first + ")"+ "(" + last + ")")
		twoDigit, _ := strconv.Atoi(first + last)
		if len(first+last) != 2 {
			panic(1)
		}
		sum = sum + twoDigit
	}
	fmt.Println(sum)
}
