package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var RED_LIMIT, GREEN_LIMIT, BLUE_LIMIT = 12, 13, 14

func main() {
	filename := os.Args[1]
	dat, _ := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	defer dat.Close()

	rd := bufio.NewReader(dat)
	gameRegex := regexp.MustCompile("Game ([0-9]+):(.+)")
	roundRegex := regexp.MustCompile("([0-9]+) (blue|red|green)")

	sum := 0
	sum_power := 0
	for {
		lineBytes, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			panic(1)
		}
		line := string(lineBytes)

		result := gameRegex.FindStringSubmatch(line)
		rounds := strings.Split(result[2], ";")
		id, _ := strconv.Atoi(result[1])
		min_red, min_green, min_blue := 0, 0, 0
		isGameValid := true
		for _, round := range rounds {
			roundGroup := roundRegex.FindAllStringSubmatch(round, -1)
			for _, group := range roundGroup {
				count_ball, _ := strconv.Atoi(group[1])
				limit := -1

				switch color := group[2]; color {
				case "red":
					limit = RED_LIMIT
					if count_ball > min_red {
						min_red = count_ball
					}
				case "green":
					limit = GREEN_LIMIT
					if count_ball > min_green {
						min_green = count_ball
					}
				case "blue":
					limit = BLUE_LIMIT
					if count_ball > min_blue {
						min_blue = count_ball
					}
				default:
					fmt.Println(color)
					panic(1)
				}

				if count_ball > limit {
					isGameValid = false
				}

			}

		}
		fmt.Printf("line: %s -> %s\n", line, isGameValid)
		if isGameValid {
			sum += id
		}
		sum_power += min_red * min_green * min_blue
	}
	fmt.Printf("sum=%d power_sum=%d\n", sum, sum_power)
}
