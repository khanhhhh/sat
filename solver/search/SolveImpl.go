package search

import "github.com/khanhhhh/sat/instance"

// Guesser :
type Guesser func(ins instance.Instance) (converged bool, variableOut instance.Variable, vaueOut bool)

// Searcher :
type Searcher func(ins instance.Instance) (sat bool, assignment map[instance.Variable]bool)

// Solve :
func Solve(ins instance.Instance, guesserIn Guesser, completeSearch bool, completeSearcher Searcher) (sat bool, assignment map[instance.Variable]bool) {
	// terminating condition
	if len(ins.ClauseMap()) == 0 {
		return true, make(map[instance.Variable]bool)
	}
	for _, variableMap := range ins.ClauseMap() {
		if len(variableMap) == 0 {
			return false, nil
		}
	}
	// guess and branch out
	converged, variable, value := guesserIn(ins)
	// unconverged path
	if converged == false {
		sat, assignment = completeSearcher(ins)
		return sat, assignment
	}
	// accept path
	{
		ins := ins.Clone()
		ins.Reduce(variable, value)
		sat, assignment = Solve(ins, guesserIn, completeSearch, completeSearcher)
		if sat {
			assignment[variable] = value
			return sat, assignment
		}
	}
	// reject path
	if completeSearch {
		ins := ins.Clone()
		ins.Reduce(variable, value)
		sat, assignment = Solve(ins, guesserIn, completeSearch, completeSearcher)
		if sat {
			assignment[variable] = value
			return sat, assignment
		}
	}
	// fail path
	return false, nil
}
