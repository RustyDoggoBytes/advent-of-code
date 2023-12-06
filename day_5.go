package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type ProductionMap struct {
	DestinationStart, SourceStart, Range int

}

func createMap(rd *bufio.Reader) []ProductionMap {
	targetMap := make([]ProductionMap, 0)
	for {
		lineBytes, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}

		line := string(lineBytes)
		if line == "" {
			break
		}

		values := strings.Split(line, " ")
		destinationRangeStart, _ := strconv.Atoi(values[0])
		sourceRangeStart, _ := strconv.Atoi(values[1])
		mapRange, _ := strconv.Atoi(values[2])

		targetMap = append(targetMap, ProductionMap{DestinationStart: destinationRangeStart, SourceStart: sourceRangeStart, Range: mapRange})
	}

	return targetMap
}

func main() {
	filename := os.Args[1]

	dat, _ := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	rd := bufio.NewReader(dat)

	seeds := make([]int, 0)
	seedToSoilMap := make([]ProductionMap, 0)
	soilToFertilizerMap := make([]ProductionMap, 0)
	fertilizerToWaterMap := make([]ProductionMap, 0)
	waterToLightMap := make([]ProductionMap, 0)
	lightToTemperatureMap := make([]ProductionMap, 0)
	temperatureToHumidityMap := make([]ProductionMap, 0)
	humidityToLocationMap := make([]ProductionMap, 0)
	for {
		lineBytes, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}

		line := string(lineBytes)
		if line == "" {
			continue
		}

		if strings.Contains(line, "seeds") {
			seedsArray := strings.Split(line, " ")
			for _, seed := range seedsArray {
				if seedInt, err := strconv.Atoi(seed); err == nil {
					seeds = append(seeds, seedInt)
				}
			}

		} else if strings.Contains(line, "seed-to-soil map:") {
			seedToSoilMap = createMap(rd)
		} else if strings.Contains(line, "soil-to-fertilizer map:") {
			soilToFertilizerMap = createMap(rd)
		} else if strings.Contains(line, "fertilizer-to-water map:") {
			fertilizerToWaterMap = createMap(rd)
		} else if strings.Contains(line, "water-to-light map:") {
			waterToLightMap = createMap(rd)
		} else if strings.Contains(line, "light-to-temperature map:") {
			lightToTemperatureMap = createMap(rd)
		} else if strings.Contains(line, "temperature-to-humidity map:") {
			temperatureToHumidityMap = createMap(rd)
		} else if strings.Contains(line,"humidity-to-location map:") {
			humidityToLocationMap = createMap(rd)
		}
	}


	lowestLocation := math.MaxInt
	for _, seed := range seeds {

		soil :=  getNextPart(seedToSoilMap, seed)
		fertilizer :=  getNextPart(soilToFertilizerMap, soil)
		water :=  getNextPart(fertilizerToWaterMap, fertilizer)
		light :=  getNextPart(waterToLightMap, water)
		temp :=  getNextPart(lightToTemperatureMap, light)
		humidity :=  getNextPart(temperatureToHumidityMap, temp)
		location :=  getNextPart(humidityToLocationMap, humidity)

		if location < lowestLocation {
			lowestLocation = location
		}

		// fmt.Println("soil", soil, "fertilizer", fertilizer, "water", water, "light", light, "temp", temp, "humidity", humidity, "location", location)
	}
	fmt.Println(lowestLocation)
}

func getNextPart(fromToMap []ProductionMap, value int) int {

	for _, productionMap := range fromToMap {
		sourceStart := productionMap.SourceStart
		if value >= sourceStart && value <= (sourceStart + productionMap.Range) { 
			add := value - sourceStart 
			return productionMap.DestinationStart + add
		}

	}
	return value
}
