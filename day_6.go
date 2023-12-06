package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	Time, Distance int
}

func (r Race) calculateNumberOfWaysToBeatRecord() int {
	count := 0
	for i := 0; i <= r.Time; i++ {
		speed := i
		movingTime := r.Time - i
		distance := movingTime * speed
		if distance > r.Distance {
			count++
		}
	}

	return count
}

func main() {
	filename := os.Args[1]
	dat, _ := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	defer dat.Close()

	rd := bufio.NewReader(dat)

	timeBytes, _, _ := rd.ReadLine()
	distanceBytes, _, _ := rd.ReadLine()

	timeLine := string(timeBytes)
	distanceLine := string(distanceBytes)

	regex := regexp.MustCompile("[0-9]+")
	timesStr := regex.FindAllString(timeLine, -1)
	distanceStr := regex.FindAllString(distanceLine, -1)

	races := make([]Race, 0)
	for i := 0; i < len(timesStr); i++ {
		time, _ := strconv.Atoi(timesStr[i])
		distance, _ := strconv.Atoi(distanceStr[i])
		races = append(races, Race{Distance: distance, Time: time})
	}

	sum := 1
	kerningDistanceStr := ""
	kerningTimeStr := "" 
	for _, race := range races {
		recordBeat := race.calculateNumberOfWaysToBeatRecord()
		if recordBeat > 0 {
			sum *= recordBeat
		}
		kerningDistanceStr += strconv.Itoa(race.Distance)
		kerningTimeStr += strconv.Itoa(race.Time)
	}

	fmt.Println(sum)

	kerningDistance, _ := strconv.Atoi(kerningDistanceStr)
	kerningTime, _ := strconv.Atoi(kerningTimeStr)
	superRace := Race {Distance: kerningDistance, Time: kerningTime}  
	fmt.Println(superRace.calculateNumberOfWaysToBeatRecord())
}
