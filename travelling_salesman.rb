# Travelling Salesman Problem
#
# The travelling salesman problem (TSP) asks the following question:
# "Given a list of cities and the distances between each pair of
# cities, what is the shortest possible route that visits each
# city and returns to the origin city?"
#
# TSP can be modelled as an undirected weighted graph, such that
# cities are the graph's vertices, paths are the graph's edges,
# and a path's distance is the edge's weight. It is a minimization
# problem starting and finishing at a specified vertex after having
# visited each other vertex exactly once. Often, the model is a
# complete graph (i.e. each pair of vertices is connected by an
# edge). If no path exists between two cities, adding an arbitrarily
# long edge will complete the graph without affecting the optimal tour.

class Graph
  attr_reader :vertices, :edges, :min_length, :min_path

  def initialize
    @edges = Hash.new

    @min_path = []
    @min_length = Float::INFINITY
  end

  def add_edge(u, v, weight)
    @edges[u] ||= {}
    @edges[u][v] = weight
  end

  def reset
    @min_path = []
    @min_length = Float::INFINITY
  end

  def find_minimal_path(start_index = 0)
    start_vertex = edges.keys[start_index]
    find_all_paths(start_vertex, [], 0)

    print "minimalPath: #{min_length}\n#{min_path.join(" ► ")}\n"
  end

  def find_all_paths(start_vertex, path = [], len)
    current_path = path.dup

    current_path << start_vertex

    unvisited_neighbors = edges[start_vertex].keys - current_path

    unvisited_neighbors.each do |neighbor|
      find_all_paths(neighbor, current_path, len + edges[start_vertex][neighbor])
    end

    if unvisited_neighbors.size.zero?
      len += edges[start_vertex][current_path.first]

      if min_length > len
        @min_length = len
        @min_path = current_path
      end
    end
  end

end

def read_data_from(file_name, loadOnly = 0)
  lines = File.read(file_name).split("\n")
  headers = lines.shift

  cities = headers.split(",").map(&:strip)

  loadOnly = [[loadOnly.to_i, 3].max, cities.size].min

  cities = cities[0, loadOnly]
  weights = []

  lines.each.with_index do |line, index|
    break if index > loadOnly

    weights << line.split(",")[0, loadOnly].map(&:strip).map(&:to_f).map { |w| (w * 1000).to_i }
  end

  [cities, weights]
end

def find_unit(elapsed)
  return [elapsed, :seconds] if elapsed > 1

  new_elapsed = elapsed.dup
  %i(milliseconds microseconds nanoseconds picoseconds).each do |unit|
    new_elapsed *= 1000
    return [new_elapsed, unit] if new_elapsed > 1
  end

  [elapsed, :unknown]
end

# ----------------------------

graph_from_file = Graph.new

size = ARGV[0].to_i
size = 3 if size.zero?

run_times = [ARGV[1].to_i, 1].max

cities_from_file, weights_from_file = read_data_from("./tsp_burgers.csv", size)

weights_from_file.each_with_index do |row, row_index|
  row.each_with_index do |weight, column_index|
    next if row_index == column_index

    graph_from_file.add_edge(cities_from_file[row_index], cities_from_file[column_index], weight)
  end
end

start_time = Time.now
run_times.times do
  graph_from_file.find_minimal_path(0)
  graph_from_file.reset
end
end_time = Time.now

print("\n\n#{start_time}\n")
print("#{end_time}\n")

elapsed, unit = find_unit((end_time - start_time) / run_times)

print("\n#{elapsed} #{unit}\n")

