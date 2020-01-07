package instance

import "math/rand"

// Variable :
type Variable = int

// Clause :
type Clause = int

// Instance :
// SAT Instance
type Instance interface {
	// basic
	Clone() (InstanceOut Instance)
	PushClause(variableMap map[Variable]bool)
	Reduce(variableIn Variable, valueIn bool)
	Evaluate(assignment map[Variable]bool) (sat bool, conflict Clause)
	// raw data
	VariableMap() (mapOut map[Variable]map[Clause]bool)
	ClauseMap() (mapOut map[Clause]map[Variable]bool)
}

// NewInstance :
// New empty SAT Instance
func NewInstance() (InstanceOut Instance) {
	return &instance{
		make(map[Variable]map[Clause]bool),
		make(map[Clause]map[Variable]bool),
	}
}

// Random3SAT :
// create a randomly generated 3-SAT Instance
func Random3SAT(numVariables int, density float64) (InstanceOut Instance) {
	numClauses := int(density * float64(numVariables))
	InstanceOut = NewInstance()
	for c := 0; c < numClauses; c++ {
		variableMap := make(map[Variable]bool)
		for i := 0; i < 3; i++ {
			v := rand.Intn(numVariables)
			s := (rand.Intn(2) == 1)
			variableMap[v] = s
		}
		InstanceOut.PushClause(variableMap)
	}
	return InstanceOut
}
