package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Card struct {
	Score, CopyCount, MatchingCount int
}

func main() {
	filename := os.Args[1]
	dat, _ := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	defer dat.Close()

	rd := bufio.NewReader(dat)

	cards := make([]Card, 0)
	id := 1
	for {
		lineBytes, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(lineBytes)
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := strings.Split(numbers[0], " ")
		haveNumbers := strings.Split(numbers[1], " ")

		winningSet := make(map[string]bool)
		for _, number := range winningNumbers {
			winningSet[number] = true
		}

		score := 0
		matchingCount := 0
		for _, number := range haveNumbers {
			if number == "" {
				continue
			}
			if _, ok := winningSet[number]; ok {
				matchingCount++
				if score == 0 {
					score++
				} else {
					score += score
				}
			}
		}

		cards = append(cards, Card{CopyCount: 1, Score: score, MatchingCount: matchingCount})
		id++
	}

	sum := 0
	for i, card := range cards {
		sum += card.Score

		for j := 0; j < card.CopyCount; j++ {
			for k := 1; k <= card.MatchingCount; k++ {
				cards[i+k].CopyCount++
			}
		}
	}
	fmt.Println("winning sums", sum)

	cardsSum := 0
	for _, card := range cards {
		cardsSum += card.CopyCount
	}
	fmt.Println("cards sums", cardsSum)

}
