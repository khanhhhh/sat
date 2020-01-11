package search

import (
	"fmt"

	"github.com/khanhhhh/sat/instance"
)

// Guesser :
type Guesser func(ins instance.Instance) (converged bool, variableOut instance.Variable, vaueOut bool)

// Searcher :
type Searcher func(ins instance.Instance) (sat bool, assignment map[instance.Variable]bool)

// IsTrivialUnsat :
func isTrivialUnsat(ins instance.Instance) (Unsat bool) {
	for _, clause := range ins.ClauseMap() {
		if len(clause) == 0 {
			return true
		}
	}
	return false
}

// IsTrivialSat :
func isTrivialSat(ins instance.Instance) (Sat bool) {
	return len(ins.ClauseMap()) == 0
}

// Solve :
func Solve(ins instance.Instance, guesserIn Guesser, completeSearch bool, completeSearcher Searcher) (sat bool, assignment map[instance.Variable]bool) {
	// terminating condition
	if isTrivialSat(ins) {
		return true, make(map[instance.Variable]bool)
	}
	if isTrivialUnsat(ins) {
		return false, nil
	}
	// guess and branch out
	converged, variable, value := guesserIn(ins)
	// unconverged path
	if converged == false {
		fmt.Println("trying complete search:", len(ins.VariableMap()))
		sat, assignment = completeSearcher(ins)
		return sat, assignment
	}
	// accept path
	{
		ins := ins.Clone()
		ins.Reduce(variable, value)
		fmt.Println("trying accept path:", len(ins.VariableMap()))
		sat, assignment = Solve(ins, guesserIn, completeSearch, completeSearcher)
		if sat {
			assignment[variable] = value
			return sat, assignment
		}
	}
	// reject path
	if completeSearch {
		ins := ins.Clone()
		fmt.Println("trying reject path:", len(ins.VariableMap()))
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
