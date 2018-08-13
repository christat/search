package search

import (
	"time"
	"fmt"
)

// Convenience wrapper to obtain extra information regarding the search performed with Benchmark_<Algorithm> variants.
type AlgorithmBenchmark struct {
	ElapsedTime time.Duration
	TotalExpansions uint
}

func (ab AlgorithmBenchmark) String() string {
	return fmt.Sprintf("Execution took %v ns and expanded %v nodes.", ab.ElapsedTime.Nanoseconds(), ab.TotalExpansions)
}

func (ab AlgorithmBenchmark) GetExpandedNodesPerSecond() float64 {
	return float64(ab.TotalExpansions) / ab.ElapsedTime.Seconds()
}
