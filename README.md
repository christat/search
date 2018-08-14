[![GoDoc](https://godoc.org/github.com/christat/search?status.svg)](https://godoc.org/github.com/christat/search)
[![Build Status](https://travis-ci.org/christat/search.svg?branch=master)](https://travis-ci.org/christat/search)
# Search (Algorithms) - a minimal library

A minimal search algorithms library, aimed (mainly) at minimization-problem solving and educational purposes.

The package is subdivided with two main sub-packages:

- **blind** algorithms
- **informed** algorithms

Every algorithm comes in two variants: regular and `Benchmark`.
The latter returns metadata in the form of execution time and number 
of node expanions made within an `AlgorithmBenchmark` struct.

Some algorithms offer the possibility to choose the underlying data 
structure implementation by setting passing a flag to the call.
See an example [here](https://github.com/christat/search/blob/master/blind/breadth_first.go#L11).

## Download/Installation

In your Go project's root directory, open a terminal and paste the following:

```
go get github.com/christat/search
```

## Blind Algorithms

- `BreadthFirst`
- `DepthFirst`
- `DepthFirstBranchAndBound`
- `Djikstra`
- `IterativeDeepening`

## Informed Algorithms

- `AStar`
- `Beam`
- `GreedyBestFirst`
- `HillClimbing`
- `IterativeDeepeningAStar`



## License

Licensed under the MIT license.
