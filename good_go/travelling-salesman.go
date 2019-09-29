// Travelling Salesman Problem
//
// The travelling salesman problem (TSP) asks the following question:
// "Given a list of cities and the distances between each pair of
// cities, what is the shortest possible route that visits each
// city and returns to the origin city?"
//
// TSP can be modelled as an undirected weighted graph, such that
// cities are the graph's vertices, paths are the graph's edges,
// and a path's distance is the edge's weight. It is a minimization
// problem starting and finishing at a specified vertex after having
// visited each other vertex exactly once. Often, the model is a
// complete graph (i.e. each pair of vertices is connected by an
// edge). If no path exists between two cities, adding an arbitrarily
// long edge will complete the graph without affecting the optimal tour.

package main

import (
	"fmt"
	"github.com/raykov/tsp_investigation/utilities"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var cities []string
var weights [][]int

func main() {
	fileName := os.Args[1]

	size, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err.Error())
	}
	if size == 0 { size = 3 }

	runTimes, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err.Error())
	}

	runTimes = int(math.Max(float64(runTimes), 1))

	cities, weights = utilities.ReadDataFromCSV(fileName, size)

	var minimalPath int
	var route []string

	start := time.Now()
	for i := 0; i < runTimes; i ++ {
		minimalPath, route = findOptimalPath(0)
	}

	t := time.Now()
	elapsed := t.Sub(start)

	elapsed = elapsed / time.Duration(runTimes)

	fmt.Printf("minimalPath: %d \n", minimalPath)
	fmt.Printf("%v\n", strings.Join(route, " ► "))

	fmt.Printf("\n\n%v\n%v\n\n%v\n", start, t, elapsed)
}
