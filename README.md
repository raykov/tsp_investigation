# Travelling Salesman Problem investigation

The travelling salesman problem (TSP) asks the following question:
 > "Given a list of cities and the distances between each pair of cities, what is the shortest possible route that visits each city and returns to the origin city?"

TSP can be modelled as an undirected weighted graph, such that cities are the graph's vertices, paths are the graph's edges, and a path's distance is the edge's weight. It is a minimization problem starting and finishing at a specified vertex after having visited each other vertex exactly once. Often, the model is a complete graph (i.e. each pair of vertices is connected by an edge). If no path exists between two cities, adding an arbitrarily long edge will complete the graph without affecting the optimal tour.

#### Test data

 * `data.csv`
    Contains test data from the internet with known answer. Used to check validity of algorithm.  
 * `tsp_burgers.csv`
    Contains test data for algorithm demonstration.

---

### Ruby

`travelling_salesman.rb`

##### Run

```sh
ruby travelling_salesman.rb [amount_of_cities_to_test] [times_to_run]
```

 * `amount_of_cities_to_test` - integer from 3 to 20.
 * `times_to_run` - integer from 1 to any reasonable number (100_000, for example). How many times it will run code to get average execution time.

---

### Go. First blood (aka. Bad Go)

`bad_go/travelling_salesman.go`

##### Run

```sh
go run ./bad_go [path_to_data_file] [amount_of_cities_to_test] [times_to_run]
```

 * `path_to_data_file` - `./tsp_burgers.csv`
 * `amount_of_cities_to_test` - integer from 3 to 20.
 * `times_to_run` - integer from 1 to any reasonable number (100_000, for example). How many times it will run code to get average execution time.
