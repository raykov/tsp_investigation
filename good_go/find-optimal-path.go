package main

import (
	"fmt"
	"math"
)

var minLength = math.MaxInt64
var minPath []int

func findOptimalPath(startVertexIndex int) (minimalPath int, route []string) {
	if startVertexIndex > len(cities) - 1 {
		panic(fmt.Sprintf("startVertexIndex %v is out of cities count %v", startVertexIndex, len(cities)))
	}

	toVisit := make([]int, len(cities)-1)[:0]
	for index, _ := range cities {
		if startVertexIndex == index {
			continue
		}

		toVisit = append(toVisit, index)
	}

	findAllPaths(startVertexIndex, &[]int{}, &toVisit, 0)

	route = make([]string, len(cities))[:0]
	for _, val := range minPath {
		route = append(route, cities[val])
	}

	return minLength, route
}

func findAllPaths(vertexIndex int, path *[]int, toVisit *[]int, length int) {
	currentPath := append([]int(nil), *path...)
	currentPath = append(currentPath, vertexIndex)

	currentToVisit := citiesToVisit(vertexIndex, toVisit)

	for _, cityIndex := range currentToVisit {
		findAllPaths(cityIndex, &currentPath, &currentToVisit, length+weights[vertexIndex][cityIndex])
	}

	if len(currentToVisit) == 0 {
		length += weights[currentPath[len(currentPath)-1]][currentPath[0]]

		if minLength > length {
			minPath = currentPath
			minLength = length
		}
	}
}

func citiesToVisit(vertex int, oldList *[]int) (newList []int) {
	newList = make([]int, len(*oldList)-1)[:0]

	for _, cityNum := range *oldList {
		if cityNum == vertex {
			continue
		}

		newList = append(newList, cityNum)
	}

	return
}
