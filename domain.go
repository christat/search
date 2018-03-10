package search

/*
	Domain represents a search problem definition.
	The type implementing this interface must provide a way to obtain both the cost (G) and heuristic (H) function, as well
	a a neighborhood access function. In instances where a weight/heuristic is not needed, the functions can be simply left
	constant (as they won't be used anyways).
*/
type Domain interface {
	Neighbors(node interface{}) ([]interface{}, error) // decision point #1: adjacency list vs adjacency matrix
	G(node interface{}) (float64, error)
	H(node interface{}) (float64, error)
}