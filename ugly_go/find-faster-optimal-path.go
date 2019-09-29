package main

import (
	"math"
	"sync"
)

var minLength = math.MaxInt64
var minPath []int

func findFasterOptimalPath(startVertexIndex int) (minimalPath int, route []string) {
	var waitGroup sync.WaitGroup

	toVisit := make([]int, len(cities)-1)[:0]
	for index, _ := range cities {
		if startVertexIndex == index {
			continue
		}

		toVisit = append(toVisit, index)
	}
	findFasterAllPaths(&waitGroup, startVertexIndex, &[]int{}, &toVisit, 0, 1)

	waitGroup.Wait()

	route = make([]string, len(cities))[:0]
	for _, val := range minPath {
		route = append(route, cities[val])
	}

	return minLength, route
}

func findFasterAllPaths(waitGroup *sync.WaitGroup, vertexIndex int, path *[]int, toVisit *[]int, length int, deepness int) {
	currentPath := append([]int(nil), *path...)
	currentPath = append(currentPath, vertexIndex)

	currentToVisit := citiesToVisit(vertexIndex, toVisit)

	for _, cityIndex := range currentToVisit {
		if deepness == 1 {
			waitGroup.Add(1)
			go func(){
				defer waitGroup.Done()

				findFasterAllPaths(waitGroup, cityIndex, &currentPath, &currentToVisit, length+weights[vertexIndex][cityIndex], deepness + 1)
			}()
		} else {
			findFasterAllPaths(waitGroup, cityIndex, &currentPath, &currentToVisit, length+weights[vertexIndex][cityIndex], deepness + 1)
		}
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
