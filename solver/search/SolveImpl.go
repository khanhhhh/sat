package search

import "github.com/khanhhhh/sat/instance"

// Guesser :
type Guesser func(ins instance.Instance) (variableOut instance.Variable, vaueOut bool)

// Solve :
func Solve(ins instance.Instance, guesserIn Guesser) (sat bool, assignment map[instance.Variable]bool) {
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
	variable, value := guesserIn(ins)
	// accept path
	{
		ins := ins.Clone()
		ins.Reduce(variable, value)
		sat, assignment = Solve(ins, guesserIn)
		if sat {
			assignment[variable] = value
			return sat, assignment
		}
	}
	// reject path
	{
		ins := ins.Clone()
		ins.Reduce(variable, value)
		sat, assignment = Solve(ins, guesserIn)
		if sat {
			assignment[variable] = value
			return sat, assignment
		}
	}
	// fail path
	return false, nil
}
