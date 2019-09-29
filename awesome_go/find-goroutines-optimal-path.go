package main

import (
	"math"
	"runtime"
	"sync"
)

type PathWithLength struct {
	Path []int
	Length int
}

var processesCount = runtime.NumCPU()

func findGoroutinesOptimalPath(startVertexIndex int) (minimalPath int, route []string) {
	var minLength = math.MaxInt64
	var minPath []int
	var waitGroup sync.WaitGroup

	toVisit := make([]int, len(cities)-1)[:0]
	for index, _ := range cities {
		if startVertexIndex == index {
			continue
		}

		toVisit = append(toVisit, index)
	}

	channel := make(chan *PathWithLength, 10)

	waitGroup.Add(1)
	go findGoroutinesAllPaths(channel, startVertexIndex, &[]int{}, &toVisit, 0, 1, &waitGroup)

	// watcher
	go func(waitGroup *sync.WaitGroup, channel chan *PathWithLength) {
		defer close(channel)
		waitGroup.Wait()
	}(&waitGroup, channel)

	for pathWithLength := range channel {
		if minLength > pathWithLength.Length {
			minPath = pathWithLength.Path
			minLength = pathWithLength.Length
		}
	}


	route = make([]string, len(cities))[:0]
	for _, val := range minPath {
		route = append(route, cities[val])
	}

	return minLength, route
}

func findGoroutinesAllPaths(channel chan *PathWithLength, vertexIndex int, path *[]int, toVisit *[]int, length int, deepness int, waitGroup *sync.WaitGroup) {
	if deepness == 1 {
		defer waitGroup.Done();
	}

	currentPath := append([]int(nil), *path...)
	currentPath = append(currentPath, vertexIndex)

	currentToVisit := citiesToVisit(vertexIndex, toVisit)

	for _, cityIndex := range currentToVisit {
		if runtime.NumGoroutine() >= processesCount-4 {
			waitGroup.Add(1)
			go func(){
				defer waitGroup.Done()
				findGoroutinesAllPaths(channel, cityIndex, &currentPath, &currentToVisit, length+weights[vertexIndex][cityIndex], deepness + 1, waitGroup)
			}()
		} else {
			findGoroutinesAllPaths(channel, cityIndex, &currentPath, &currentToVisit, length+weights[vertexIndex][cityIndex], deepness + 1, waitGroup)
		}

	}

	if len(currentToVisit) == 0 {
		length += weights[currentPath[len(currentPath)-1]][currentPath[0]]
		channel <- &PathWithLength{Path: currentPath,  Length: length}
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
