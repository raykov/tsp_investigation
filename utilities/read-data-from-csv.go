package utilities

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadDataFromCSV(fileName string, takeOnly int) (cities []string, weights [][]int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if cities, err = header(scanner); err != nil {
		log.Fatal(err)
	}

	switch {
	case takeOnly < 3:
		takeOnly = 3
	case takeOnly > len(cities):
		takeOnly = len(cities)
	}

	cities = cities[:takeOnly]

	weights = make([][]int, len(cities))[:0]

	for i := 0; i < len(cities); i++ {
		scanner.Scan()
		weights = append(weights, toInts(strings.Split(scanner.Text(), ",")[:takeOnly]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

type fileScanner interface {
	Scan() bool
	Text() string
}

func header(scanner fileScanner) (cities []string , err error) {
	if !scanner.Scan() {
		err = errors.New("couldn't read a header line")
		return
	}

	cities = strings.Split(scanner.Text(), ",")

	return
}

func toInts(arr []string) []int {
	newArr := make([]int, len(arr))[:0]

	for _, val := range arr {
		floatVal, err := strconv.ParseFloat(strings.TrimSpace(val), 32)
		newVal := int(floatVal * 1000) // convert kms to meters
		if err != nil {
			log.Fatal(err)
		}

		newArr = append(newArr, newVal)
	}

	return newArr
}
