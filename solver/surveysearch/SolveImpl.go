package surveysearch

import (
	"fmt"
	"time"

	"github.com/khanhhhh/sat/guesser/surveydecimation"
	"github.com/khanhhhh/sat/guesser/unitpropagation"
	"github.com/khanhhhh/sat/instance"
	"github.com/khanhhhh/sat/solver/cdcl"
	"github.com/khanhhhh/sat/solver/search"
)

var guesser = func(ins instance.Instance) (converged bool, variableOut instance.Variable, valueOut bool) {
	var nonTrivial bool
	// unitpropagation
	converged, variableOut, valueOut = unitpropagation.Guess(ins)
	if converged {
		return converged, variableOut, valueOut
	}
	// surveydecimation
	t := time.Now()
	converged, nonTrivial, variableOut, valueOut = surveydecimation.Guess(ins, 1.0)
	fmt.Println("\tGuess time:", time.Since(t))
	if converged && nonTrivial {
		return converged, variableOut, valueOut
	}
	return converged, variableOut, valueOut
}

// Solve :
func Solve(ins instance.Instance) (sat bool, assignment map[instance.Variable]bool) {
	return search.Solve(ins, guesser, true, cdcl.Solve)
}
