package search

/*
	The State interface models a specific state in the state space of a problem.
	We need as bare minimum two functions:
		- Asserting if a State Equals() another
		- Obtaining the Neighbors() of this State
*/
type State interface {
	Equals(other State) bool
	Name() string
	Neighbors() []State
}

// HeuristicState composes a regular State, requiring additionally an inter-state Cost() function.
type WeightedState interface {
	Cost(target State) float64
	State
}

// HeuristicState composes WeightedState, requiring additionally a Heuristic() function.
type HeuristicState interface {
	Heuristic() float64
	WeightedState
}
