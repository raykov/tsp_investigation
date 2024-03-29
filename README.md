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

 * `travelling_salesman.rb`

##### Run

```sh
ruby travelling_salesman.rb [amount_of_cities_to_test] [times_to_run]
```

 * `amount_of_cities_to_test` - integer from 3 to 20.
 * `times_to_run` - integer from 1 to any reasonable number (100_000, for example). How many times it will run code to get average execution time.

---

### Go. First blood (aka. Bad Go)

 * `bad_go/travelling_salesman.go`

##### Run

```sh
go run ./bad_go [path_to_data_file] [amount_of_cities_to_test] [times_to_run]
```

 * `path_to_data_file` - `./tsp_burgers.csv`
 * `amount_of_cities_to_test` - integer from 3 to 20.
 * `times_to_run` - integer from 1 to any reasonable number (100_000, for example). How many times it will run code to get average execution time.

---

### Good Go

 * `good_go/travelling_salesman.go` - runner
 * `good_go/find-optimal-path.go` - algorithm

##### Run

```sh
go run ./good_go [path_to_data_file] [amount_of_cities_to_test] [times_to_run]
```

 * `path_to_data_file` - `./tsp_burgers.csv`
 * `amount_of_cities_to_test` - integer from 3 to 20.
 * `times_to_run` - integer from 1 to any reasonable number (100_000, for example). How many times it will run code to get average execution time.

---

### Ugly Go

 * `ugly_go/travelling_salesman.go` - runner
 * `ugly_go/find-faster-optimal-path.go` - algorithm

##### Run

```sh
go run ./ugly_go [path_to_data_file] [amount_of_cities_to_test] [times_to_run]
```

 * `path_to_data_file` - `./tsp_burgers.csv`
 * `amount_of_cities_to_test` - integer from 3 to 20.
 * `times_to_run` - integer from 1 to any reasonable number (100_000, for example). How many times it will run code to get average execution time.

---

### Awesome Go

 * `awesome_go/travelling_salesman.go` - runner
 * `awesome_go/find-goroutines-optimal-path.go` - algorithm

##### Run

```sh
go run ./awesome_go [path_to_data_file] [amount_of_cities_to_test] [times_to_run]
```

 * `path_to_data_file` - `./tsp_burgers.csv`
 * `amount_of_cities_to_test` - integer from 3 to 20.
 * `times_to_run` - integer from 1 to any reasonable number (100_000, for example). How many times it will run code to get average execution time.

---

### Results

```shell script
> go run ./good_go ./tsp_burgers.csv 11 1
minimalPath: 7550
Office ► EastBurger ► Markthalleneun ► Windburger ► Angry Chicken ► Captain burger 36 ► Burgermeister ► Pacifico Berlin ► Rosengarten am Engelbecken ► BRGRS BRGRS - Organic Burgers ► Wurstpate


2019-09-29 14:43:50.21425 +0200 CEST m=+0.000203276
2019-09-29 14:43:50.986152 +0200 CEST m=+0.772101059

771.897783ms
```

### Benchmarks

You can find benchmark results for different data sets (from 3 cities up to 13. 14 only for Good Go implementation) inside [`/benchmarks`](./benchmarks) folder.

#### Results for 3 cities

![](./benchmarks/3_cities_minimum.png)

#### Results for 6 cities

![](./benchmarks/6_cities.png)

#### Results for 11 cities

![](./benchmarks/11_cities.png)

#### Results for 13 cities

![](./benchmarks/13_cities_maximum.png)
